package routes

import (
	"github.com/Asad2730/RentCar/handlers"
	"github.com/gin-gonic/gin"
)

func DriverRoutes(r *gin.RouterGroup) {
	r.GET("/", handlers.GetDriver)
	r.POST("/", handlers.InsertDriver)
	r.GET("/:id", handlers.GetDriver)
	r.PUT("/:id", handlers.UpdateDriver)
	r.DELETE("/:id", handlers.DeleteDriver)
}

func AddressRoutes(r *gin.RouterGroup) {
	r.POST("/", handlers.InsertAddress)
	r.PUT("/:id", handlers.UpdateAddress)
	r.DELETE("/:id", handlers.DeleteAddress)
}

func DocumentRoutes(r *gin.RouterGroup) {
	r.POST("/", handlers.InsertDocument)
	r.PUT("/:id", handlers.UpdateDocument)
	r.DELETE("/:id", handlers.DeleteDocument)
}

func DocumentTypeRoutes(r *gin.RouterGroup) {
	r.GET("/", handlers.GetDocumentsType)
	r.POST("/", handlers.InsertDocumentType)
	r.PUT("/:id", handlers.UpdateDocumentType)
	r.DELETE("/:id", handlers.DeleteDocumentType)
}
