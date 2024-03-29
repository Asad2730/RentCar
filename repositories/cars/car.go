package cars

import (
	"time"

	"github.com/Asad2730/RentCar/conn"
	"github.com/Asad2730/RentCar/models"
	"github.com/Asad2730/RentCar/repositories"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func InsertCar(c *gin.Context) error {
	var car models.Car

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &car); err != nil {
		return err
	}

	car.CreatedAt = time.Now().String()

	if err := conn.Db.Create(&car).Error; err != nil {
		return err
	}

	return nil
}

func GetCar(c *gin.Context) (*models.CarResult, error) {

	id, err := repositories.ConvertToInt32(c.Param("id"))

	if err != nil {
		return nil, err
	}

	var car models.Car
	if err := conn.Db.First(&car, &models.Car{
		Id:        id,
		DeletedAt: "",
	}).Error; err != nil {
		return nil, err
	}

	carResult, err := GetCarResults(c, &car)

	if err != nil {
		return nil, err
	}

	return carResult, nil
}

func GetCars(c *gin.Context) (*models.CarList, error) {
	var cars []*models.Car
	if err := conn.Db.Find(&cars, &models.Car{
		DeletedAt: "",
	}).Error; err != nil {
		return nil, err
	}
	carList := &models.CarList{}

	for _, car := range cars {
		carResult, err := GetCarResults(c, car)
		if err != nil {
			return nil, err
		}
		carList.Cars = append(carList.Cars, carResult)
	}

	return carList, nil
}

func UpdateCar(c *gin.Context) (*models.Car, error) {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return nil, err
	}

	var car models.Car
	if err := conn.Db.First(&car, &models.Car{Id: id}).Error; err != nil {
		return nil, err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(data, &car); err != nil {
		return nil, err
	}

	car.UpdatedAt = time.Now().String()

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
	if err := conn.Db.First(&car, &models.Car{Id: id}).Error; err != nil {
		return err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &car); err != nil {
		return err
	}

	car.DeletedAt = time.Now().String()

	if err := conn.Db.Save(&car).Error; err != nil {
		return err
	}
	return nil
}

func GetCarResults(c *gin.Context, car *models.Car) (*models.CarResult, error) {
	model, err := GetModel(c, car.Id)
	if err != nil {
		return nil, err
	}

	manufacturer, err := GetManufacturer(c, car.Id)
	if err != nil {
		return nil, err
	}

	color, err := GetColor(c, car.Id)
	if err != nil {
		return nil, err
	}

	bodyType, err := GetBodyType(c, car.Id)
	if err != nil {
		return nil, err
	}

	engineType, err := GetEngineType(c, car.Id)
	if err != nil {
		return nil, err
	}

	carResult := &models.CarResult{
		Car:          car,
		Model:        model,
		Manufacturer: manufacturer,
		Color:        color,
		BodyType:     bodyType,
		EngineType:   engineType,
	}

	return carResult, nil
}
