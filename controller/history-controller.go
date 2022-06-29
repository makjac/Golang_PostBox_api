package controller

import (
	"log"
	"net/http"

	"github.com/PBB-api/models"
	"github.com/gin-gonic/gin"
)

func GetHistory(ctx *gin.Context) {
	parcelUuidHader := ctx.GetHeader("parcel_uuid")
	var history []models.History
	err := dbConnect.Model(&history).Where("parcel_uuid = ?", parcelUuidHader).Order("sp_id DESC").Select()

	if err != nil {
		log.Printf("Error getting parcel history: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Opps... Something went wrong",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"staus":   http.StatusOK,
		"message": "History successfully",
		"data":    history,
	})
	return
}
