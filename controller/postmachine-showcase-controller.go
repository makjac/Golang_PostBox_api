package controller

import (
	"log"
	"net/http"

	"github.com/PBB-api/models"
	"github.com/gin-gonic/gin"
)

//GET All person table records
func PostmachineShowcase(c *gin.Context) {
	cityHader := c.GetHeader("city")

	var postmachineShowcase []models.PostmachineShowcase
	if cityHader == "" {
		err := dbConnect.Model(&postmachineShowcase).Select()
		if err != nil {
			log.Printf("Error to getting post machine: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Opps... Something went wrong",
			})
			return
		}
	} else {
		err := dbConnect.Model(&postmachineShowcase).Where("mp_city = ?", cityHader).Where("box_type = mała   ").Select()
		if err != nil {
			log.Printf("Error to getting post machine: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Opps... Something went wrong",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"staus":   http.StatusOK,
		"message": "post machine",
		"data":    postmachineShowcase,
	})
	return
}
