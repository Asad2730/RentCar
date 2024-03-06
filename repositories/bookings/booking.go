package bookings

import (
	"github.com/Asad2730/RentCar/conn"
	"github.com/Asad2730/RentCar/models"
	"github.com/Asad2730/RentCar/repositories"
	"github.com/gin-gonic/gin"
)

func GetBooking(c *gin.Context) (*models.BookingResult, error) {

	id, err := repositories.ConvertToInt32(c.Param("id"))

	if err != nil {
		return nil, err
	}

	bookingRequest, err := GetBookingRequest(c, id)

	if err != nil {
		return nil, err
	}

	bookingRequestResult, err := GetBookingRequestResults(c, bookingRequest)

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

func GetBookingRequestResults(c *gin.Context, request *models.BookingRequest) (*models.BookingResult, error) {

	closing, err := GetBookingClosing(c, request.Id)

	if err != nil {
		return nil, err
	}

	detail, err := GetBookingDetail(c, request.Id)

	if err != nil {
		return nil, err
	}

	feedback, err := GetBookingFeedBack(c, request.Id)

	if err != nil {
		return nil, err
	}

	bookingResult := &models.BookingResult{
		BookingRequest:  request,
		BookingClosing:  closing,
		BookingDetail:   detail,
		BookingFeedBack: feedback,
	}

	return bookingResult, nil
}
