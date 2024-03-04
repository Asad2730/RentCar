package conn

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=123 dbname=CarRental port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate)
	db.AutoMigrate(&protobufModel.Product{})
	db.AutoMigrate(&protobufModel.Cart{})
	Db = db
}