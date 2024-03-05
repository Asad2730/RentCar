package cars

import (
	"time"

	"github.com/Asad2730/RentCar/conn"
	"github.com/Asad2730/RentCar/models"
	"github.com/Asad2730/RentCar/repositories"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func GetEngineTypeCount() (int, error) {
	var count int64
	if err := conn.Db.Model(&models.EngineType{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func InsertEngineType(c *gin.Context) error {
	var body models.EngineType

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &body); err != nil {
		return err
	}

	count, err := GetEngineTypeCount()
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

func GetEngineType(c *gin.Context, id int32) (*models.EngineType, error) {

	var body models.EngineType
	if err := conn.Db.Where(&models.EngineType{
		CarId:     id,
		DeletedAt: "",
	}).First(&body).Error; err != nil {
		return nil, err
	}

	return &body, nil
}

func UpdateEngineType(c *gin.Context) (*models.EngineType, error) {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return nil, err
	}

	var body models.EngineType
	if err := conn.Db.Where(&models.EngineType{Id: id}).First(&body).Error; err != nil {
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

func DeleteEngineType(c *gin.Context) error {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return err
	}

	var body models.EngineType
	if err := conn.Db.Where(&models.EngineType{Id: id}).First(&body).Error; err != nil {
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
