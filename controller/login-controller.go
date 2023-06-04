package controller

import (
	"api_microservice/service"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserName string `form:"username" json:"username"`
}

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	jwtService service.JWTService
}

func NewLoginController(jwtService service.JWTService) LoginController {
	return &loginController{
		jwtService: jwtService,
	}
}

func (lc *loginController) Login(ctx *gin.Context) string {
	var UserCredentail User
	err := ctx.ShouldBind(&UserCredentail)
	if err != nil {
		return ""
	}
	isAuth := lc.jwtService.GenerateToken(UserCredentail.UserName)
	return isAuth
}
