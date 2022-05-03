package controller

import (
	"log"
	"net/http"

	"github.com/PBB-api/models"
	"github.com/gin-gonic/gin"
)

func GetAdressID(c *gin.Context) {
	var Adress []models.Adress
	err := dbConnect.Model(&Adress).Where("id = ?", c.Param("id")).Select()

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
		"message": "adress",
		"data":    Adress,
	})
	return
}
