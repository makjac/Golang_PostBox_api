package controller

import (
	"github.com/PBB-api/dto"
	"github.com/PBB-api/service"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	//loginService service.LoginService
	jWtService service.JWTService
}

func NewLoginController( //loginService service.LoginService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		//loginService: loginService,
		jWtService: jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return ""
	}

	isAuthenticated, uuid, err := CheckUser(credentials.Username, credentials.Passwd)
	//fmt.Println(err)
	//isAuthenticated := controller.loginService.Login(credentials.Username, credentials.Passwd)
	if isAuthenticated {
		return controller.jWtService.GenerateToken(credentials.Username, uuid)
	}
	return ""
}
