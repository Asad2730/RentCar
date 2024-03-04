package conn

import (
	"github.com/Asad2730/RentCar/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=123 dbname=CarRental port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Car{})
	db.AutoMigrate(&models.BodyType{})
	db.AutoMigrate(&models.Color{})
	db.AutoMigrate(&models.EngineType{})
	db.AutoMigrate(&models.Manufacturer{})
	db.AutoMigrate(&models.Model{})
	Db = db
}
