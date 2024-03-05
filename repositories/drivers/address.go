package drivers

import (
	"time"

	"github.com/Asad2730/RentCar/conn"
	"github.com/Asad2730/RentCar/models"
	"github.com/Asad2730/RentCar/repositories"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func GetAddressCount() (int, error) {
	var count int64
	if err := conn.Db.Model(&models.Address{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func InsertAddress(c *gin.Context) error {
	var address models.Address

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &address); err != nil {
		return err
	}

	count, err := GetAddressCount()
	if err != nil {
		return err
	}
	address.Id = int32(count)
	address.CreatedAt = time.Now().String()

	if err := conn.Db.Create(&address).Error; err != nil {
		return err
	}

	return nil
}

func GetAddress(c *gin.Context, id int32) (*models.Address, error) {

	var address models.Address
	if err := conn.Db.Where(&models.Address{
		Id:        id,
		DeletedAt: "",
	}).First(&address).Error; err != nil {
		return nil, err
	}

	return &address, nil
}

func UpdateAddress(c *gin.Context) (*models.Address, error) {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return nil, err
	}

	var address models.Address
	if err := conn.Db.Where(&models.Address{Id: id}).First(&address).Error; err != nil {
		return nil, err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(data, &address); err != nil {
		return nil, err
	}

	address.UpdatedAt = time.Now().String()

	if err := conn.Db.Save(&address).Error; err != nil {
		return nil, err
	}
	return &address, nil
}

func DeleteAddress(c *gin.Context) error {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return err
	}

	var address models.Address
	if err := conn.Db.Where(&models.Address{Id: id}).First(&address).Error; err != nil {
		return err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &address); err != nil {
		return err
	}

	address.DeletedAt = time.Now().String()

	if err := conn.Db.Save(&address).Error; err != nil {
		return err
	}
	return nil
}
