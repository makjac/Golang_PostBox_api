package routers

import (
	"github.com/PBB-api/controller"
	"github.com/PBB-api/middlewares"
	"github.com/gin-gonic/gin"
)

func jwtRouters(router *gin.Engine) {
	jwtRouters := router.Group("/api", middlewares.AuthorizeJWT())
	{
		jwtRouters.GET("/history/:id", controller.GetParcelHistory)
		jwtRouters.GET("/adress/:id", controller.GetAdressID)
		jwtRouters.GET("/showcase", controller.GetShowcase)
		jwtRouters.POST("/showcase/update", controller.UpdateShowcase)
		jwtRouters.GET("/parcels/send", controller.SenderParcels)
		jwtRouters.GET("/parcels/incoming", controller.ReceiverParcels)
		jwtRouters.POST("/parcels/create", controller.CreateParcel)
	}
}
