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

	db.AutoMigrate(&models.Car{}, &models.BodyType{}, &models.Color{},
		&models.EngineType{}, &models.Manufacturer{}, &models.Model{})

	db.AutoMigrate(&models.Driver{}, &models.Address{}, &models.Document{}, &models.DocumentType{})

	db.AutoMigrate(&models.BookingRequest{}, &models.BookingClosing{},
		&models.BookingFeedBack{}, &models.BookingDetail{})

	Db = db
}
