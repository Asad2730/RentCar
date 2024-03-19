package customers

import (
	"time"

	"github.com/Asad2730/RentCar/conn"
	"github.com/Asad2730/RentCar/models"
	"github.com/Asad2730/RentCar/repositories"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func InsertCustomer(c *gin.Context) error {
	var customer models.Customer
	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &customer); err != nil {
		return err
	}

	customer.CreatedAt = time.Now().String()
	if err := conn.Db.Create(&customer).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCustomer(c *gin.Context) (*models.Customer, error) {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return nil, err
	}

	var customer models.Customer
	if err := conn.Db.First(&customer, &models.Customer{Id: id}).Error; err != nil {
		return nil, err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(data, &customer); err != nil {
		return nil, err
	}

	customer.UpdatedAt = time.Now().String()

	if err := conn.Db.Save(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func DeleteCustomer(c *gin.Context) error {
	id, err := repositories.ConvertToInt32(c.Param("id"))
	if err != nil {
		return err
	}

	var customer models.Customer
	if err := conn.Db.First(&customer, &models.Car{Id: id}).Error; err != nil {
		return err
	}

	data, err := repositories.ReadRequest(c.Request.Body)
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(data, &customer); err != nil {
		return err
	}

	customer.DeletedAt = time.Now().String()

	if err := conn.Db.Save(&customer).Error; err != nil {
		return err
	}
	return nil
}
