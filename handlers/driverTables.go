package handlers

import (
	"github.com/Asad2730/RentCar/repositories/drivers"
	"github.com/Asad2730/RentCar/repositories/response"
	"github.com/gin-gonic/gin"
)

// Driver
func InsertDriver(c *gin.Context) {

	if err := drivers.InsertDriver(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}

	c.ProtoBuf(201, response.GetResponse("Driver Added!"))
}

func GetDriver(c *gin.Context) {
	res, err := drivers.GetDriver(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func GetDrivers(c *gin.Context) {
	res, err := drivers.GetDrivers(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func UpdateDriver(c *gin.Context) {
	res, err := drivers.UpdateDriver(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func DeleteDriver(c *gin.Context) {

	if err := drivers.DeleteDriver(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, response.GetResponse("Deleted SuccessFully!"))
}

// Address
func InsertAddress(c *gin.Context) {
	if err := drivers.InsertAddress(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(201, response.GetResponse("Address Added!"))
}

func UpdateAddress(c *gin.Context) {
	res, err := drivers.UpdateAddress(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func DeleteAddress(c *gin.Context) {
	if err := drivers.DeleteAddress(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, response.GetResponse("Deleted SuccessFully!"))
}

// Document
func InsertDocument(c *gin.Context) {
	if err := drivers.InsertDocument(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(201, response.GetResponse("Document Added!"))
}

func UpdateDocument(c *gin.Context) {
	res, err := drivers.UpdateDocument(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func DeleteDocument(c *gin.Context) {
	if err := drivers.DeleteDocument(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, response.GetResponse("Deleted Successfully!"))
}

// DocumentType
func InsertDocumentType(c *gin.Context) {
	if err := drivers.InsertDocumentType(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(201, response.GetResponse("DocumentType Added!"))
}

func GetDocumentsType(c *gin.Context) {
	res, err := drivers.GetDocumentsType(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}

	c.ProtoBuf(200, res)
}

func UpdateDocumentType(c *gin.Context) {
	res, err := drivers.UpdateDocumentType(c)
	if err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, res)
}

func DeleteDocumentType(c *gin.Context) {
	if err := drivers.DeleteDocumentType(c); err != nil {
		c.ProtoBuf(500, response.GetResponse(err.Error()))
		return
	}
	c.ProtoBuf(200, response.GetResponse("Deleted Successfully!"))
}
