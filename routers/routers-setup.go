package routers

import (
	"net/http"

	"github.com/PBB-api/controller"
	"github.com/PBB-api/middlewares"
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

func RtrSetup(router *gin.Engine) {

	router.GET("/", welcome)
	router.NoRoute(notFound)
	PersonRouter(router)

	apiRouters := router.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRouters.GET("/histośry/:id", controller.GetParcelHistory)
		apiRouters.GET("/adress/:id", controller.GetAdressID)
	}

	router.GET("/register", func(ctx *gin.Context) {
		ctx.JSON(200, registerController.Register(ctx))
	})

	router.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	router.GET("/activated", func(c *gin.Context) {
		c.HTML(http.StatusOK, "account-activated.html", gin.H{
			"content": "This is an account-activated page...",
		})
	})

}

func login(c *gin.Context) {

}

func welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome-page.html", gin.H{
		"staus":   http.StatusOK,
		"message": "oops... something is wrong",
	})
	return
}

func notFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"staus":   http.StatusNotFound,
		"message": "oops... something is wrong",
	})
	return
}
