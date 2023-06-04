package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"api_microservice/controller"
	"api_microservice/middleware"
	"api_microservice/service"

	"github.com/gin-gonic/gin"
)

var (
	jwtService      service.JWTService         = service.NewJWTService()
	loginController controller.LoginController = controller.NewLoginController(jwtService)
)

func main() {
	router := gin.New()

	// Create reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "http://localhost:5000",
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{"data": "App up and running"})
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

	router.Any("/api/*path", middleware.AuthorizeToken(), func(ctx *gin.Context) {
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	})

	router.Run(":8005")
}
