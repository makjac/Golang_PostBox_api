package controller

import (
	"fmt"
	"net/http"

	"github.com/PBB-api/models"
	"github.com/PBB-api/service"
	"github.com/gin-gonic/gin"
)

type RegisterController interface {
	Register(ctx *gin.Context) error
}

type registerController struct {
	JWTregisterService service.JWTRegisterService
}

func NewRegisterController(JWTregisterService service.JWTRegisterService) RegisterController {
	return &registerController{
		JWTregisterService: JWTregisterService,
	}
}

func (controller *registerController) Register(ctx *gin.Context) error {
	var Form models.RegisterForm
	err := ctx.ShouldBind(&Form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Opps... Something went wrong with the request",
		})
		return fmt.Errorf("Wrong request")
	}

	Form.Passwd, err = service.HashPasswd(Form.Passwd)
	if err != nil {
		return fmt.Errorf("Password hash error")
	}

	//err = dbConnect.Model(&Form).Select()
	err = dbConnect.Model(nil).Column("register_form('Kokon', 'Kiszko', '123456789', 'kiki@op.pl', 'koki, 'iop098')").Select()
	//_, err = dbConnect.DB().Query(nil, "register_form('Kokon', 'Kiszko', '123456789', 'kiki@op.pl', 'koki, 'iop098')")
	//err = dbConnect.QueryRow(ctx, "SELECT register_form('Kokon', 'Kiszko', '123456789', 'kiki@op.pl', 'koki, 'iop098')")
	//err = dbConnect.Model(nil).Select()

	return err
}
