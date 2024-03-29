package bookings

import (
	"time"

	"github.com/Asad2730/RentCar/conn"
	"github.com/Asad2730/RentCar/models"
	"github.com/Asad2730/RentCar/repositories"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func InsertBookingDetail(c *gin.Context) error {
	var bookingDetail models.BookingDetail
	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &bookingDetail); err != nil {
		return err
	}

	bookingDetail.CreatedAt = time.Now().String()

	if err := conn.Db.Create(&bookingDetail).Error; err != nil {
		return err
	}

	return nil
}

func GetBookingDetail(c *gin.Context, id int32) (*models.BookingDetail, error) {

	var body models.BookingDetail
	if err := conn.Db.First(&body, &models.BookingDetail{
		BookingId: id,
		DeletedAt: "",
	}).Error; err != nil {
		return nil, err
	}

	return &body, nil
}

func UpdateBookingDetail(c *gin.Context) (*models.BookingDetail, error) {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return nil, err
	}

	var bookingDetail models.BookingDetail
	if err := conn.Db.First(&bookingDetail, &models.BookingDetail{Id: id}).Error; err != nil {
		return nil, err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(data, &bookingDetail); err != nil {
		return nil, err
	}

	bookingDetail.UpdatedAt = time.Now().String()

	if err := conn.Db.Save(&bookingDetail).Error; err != nil {
		return nil, err
	}
	return &bookingDetail, nil
}

func DeleteBookingDetail(c *gin.Context) error {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return err
	}

	var bookingDetail models.BookingDetail
	if err := conn.Db.First(&bookingDetail, &models.BookingDetail{Id: id}).Error; err != nil {
		return err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &bookingDetail); err != nil {
		return err
	}

	bookingDetail.DeletedAt = time.Now().String()

	if err := conn.Db.Save(&bookingDetail).Error; err != nil {
		return err
	}
	return nil
}
