package handlers

import (
	"github.com/Asad2730/RentCar/repositories/car"
	"github.com/Asad2730/RentCar/repositories/response"
	"github.com/gin-gonic/gin"
)

func InsertCar(c *gin.Context) {

	if err := car.InsertCar(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}

	c.ProtoBuf(201, response.GetResponse("Car Added!"))
}

func GetCar(c *gin.Context) {
	res, err := car.GetCar(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func GetCars(c *gin.Context) {
	res, err := car.GetCars(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func UpdateCar(c *gin.Context) {
	res, err := car.UpdateCar(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func DeleteCar(c *gin.Context) {

	if err := car.DeleteCar(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, response.GetResponse("Deleted SuccessFully!"))
}
