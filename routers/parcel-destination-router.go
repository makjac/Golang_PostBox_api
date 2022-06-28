package routers

import (
	"github.com/PBB-api/controller"
	"github.com/gin-gonic/gin"
)

func ParcelDestinationRouter(router *gin.Engine) {
	router.GET("/parcel/destination", controller.ParcelDestination)
}
