package cars

import (
	"time"

	"github.com/Asad2730/RentCar/conn"
	"github.com/Asad2730/RentCar/models"
	"github.com/Asad2730/RentCar/repositories"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func GetManufacturerCount() (int, error) {
	var count int64
	if err := conn.Db.Model(&models.Manufacturer{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func InsertManufacturer(c *gin.Context) error {
	var body models.Manufacturer

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &body); err != nil {
		return err
	}

	count, err := GetManufacturerCount()
	if err != nil {
		return err
	}
	body.Id = int32(count)
	body.CreatedAt = time.Now().String()

	if err := conn.Db.Create(&body).Error; err != nil {
		return err
	}

	return nil
}

func GetManufacturer(c *gin.Context, id int32) (*models.Manufacturer, error) {

	var body models.Manufacturer
	if err := conn.Db.Where(&models.Manufacturer{
		Id:        id,
		DeletedAt: "",
	}).First(&body).Error; err != nil {
		return nil, err
	}

	return &body, nil
}

func UpdateManufacturer(c *gin.Context) (*models.Manufacturer, error) {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return nil, err
	}

	var body models.Manufacturer
	if err := conn.Db.Where(&models.Manufacturer{Id: id}).First(&body).Error; err != nil {
		return nil, err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(data, &body); err != nil {
		return nil, err
	}

	body.UpdatedAt = time.Now().String()

	if err := conn.Db.Save(&body).Error; err != nil {
		return nil, err
	}
	return &body, nil
}

func DeleteManufacturer(c *gin.Context) error {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return err
	}

	var body models.Manufacturer
	if err := conn.Db.Where(&models.Manufacturer{Id: id}).First(&body).Error; err != nil {
		return err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &body); err != nil {
		return err
	}

	body.DeletedAt = time.Now().String()

	if err := conn.Db.Save(&body).Error; err != nil {
		return err
	}
	return nil
}
