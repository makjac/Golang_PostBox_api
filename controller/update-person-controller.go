package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func buildUpdateSql(usrID string, name string, surname string, phone string, email string) string {
	id, _ := strconv.Atoi(usrID)
	return fmt.Sprintf("SELECT update_person('%d', '%s', '%s', '%s', '%s')", id, name, surname, phone, email)
}

func UpdateShowcase(ctx *gin.Context) {
	IDHader := ctx.GetHeader("usrID")
	nameHader := ctx.GetHeader("name")
	surnameHader := ctx.GetHeader("surname")
	phoneHader := ctx.GetHeader("phone")
	emailHader := ctx.GetHeader("email")

	_, err := dbConnect.Exec(buildUpdateSql(IDHader, nameHader, surnameHader, phoneHader, emailHader))

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
		"message": nameHader,
	})

	return

}
