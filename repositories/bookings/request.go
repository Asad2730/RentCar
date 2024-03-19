package bookings

import (
	"time"

	"github.com/Asad2730/RentCar/conn"
	"github.com/Asad2730/RentCar/models"
	"github.com/Asad2730/RentCar/repositories"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func getBookingRequest(condition *models.BookingRequest) ([]*models.BookingRequest, error) {
	var body []*models.BookingRequest
	if err := conn.Db.Find(&body, condition).Error; err != nil {
		return nil, err
	}
	return body, nil
}

func InsertBookingRequest(c *gin.Context) error {
	var bookingRequest models.BookingRequest

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &bookingRequest); err != nil {
		return err
	}

	bookingRequest.CreatedAt = time.Now().String()

	if err := conn.Db.Create(&bookingRequest).Error; err != nil {
		return err
	}

	return nil
}

func GetBookingRequest(id int32) (*models.BookingRequest, error) {

	var body models.BookingRequest
	if err := conn.Db.First(&body, &models.BookingRequest{
		Id:        id,
		DeletedAt: "",
	}).Error; err != nil {
		return nil, err
	}

	return &body, nil
}

func GetBookingRequestByBoId(id int32) ([]*models.BookingRequest, error) {

	body, err := getBookingRequest(&models.BookingRequest{BoId: id, DeletedAt: ""})
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetBookingRequestByCustomerId(id int32) ([]*models.BookingRequest, error) {

	body, err := getBookingRequest(&models.BookingRequest{CustomerId: id, DeletedAt: ""})
	if err != nil {
		return nil, err
	}
	return body, nil
}

func UpdateBookingRequest(c *gin.Context) (*models.BookingRequest, error) {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return nil, err
	}

	var bookingRequest models.BookingRequest
	if err := conn.Db.First(&bookingRequest, &models.BookingRequest{Id: id}).Error; err != nil {
		return nil, err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(data, &bookingRequest); err != nil {
		return nil, err
	}

	bookingRequest.UpdatedAt = time.Now().String()

	if err := conn.Db.Save(&bookingRequest).Error; err != nil {
		return nil, err
	}
	return &bookingRequest, nil
}

func DeleteBookingRequest(c *gin.Context) error {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return err
	}

	var bookingRequest models.BookingRequest
	if err := conn.Db.First(&bookingRequest, &models.BookingRequest{Id: id}).Error; err != nil {
		return err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &bookingRequest); err != nil {
		return err
	}

	bookingRequest.DeletedAt = time.Now().String()

	if err := conn.Db.Save(&bookingRequest).Error; err != nil {
		return err
	}
	return nil
}
