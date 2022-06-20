package controller

import (
	"log"
	"net/http"

	"github.com/PBB-api/models"
	"github.com/gin-gonic/gin"
)

//GET All person table records
func GetShowcase(c *gin.Context) {
	uuid, _ := c.Get("uuid")

	var Showcase []models.Showcase
	err := dbConnect.Model(&Showcase).Where("user_uuid = ?", uuid).Select()

	if err != nil {
		log.Printf("Error to getting persons: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Opps... Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"staus":   http.StatusOK,
		"message": "Showcase",
		"data":    Showcase,
	})
	return
}
