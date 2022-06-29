package routers

import (
	"github.com/PBB-api/controller"
	"github.com/gin-gonic/gin"
)

func HistoryRouter(router *gin.Engine) {
	router.GET("/parcel/history", controller.GetHistory)
}
