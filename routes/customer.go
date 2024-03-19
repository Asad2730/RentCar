package routes

import (
	"github.com/Asad2730/RentCar/handlers"
	"github.com/gin-gonic/gin"
)

func CustomerRoutes(r *gin.RouterGroup) {
	r.POST("/", handlers.InsertCustomer)
	r.PUT("/:id", handlers.UpdateCustomer)
	r.DELETE("/:id", handlers.DeleteCustomer)
}
