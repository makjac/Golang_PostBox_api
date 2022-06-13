package controller

import (
	"fmt"
	"net/http"

	"github.com/PBB-api/models"
	"github.com/PBB-api/service"
	"github.com/gin-gonic/gin"
	mails "github.com/gin-gonic/mail"
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

func buildSql(name string, surname string, phone string, email string, login string, passwd string) string {
	return fmt.Sprintf("SELECT register_form('%s', '%s', '%s', '%s', '%s', '%s')", name, surname, phone, email, login, passwd)
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

	_, err = dbConnect.Exec(buildSql(Form.Name, Form.Surname, Form.Phone, Form.Email, Form.Login, Form.Passwd))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Opps... Something went wrong with the request",
		})
		return fmt.Errorf("Wrong request")
	}

	ctx.JSON(http.StatusOK, gin.H{
		"staus":   http.StatusOK,
		"message": "register",
	})

	mails.VerficationMail()

	return err
}
