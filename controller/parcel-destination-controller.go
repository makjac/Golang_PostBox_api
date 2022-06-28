package controller

import (
	"log"
	"net/http"

	"github.com/PBB-api/models"
	"github.com/gin-gonic/gin"
)

//GET All person table records
func ParcelDestination(ctx *gin.Context) {
	parcelUuidHeader := ctx.GetHeader("puuid")

	var parcelDestination []models.ParcelDestination
	err := dbConnect.Model(&parcelDestination).Where("parcel_uuid = ?", parcelUuidHeader).Select()

	if err != nil {
		log.Printf("Error to getting persons: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Opps... Something went wrong",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"staus":   http.StatusOK,
		"message": "parcel destination",
		"data":    parcelDestination,
	})
	return
}
