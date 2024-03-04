package car

import (
	"time"

	"github.com/Asad2730/RentCar/conn"
	"github.com/Asad2730/RentCar/models"
	"github.com/Asad2730/RentCar/repositories"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func GetCarCount() (int, error) {
	var count int64
	if err := conn.Db.Model(&models.Car{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func InsertCar(c *gin.Context) error {
	var car models.Car

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &car); err != nil {
		return err
	}

	count, err := GetCarCount()
	if err != nil {
		return err
	}
	car.Id = int32(count)
	car.TimeStamp.CreatedAt = time.Now().String()

	if err := conn.Db.Create(&car).Error; err != nil {
		return err
	}

	return nil
}

func GetCar(c *gin.Context) (*models.Car, error) {
	id, err := repositories.ConvertToInt32(c.Param("id"))

	if err != nil {
		return nil, err
	}

	var car models.Car
	if err := conn.Db.Where(&models.Car{Id: id,
		TimeStamp: &models.TimeStamp{
			DeletedAt: "",
		}}).First(&car).Error; err != nil {
		return nil, err
	}

	return &car, nil
}

func GetCars(c *gin.Context) (*models.CarList, error) {
	var cars []*models.Car
	if err := conn.Db.Where(&models.Car{
		TimeStamp: &models.TimeStamp{
			DeletedAt: "",
		},
	}).Find(&cars).Error; err != nil {
		return nil, err
	}
	carList := &models.CarList{
		Cars: cars,
	}

	return carList, nil
}

func UpdateCar(c *gin.Context) (*models.Car, error) {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return nil, err
	}

	var car models.Car
	if err := conn.Db.Where(&models.Car{Id: id}).First(&car).Error; err != nil {
		return nil, err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(data, &car); err != nil {
		return nil, err
	}

	car.TimeStamp.UpdatedAt = time.Now().String()

	if err := conn.Db.Save(&car).Error; err != nil {
		return nil, err
	}
	return &car, nil
}

func DeleteCar(c *gin.Context) error {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return err
	}

	var car models.Car
	if err := conn.Db.Where(&models.Car{Id: id}).First(&car).Error; err != nil {
		return err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &car); err != nil {
		return err
	}

	car.TimeStamp.DeletedAt = time.Now().String()

	if err := conn.Db.Save(&car).Error; err != nil {
		return err
	}
	return nil
}
