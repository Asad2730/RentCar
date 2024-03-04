package response

import "github.com/Asad2730/RentCar/models"

func GetResponse(msg string) *models.ResponseMsg {
	return &models.ResponseMsg{
		Message: msg,
	}
}
