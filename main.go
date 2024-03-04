package main

import (
	"github.com/Asad2730/RentCar/conn"
	"github.com/Asad2730/RentCar/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	conn.Connect()
}

func main() {
	r := gin.Default()
	api := r.Group("/api")
	car := api.Group("/car")
	routes.CartRoutes(car)
	r.Run("0.0.0.0:3000")
}
