package cars

import (
	"time"

	"github.com/Asad2730/RentCar/conn"
	"github.com/Asad2730/RentCar/models"
	"github.com/Asad2730/RentCar/repositories"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func GetBodyTypeCount() (int, error) {
	var count int64
	if err := conn.Db.Model(&models.BodyType{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func InsertBodyType(c *gin.Context) error {
	var body models.BodyType

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &body); err != nil {
		return err
	}

	count, err := GetBodyTypeCount()
	if err != nil {
		return err
	}
	body.Id = int32(count)
	body.TimeStamp.CreatedAt = time.Now().String()

	if err := conn.Db.Create(&body).Error; err != nil {
		return err
	}

	return nil
}

func GetBodyType(c *gin.Context, id int32) (*models.BodyType, error) {

	var body models.BodyType
	if err := conn.Db.Where(&models.BodyType{Id: id,
		TimeStamp: &models.TimeStamp{
			DeletedAt: "",
		}}).First(&body).Error; err != nil {
		return nil, err
	}

	return &body, nil
}

func UpdateBodyType(c *gin.Context) (*models.BodyType, error) {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return nil, err
	}

	var body models.BodyType
	if err := conn.Db.Where(&models.BodyType{Id: id}).First(&body).Error; err != nil {
		return nil, err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(data, &body); err != nil {
		return nil, err
	}

	body.TimeStamp.UpdatedAt = time.Now().String()

	if err := conn.Db.Save(&body).Error; err != nil {
		return nil, err
	}
	return &body, nil
}

func DeleteBodyType(c *gin.Context) error {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return err
	}

	var body models.BodyType
	if err := conn.Db.Where(&models.BodyType{Id: id}).First(&body).Error; err != nil {
		return err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &body); err != nil {
		return err
	}

	body.TimeStamp.DeletedAt = time.Now().String()

	if err := conn.Db.Save(&body).Error; err != nil {
		return err
	}
	return nil
}
