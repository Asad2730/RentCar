package cars

import (
	"time"

	"github.com/Asad2730/RentCar/conn"
	"github.com/Asad2730/RentCar/models"
	"github.com/Asad2730/RentCar/repositories"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func GetColorCount() (int, error) {
	var count int64
	if err := conn.Db.Model(&models.Color{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func InsertColor(c *gin.Context) error {
	var body models.Color

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &body); err != nil {
		return err
	}

	count, err := GetColorCount()
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

func GetColor(c *gin.Context, id int32) (*models.Color, error) {

	var body models.Color
	if err := conn.Db.Where(&models.Color{
		CarId:     id,
		DeletedAt: "",
	}).First(&body).Error; err != nil {
		return nil, err
	}

	return &body, nil
}

func UpdateColor(c *gin.Context) (*models.Color, error) {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return nil, err
	}

	var body models.Color
	if err := conn.Db.Where(&models.Color{Id: id}).First(&body).Error; err != nil {
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

func DeleteColor(c *gin.Context) error {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return err
	}

	var body models.Color
	if err := conn.Db.Where(&models.Color{Id: id}).First(&body).Error; err != nil {
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
