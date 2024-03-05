package drivers

import (
	"time"

	"github.com/Asad2730/RentCar/conn"
	"github.com/Asad2730/RentCar/models"
	"github.com/Asad2730/RentCar/repositories"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func GetDriverCount() (int, error) {
	var count int64
	if err := conn.Db.Model(&models.Driver{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func InsertDriver(c *gin.Context) error {
	var body models.Driver

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &body); err != nil {
		return err
	}

	count, err := GetDriverCount()
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

func GetDriver(c *gin.Context) (*models.DriverResult, error) {

	id, err := repositories.ConvertToInt32(c.Param("id"))

	if err != nil {
		return nil, err
	}

	var body models.Driver
	if err := conn.Db.Where(&models.Driver{
		Id:        id,
		DeletedAt: "",
	}).First(&body).Error; err != nil {
		return nil, err
	}

	driverResult, err := GetDriverResults(c, &body)

	if err != nil {
		return nil, err
	}
	return driverResult, nil
}

func GetDrivers(c *gin.Context) (*models.DriverList, error) {
	var drivers []*models.Driver
	if err := conn.Db.Where(&models.Driver{
		DeletedAt: "",
	}).Find(&drivers).Error; err != nil {
		return nil, err
	}
	driverList := &models.DriverList{}

	for _, car := range drivers {
		driverResult, err := GetDriverResults(c, car)
		if err != nil {
			return nil, err
		}
		driverList.DriverResult = append(driverList.DriverResult, driverResult)
	}

	return driverList, nil
}

func UpdateDriver(c *gin.Context) (*models.Driver, error) {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return nil, err
	}

	var driver models.Driver
	if err := conn.Db.Where(&models.Driver{Id: id}).First(&driver).Error; err != nil {
		return nil, err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(data, &driver); err != nil {
		return nil, err
	}

	driver.UpdatedAt = time.Now().String()

	if err := conn.Db.Save(&driver).Error; err != nil {
		return nil, err
	}
	return &driver, nil
}

func DeleteDriver(c *gin.Context) error {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return err
	}

	var driver models.Driver
	if err := conn.Db.Where(&models.Car{Id: id}).First(&driver).Error; err != nil {
		return err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &driver); err != nil {
		return err
	}

	driver.DeletedAt = time.Now().String()

	if err := conn.Db.Save(&driver).Error; err != nil {
		return err
	}
	return nil
}

func GetDriverResults(c *gin.Context, driver *models.Driver) (*models.DriverResult, error) {

	address, err := GetAddress(c, driver.AddressId)
	if err != nil {
		return nil, err
	}
	documnetResult, err := GetDocumentResult(c, driver.DocumentId)
	if err != nil {
		return nil, err
	}

	driverResult := &models.DriverResult{
		Driver:         driver,
		Address:        address,
		DocumentResult: documnetResult,
	}
	return driverResult, nil
}