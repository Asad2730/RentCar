package cars

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
	if err := conn.Db.Where(&models.Car{
		Id:        id,
		DeletedAt: "",
	}).First(&car).Error; err != nil {
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
	if err := conn.Db.Where(&models.Car{
		DeletedAt: "",
	}).Find(&cars).Error; err != nil {
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

	car.DeletedAt = time.Now().String()

	if err := conn.Db.Save(&car).Error; err != nil {
		return err
	}
	return nil
}

func GetCarResults(c *gin.Context, car *models.Car) (*models.CarResult, error) {
	model, err := GetModel(c, car.ModelId)
	if err != nil {
		return nil, err
	}

	manufacturer, err := GetManufacturer(c, car.ManufacturerId)
	if err != nil {
		return nil, err
	}

	color, err := GetColor(c, car.ColorId)
	if err != nil {
		return nil, err
	}

	bodyType, err := GetBodyType(c, car.BodyTypeId)
	if err != nil {
		return nil, err
	}

	engineType, err := GetEngineType(c, car.EngineTypeId)
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
