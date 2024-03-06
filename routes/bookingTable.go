package routes

import (
	"github.com/Asad2730/RentCar/handlers"
	"github.com/gin-gonic/gin"
)

func Booking(r *gin.RouterGroup) {
	r.GET("/", handlers.GetBookings)
	r.GET("/:id", handlers.GetBooking)
}

func BookingRequestRoutes(r *gin.RouterGroup) {
	r.POST("/", handlers.InsertBookingRequest)
	r.PUT("/:id", handlers.UpdateBookingRequest)
	r.DELETE("/:id", handlers.DeleteBookingRequest)
}

func BookingClosingRoutes(r *gin.RouterGroup) {
	r.POST("/", handlers.InsertBookingClosing)
	r.PUT("/:id", handlers.UpdateBookingClosing)
	r.DELETE("/:id", handlers.DeleteBookingClosing)
}

func BookingDetailRoutes(r *gin.RouterGroup) {
	r.POST("/", handlers.InsertBookingDetail)
	r.PUT("/:id", handlers.UpdateBookingDetail)
	r.DELETE("/:id", handlers.DeleteBookingDetail)
}

func BookingFeedbackRoutes(r *gin.RouterGroup) {
	r.POST("/", handlers.InsertBookingFeedBack)
	r.PUT("/:id", handlers.UpdateBookingFeedBack)
	r.DELETE("/:id", handlers.DeleteBookingFeedBack)
}
