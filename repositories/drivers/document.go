package drivers

import (
	"time"

	"github.com/Asad2730/RentCar/conn"
	"github.com/Asad2730/RentCar/models"
	"github.com/Asad2730/RentCar/repositories"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func InsertDocument(c *gin.Context) error {
	var document models.Document

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &document); err != nil {
		return err
	}

	document.CreatedAt = time.Now().String()

	if err := conn.Db.Create(&document).Error; err != nil {
		return err
	}

	return nil
}

func GetDocumentResult(c *gin.Context, id int32) (*models.DocumentResult, error) {

	var document models.Document
	if err := conn.Db.Where(&models.Document{
		DriverId:  id,
		DeletedAt: "",
	}).First(&document).Error; err != nil {
		return nil, err
	}

	doumentType, err := GetDocumentType(c, document.DocumentTypeId)

	if err != nil {
		return nil, err
	}
	documentResult := &models.DocumentResult{
		Document:     &document,
		DocumentType: doumentType,
	}

	return documentResult, nil
}

func UpdateDocument(c *gin.Context) (*models.Document, error) {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return nil, err
	}

	var document models.Document
	if err := conn.Db.Where(&models.Document{Id: id}).First(&document).Error; err != nil {
		return nil, err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(data, &document); err != nil {
		return nil, err
	}

	document.UpdatedAt = time.Now().String()

	if err := conn.Db.Save(&document).Error; err != nil {
		return nil, err
	}
	return &document, nil
}

func DeleteDocument(c *gin.Context) error {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return err
	}

	var document models.Document
	if err := conn.Db.Where(&models.Document{Id: id}).First(&document).Error; err != nil {
		return err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &document); err != nil {
		return err
	}

	document.DeletedAt = time.Now().String()

	if err := conn.Db.Save(&document).Error; err != nil {
		return err
	}
	return nil
}
