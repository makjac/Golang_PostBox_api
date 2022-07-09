package routers

import (
	"net/http"

	"github.com/PBB-api/controller"
	"github.com/PBB-api/service"
	"github.com/gin-gonic/gin"
)

var (
	//loginService service.LoginService = service.NewLoginService()
	jwtService service.JWTService = service.NewJWTService()

	loginController    controller.LoginController    = controller.NewLoginController(jwtService)
	jwtRegisterService service.JWTRegisterService    = service.NewJWTRegisterService()
	registerController controller.RegisterController = controller.NewRegisterController(jwtRegisterService)
)

func authRouters(router *gin.Engine) {
	router.POST("/register", func(ctx *gin.Context) {
		ctx.JSON(200, registerController.Register(ctx))
	})

	router.GET("/activate/:token", func(ctx *gin.Context) {
		ctx.JSON(200, registerController.Activate(ctx))
	})

	router.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"staus":   http.StatusOK,
				"message": "logged in",
				"token":   token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"staus":   http.StatusUnauthorized,
				"message": "Invalid login or password",
				"token":   "",
			})
		}
	})
}
