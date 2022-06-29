package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func buildcreateParcelSql(receiver_username string, receiver_pm_id string, sender_uuid string, sender_pm_id string, parcel_type string, parcel_name string) string {
	Rid, _ := strconv.Atoi(receiver_pm_id)
	Sid, _ := strconv.Atoi(sender_pm_id)
	PTid, _ := strconv.Atoi(parcel_type)
	return fmt.Sprintf("SELECT create_parcel('%s', '%d', '%s', '%d', '%d', '%s')", receiver_username, Rid, sender_uuid, Sid, PTid, parcel_name)
}

func CreateParcel(ctx *gin.Context) {
	sender_usernameHader := ctx.GetHeader("sender_username")
	receiver_usernameHader := ctx.GetHeader("receiver_username")
	receiver_pm_idHader := ctx.GetHeader("receiver_pm_id")
	sender_pm_idHader := ctx.GetHeader("sender_pm_id")
	parcel_typeHader := ctx.GetHeader("parcel_type")
	parcel_nameHader := ctx.GetHeader("parcel_name")

	_, err := dbConnect.Exec(buildcreateParcelSql(receiver_usernameHader, receiver_pm_idHader, sender_usernameHader, sender_pm_idHader, parcel_typeHader, parcel_nameHader))

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
		"message": "Parcel Created",
	})

	return

}
