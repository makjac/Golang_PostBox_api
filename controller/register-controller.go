package controller

import (
	"fmt"
	"net/http"

	"github.com/PBB-api/models"
	"github.com/gin-gonic/gin"
)

func register(ctx *gin.Context) error {
	var Form models.RegisterForm
	err := ctx.ShouldBind(&Form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Opps... Something went wrong with the request",
		})
		return fmt.Errorf("Wrong request")
	}
	return err
}
