package controller

import (
	"log"
	"net/http"

	"github.com/PBB-api/models"
	"github.com/gin-gonic/gin"
)

func ReceiverParcels(ctx *gin.Context) {
	uuid, _ := ctx.Get("uuid")

	var parcelShowcase []models.ParcelShowcase
	err := dbConnect.Model(&parcelShowcase).Where("receiver_uuid = ?", uuid).Select()

	if err != nil {
		log.Printf("Error to getting parcels: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Opps... Something went wrong",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"staus":   http.StatusOK,
		"message": "Receiver Parcels",
		"data":    parcelShowcase,
	})
	return
}
