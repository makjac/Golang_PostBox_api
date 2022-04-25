package controller

import (
	"log"
	"net/http"

	"github.com/PBB-api/models"
	"github.com/gin-gonic/gin"
)

//GET All person table records
func GetAllPersons(c *gin.Context) {
	var Person []models.Person
	err := dbConnect.Model(&Person).Select()

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
		"message": "all persons",
		"data":    Person,
	})
	return
}
