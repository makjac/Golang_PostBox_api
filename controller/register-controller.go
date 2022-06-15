package controller

import (
	"fmt"
	"net/http"

	mails "github.com/PBB-api/mail"
	"github.com/PBB-api/models"
	"github.com/PBB-api/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type RegisterController interface {
	Register(ctx *gin.Context) error
	Activate(ctx *gin.Context) error
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

	var token string
	token, err = controller.JWTregisterService.GenerateToken(Form.Login)

	if err != nil {
	}

	mails.VerficationMail(Form.Email, Form.Name, "http://makjac.pl:8080/activate/"+token)

	return nil
}

func (controller *registerController) Activate(ctx *gin.Context) error {
	var tokenString = ctx.Param("token")

	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Wrong token",
		})
		return fmt.Errorf("Wrong token")
	}

	token, _ := service.NewJWTRegisterService().ValidateToken(tokenString)

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)

		auth := models.Auth{
			Username: claims["login"].(string),
			Active:   true,
		}

		_, err := dbConnect.Model(&auth).Set("active = ?active").Where("username = ?username").Update()

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusConflict,
				"message": "Unable to apdate the record",
			})
			return fmt.Errorf("Unable to apdate the record")
		}

		ctx.JSON(http.StatusOK, gin.H{
			"staus":   http.StatusOK,
			"message": "Active",
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Wrong token",
		})
		return fmt.Errorf("Wrong token")
	}

	return nil
}
