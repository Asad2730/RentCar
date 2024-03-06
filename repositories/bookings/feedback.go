package bookings

import (
	"time"

	"github.com/Asad2730/RentCar/conn"
	"github.com/Asad2730/RentCar/models"
	"github.com/Asad2730/RentCar/repositories"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func InsertBookingFeedback(c *gin.Context) error {
	var bookingFeedback models.BookingFeedBack
	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &bookingFeedback); err != nil {
		return err
	}

	bookingFeedback.CreatedAt = time.Now().String()

	if err := conn.Db.Create(&bookingFeedback).Error; err != nil {
		return err
	}

	return nil
}

func UpdateBookingFeedback(c *gin.Context) (*models.BookingFeedBack, error) {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return nil, err
	}

	var bookingFeedback models.BookingFeedBack
	if err := conn.Db.Where(&models.BookingFeedBack{Id: id}).First(&bookingFeedback).Error; err != nil {
		return nil, err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(data, &bookingFeedback); err != nil {
		return nil, err
	}

	bookingFeedback.UpdatedAt = time.Now().String()

	if err := conn.Db.Save(&bookingFeedback).Error; err != nil {
		return nil, err
	}
	return &bookingFeedback, nil
}

func DeleteBookingFeedback(c *gin.Context) error {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return err
	}

	var bookingFeedback models.BookingFeedBack
	if err := conn.Db.Where(&models.BookingFeedBack{Id: id}).First(&bookingFeedback).Error; err != nil {
		return err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &bookingFeedback); err != nil {
		return err
	}

	bookingFeedback.DeletedAt = time.Now().String()

	if err := conn.Db.Save(&bookingFeedback).Error; err != nil {
		return err
	}
	return nil
}
