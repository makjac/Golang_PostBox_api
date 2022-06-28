package routers

import (
	"github.com/PBB-api/controller"
	"github.com/gin-gonic/gin"
)

func PostmachineRouter(router *gin.Engine) {
	router.GET("/postmachine", controller.PostmachineShowcase)
}
