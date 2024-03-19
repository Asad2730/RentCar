package handlers

import (
	"github.com/Asad2730/RentCar/repositories/customers"
	"github.com/Asad2730/RentCar/repositories/response"
	"github.com/gin-gonic/gin"
)

func InsertCustomer(c *gin.Context) {

	if err := customers.InsertCustomer(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}

	c.ProtoBuf(201, response.GetResponse("Car Added!"))
}

func UpdateCustomer(c *gin.Context) {
	res, err := customers.UpdateCustomer(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func DeleteCustomer(c *gin.Context) {

	if err := customers.DeleteCustomer(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, response.GetResponse("Deleted SuccessFully!"))
}
