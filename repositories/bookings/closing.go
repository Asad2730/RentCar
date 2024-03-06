package bookings

import (
	"time"

	"github.com/Asad2730/RentCar/conn"
	"github.com/Asad2730/RentCar/models"
	"github.com/Asad2730/RentCar/repositories"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func InsertBookingClosing(c *gin.Context) error {
	var bookingClosing models.BookingClosing
	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &bookingClosing); err != nil {
		return err
	}

	bookingClosing.CreatedAt = time.Now().String()

	if err := conn.Db.Create(&bookingClosing).Error; err != nil {
		return err
	}

	return nil
}

func GetBookingClosing(c *gin.Context, id int32) (*models.BookingClosing, error) {

	var body models.BookingClosing
	if err := conn.Db.Where(&models.BookingClosing{
		BookingId: id,
		DeletedAt: "",
	}).First(&body).Error; err != nil {
		return nil, err
	}

	return &body, nil
}

func UpdateBookingClosing(c *gin.Context) (*models.BookingClosing, error) {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return nil, err
	}

	var bookingClosing models.BookingClosing
	if err := conn.Db.Where(&models.BookingClosing{Id: id}).First(&bookingClosing).Error; err != nil {
		return nil, err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(data, &bookingClosing); err != nil {
		return nil, err
	}

	bookingClosing.UpdatedAt = time.Now().String()

	if err := conn.Db.Save(&bookingClosing).Error; err != nil {
		return nil, err
	}
	return &bookingClosing, nil
}

func DeleteBookingClosing(c *gin.Context) error {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return err
	}

	var bookingClosing models.BookingClosing
	if err := conn.Db.Where(&models.BookingClosing{Id: id}).First(&bookingClosing).Error; err != nil {
		return err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &bookingClosing); err != nil {
		return err
	}

	bookingClosing.DeletedAt = time.Now().String()

	if err := conn.Db.Save(&bookingClosing).Error; err != nil {
		return err
	}
	return nil
}
