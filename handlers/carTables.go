package handlers

import (
	"github.com/Asad2730/RentCar/repositories/cars"
	"github.com/Asad2730/RentCar/repositories/response"
	"github.com/gin-gonic/gin"
)

// Car
func InsertCar(c *gin.Context) {

	if err := cars.InsertCar(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}

	c.ProtoBuf(201, response.GetResponse("Car Added!"))
}

func GetCar(c *gin.Context) {
	res, err := cars.GetCar(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func GetCars(c *gin.Context) {
	res, err := cars.GetCars(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func UpdateCar(c *gin.Context) {
	res, err := cars.UpdateCar(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func DeleteCar(c *gin.Context) {

	if err := cars.DeleteCar(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, response.GetResponse("Deleted SuccessFully!"))
}

// BodyType
func InsertBodyType(c *gin.Context) {
	if err := cars.InsertBodyType(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(201, response.GetResponse("BodyType Added!"))
}

func UpdateBodyType(c *gin.Context) {
	res, err := cars.UpdateBodyType(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func DeleteBodyType(c *gin.Context) {
	if err := cars.DeleteBodyType(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, response.GetResponse("Deleted SuccessFully!"))
}

// Color
func InsertColor(c *gin.Context) {
	if err := cars.InsertColor(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(201, response.GetResponse("Color Added!"))
}

func UpdateColor(c *gin.Context) {
	res, err := cars.UpdateColor(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func DeleteColor(c *gin.Context) {
	if err := cars.DeleteColor(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, response.GetResponse("Deleted Successfully!"))
}

// EngineType
func InsertEngineType(c *gin.Context) {
	if err := cars.InsertEngineType(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(201, response.GetResponse("EngineType Added!"))
}

func UpdateEngineType(c *gin.Context) {
	res, err := cars.UpdateEngineType(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func DeleteEngineType(c *gin.Context) {
	if err := cars.DeleteEngineType(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, response.GetResponse("Deleted Successfully!"))
}

// Manufacturer
func InsertManufacturer(c *gin.Context) {
	if err := cars.InsertManufacturer(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(201, response.GetResponse("Manufacturer Added!"))
}

func UpdateManufacturer(c *gin.Context) {
	res, err := cars.UpdateManufacturer(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func DeleteManufacturer(c *gin.Context) {
	if err := cars.DeleteManufacturer(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, response.GetResponse("Deleted Successfully!"))
}

// Model
func InsertModel(c *gin.Context) {
	if err := cars.InsertModel(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(201, response.GetResponse("Model Added!"))
}

func UpdateModel(c *gin.Context) {
	res, err := cars.UpdateModel(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func DeleteModel(c *gin.Context) {
	if err := cars.DeleteModel(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, response.GetResponse("Deleted Successfully!"))
}
