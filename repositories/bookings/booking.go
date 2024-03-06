package bookings

import (
	"github.com/Asad2730/RentCar/conn"
	"github.com/Asad2730/RentCar/models"
	"github.com/Asad2730/RentCar/repositories"
	"github.com/gin-gonic/gin"
)

// see which model our hich id should be passed in them...
// Booking id is of BookingRequest Table so just use that..

func GetBooking(c *gin.Context) (*models.BookingResult, error) {

	id, err := repositories.ConvertToInt32(c.Param("id"))

	if err != nil {
		return nil, err
	}

	var bookingRequest models.BookingRequest
	if err := conn.Db.Where(&models.BookingRequest{
		Id:        id,
		DeletedAt: "",
	}).First(&bookingRequest).Error; err != nil {
		return nil, err
	}

	bookingRequestResult, err := GetBookingRequestResults(c, &bookingRequest)

	if err != nil {
		return nil, err
	}

	return bookingRequestResult, nil
}

func GetBookings(c *gin.Context) (*models.BookingList, error) {
	var bookingRequests []*models.BookingRequest
	if err := conn.Db.Where(&models.BookingRequest{
		DeletedAt: "",
	}).Find(&bookingRequests).Error; err != nil {
		return nil, err
	}
	bookingRequestList := &models.BookingList{}

	for _, bookingRequest := range bookingRequests {
		bookingRequestResult, err := GetBookingRequestResults(c, bookingRequest)
		if err != nil {
			return nil, err
		}
		bookingRequestList.Bookings = append(bookingRequestList.Bookings, bookingRequestResult)
	}

	return bookingRequestList, nil
}

// make get functions in others aswell...
func GetBookingRequestResults(c *gin.Context, bookingRequest *models.BookingRequest) (*models.BookingResult, error) {
	return nil, nil
}
