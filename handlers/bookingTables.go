package handlers

import (
	"github.com/Asad2730/RentCar/repositories/bookings"
	"github.com/Asad2730/RentCar/repositories/response"
	"github.com/gin-gonic/gin"
)

// Request
func InsertBookingRequest(c *gin.Context) {

	if err := bookings.InsertBookingRequest(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}

	c.ProtoBuf(201, response.GetResponse("Booking request send!"))
}

func UpdateBookingRequest(c *gin.Context) {
	res, err := bookings.UpdateBookingRequest(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func DeleteBookingRequest(c *gin.Context) {

	if err := bookings.DeleteBookingRequest(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, response.GetResponse("Deleted SuccessFully!"))
}

// closing

func InsertBookingClosing(c *gin.Context) {
	if err := bookings.InsertBookingClosing(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(201, response.GetResponse("BodyType Added!"))
}

func UpdateBookingClosing(c *gin.Context) {
	res, err := bookings.UpdateBookingClosing(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func DeleteBookingClosing(c *gin.Context) {
	if err := bookings.DeleteBookingClosing(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, response.GetResponse("Deleted SuccessFully!"))
}

// detail
func InsertBookingDetail(c *gin.Context) {
	if err := bookings.InsertBookingDetail(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(201, response.GetResponse("Booking Detail Added!"))
}

func UpdateBookingDetail(c *gin.Context) {
	res, err := bookings.UpdateBookingDetail(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func DeleteBookingDetail(c *gin.Context) {
	if err := bookings.DeleteBookingDetail(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, response.GetResponse("Deleted Successfully!"))
}

// feedback
func InsertBookingFeedBack(c *gin.Context) {
	if err := bookings.InsertBookingFeedback(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(201, response.GetResponse("Booking Feedback Added!"))
}

func UpdateBookingFeedBack(c *gin.Context) {
	res, err := bookings.UpdateBookingFeedback(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func DeleteBookingFeedBack(c *gin.Context) {
	if err := bookings.DeleteBookingFeedback(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, response.GetResponse("Deleted Successfully!"))
}

//booking

func GetBooking(c *gin.Context) {
	res, err := bookings.GetBooking(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func GetBookings(c *gin.Context) {
	res, err := bookings.GetBookings(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func GetBookingByCustomerId(c *gin.Context) {
	res, err := bookings.GetBookingByCustomerId(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}
