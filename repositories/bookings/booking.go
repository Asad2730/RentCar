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

	bookingRequest, err := GetBookingRequest(id)

	if err != nil {
		return nil, err
	}

	bookingRequestResult, err := getBookingRequestResults(c, bookingRequest)

	if err != nil {
		return nil, err
	}

	return bookingRequestResult, nil
}

func GetBookingByBoId(c *gin.Context) (*models.BookingList, error) {
	id, err := repositories.ConvertToInt32(c.Param("boId"))

	if err != nil {
		return nil, err
	}
	bookingRequests, err := GetBookingRequestByBoId(id)

	if err != nil {
		return nil, err
	}

	bookingRequestList, err := populateBookingList(c, bookingRequests)
	if err != nil {
		return nil, err
	}

	return bookingRequestList, nil

}

func GetBookingByCustomerId(c *gin.Context) (*models.BookingList, error) {
	id, err := repositories.ConvertToInt32(c.Param("customerId"))

	if err != nil {
		return nil, err
	}
	bookingRequests, err := GetBookingRequestByCustomerId(id)

	if err != nil {
		return nil, err
	}

	bookingRequestList, err := populateBookingList(c, bookingRequests)
	if err != nil {
		return nil, err
	}

	return bookingRequestList, nil

}

func GetBookings(c *gin.Context) (*models.BookingList, error) {

	var bookingRequests []*models.BookingRequest
	if err := conn.Db.Find(&bookingRequests, &models.BookingRequest{DeletedAt: ""}).Error; err != nil {
		return nil, err
	}
	bookingRequestList, err := populateBookingList(c, bookingRequests)
	if err != nil {
		return nil, err
	}
	return bookingRequestList, nil
}

func populateBookingList(c *gin.Context, bookingRequests []*models.BookingRequest) (*models.BookingList, error) {

	bookingRequestList := &models.BookingList{}
	for _, bookingRequest := range bookingRequests {
		bookingRequestResult, err := getBookingRequestResults(c, bookingRequest)
		if err != nil {
			return nil, err
		}
		bookingRequestList.Bookings = append(bookingRequestList.Bookings, bookingRequestResult)
	}

	return bookingRequestList, nil
}

func getBookingRequestResults(c *gin.Context, request *models.BookingRequest) (*models.BookingResult, error) {

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
