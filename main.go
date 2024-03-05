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

	routeGroup := []struct {
		path   string
		group  *gin.RouterGroup
		routes func(*gin.RouterGroup)
	}{
		{"/car", api.Group("/car"), routes.CartRoutes},
		{"/bodytype", api.Group("/bodytype"), routes.BodyTypeRoutes},
		{"/color", api.Group("/color"), routes.ColorRoutes},
		{"/engine", api.Group("/engine"), routes.EngineRoutes},
		{"/manufacturer", api.Group("/manufacturer"), routes.ManufacturerRoutes},
		{"/model", api.Group("/model"), routes.ModelRoutes},
		{"/driver", api.Group("/driver"), routes.DriverRoutes},
		{"/address", api.Group("/address"), routes.AddressRoutes},
		{"/document", api.Group("/document"), routes.DocumentRoutes},
		{"/documentType", api.Group("/documentType"), routes.DocumentTypeRoutes},
	}

	for _, rg := range routeGroup {
		rg.routes(rg.group)
	}

	r.Run("0.0.0.0:3000")
}
