package controller

import (
	"log"
	"net/http"

	"github.com/PBB-api/models"
	"github.com/gin-gonic/gin"
)

func GetParcelHistory(c *gin.Context) {
	var historyID = "id = " + c.Param("id")
	var Parcel_history []models.Parcel_history
	err := dbConnect.Model(&Parcel_history).Where(historyID).Select()

	if err != nil {
		log.Printf("Error getting parcel history: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Opps... Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"staus":   http.StatusOK,
		"message": "History successfully",
		"data":    Parcel_history,
	})
	return
}
