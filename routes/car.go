package routes

import (
	"github.com/Asad2730/RentCar/handlers"
	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.RouterGroup) {
	r.GET("/", handlers.GetCars)
	r.POST("/", handlers.InsertCar)
	r.GET("/:id", handlers.GetCar)
	r.PUT("/:id", handlers.UpdateCar)
	r.DELETE("/:id", handlers.DeleteCar)
}
