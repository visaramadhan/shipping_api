package router

import (
	"github.com/visaramadhan/shipping_api_visa/api/shiping"
	"github.com/visaramadhan/shipping_api_visa/config"
	"github.com/visaramadhan/shipping_api_visa/manager"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) error {

	infraManager := manager.NewInfraManager(config.Cfg)
	serviceManager := manager.NewRepoManager(infraManager)
	repoManager := manager.NewServiceManager(serviceManager)

	shipingHandler := shiping.NewShipingHandler(repoManager.ShipingService())

	v1 := router.Group("/api/v1")
	{
		eCommerce := v1.Group("/e-commerce")
		{
			shiping := eCommerce.Group("/shiping")
			{
				shiping.POST("/register-shiping", shipingHandler.AddShiping)
				shiping.GET("/list-shiping", shipingHandler.ListShipings)
				shiping.GET("/:id", shipingHandler.GetShipingById)
				shiping.GET("/cost", shipingHandler.CalculateShippingCost)
			}
		}
	}

	return router.Run()

}
