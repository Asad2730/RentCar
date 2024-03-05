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

func BodyTypeRoutes(r *gin.RouterGroup) {
	r.POST("/", handlers.InsertBodyType)
	r.PUT("/:id", handlers.UpdateBodyType)
	r.DELETE("/:id", handlers.DeleteBodyType)
}

func ColorRoutes(r *gin.RouterGroup) {
	r.POST("/", handlers.InsertColor)
	r.PUT("/:id", handlers.UpdateColor)
	r.DELETE("/:id", handlers.DeleteColor)
}

func EngineRoutes(r *gin.RouterGroup) {
	r.POST("/", handlers.InsertEngineType)
	r.PUT("/:id", handlers.UpdateEngineType)
	r.DELETE("/:id", handlers.DeleteEngineType)
}

func ManufacturerRoutes(r *gin.RouterGroup) {
	r.POST("/", handlers.InsertManufacturer)
	r.PUT("/:id", handlers.UpdateManufacturer)
	r.DELETE("/:id", handlers.DeleteManufacturer)
}

func ModelRoutes(r *gin.RouterGroup) {
	r.POST("/", handlers.InsertModel)
	r.PUT("/:id", handlers.UpdateModel)
	r.DELETE("/:id", handlers.DeleteModel)
}
