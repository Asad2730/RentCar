package drivers

import (
	"time"

	"github.com/Asad2730/RentCar/conn"
	"github.com/Asad2730/RentCar/models"
	"github.com/Asad2730/RentCar/repositories"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func GetDocumentTypeCount() (int, error) {
	var count int64
	if err := conn.Db.Model(&models.DocumentType{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func InsertDocumentType(c *gin.Context) error {
	var docType models.DocumentType

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &docType); err != nil {
		return err
	}

	count, err := GetDocumentTypeCount()
	if err != nil {
		return err
	}
	docType.Id = int32(count)
	docType.CreatedAt = time.Now().String()

	if err := conn.Db.Create(&docType).Error; err != nil {
		return err
	}

	return nil
}

func GetDocumentType(c *gin.Context, id int32) (*models.DocumentType, error) {

	var docType models.DocumentType
	if err := conn.Db.Where(&models.DocumentType{
		Id:        id,
		DeletedAt: "",
	}).First(&docType).Error; err != nil {
		return nil, err
	}

	return &docType, nil
}

func GetDocumentsType(c *gin.Context) (*models.DocumentTypeList, error) {

	var docTypes []*models.DocumentType
	if err := conn.Db.Where(&models.DocumentType{
		DeletedAt: "",
	}).Find(&docTypes).Error; err != nil {
		return nil, err
	}

	documentTypeList := &models.DocumentTypeList{
		DocumentTypeList: docTypes,
	}
	return documentTypeList, nil
}

func UpdateDocumentType(c *gin.Context) (*models.DocumentType, error) {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return nil, err
	}

	var docType models.DocumentType
	if err := conn.Db.Where(&models.DocumentType{Id: id}).First(&docType).Error; err != nil {
		return nil, err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(data, &docType); err != nil {
		return nil, err
	}

	docType.UpdatedAt = time.Now().String()

	if err := conn.Db.Save(&docType).Error; err != nil {
		return nil, err
	}
	return &docType, nil
}

func DeleteDocumentType(c *gin.Context) error {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return err
	}

	var docType models.DocumentType
	if err := conn.Db.Where(&models.DocumentType{Id: id}).First(&docType).Error; err != nil {
		return err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &docType); err != nil {
		return err
	}

	docType.DeletedAt = time.Now().String()

	if err := conn.Db.Save(&docType).Error; err != nil {
		return err
	}
	return nil
}
