package middleware

import (
	"api_microservice/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizeToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")

		tokenUser := authHeader[len(SCHEMA):]

		token, err := service.NewJWTService().ValidateToken(tokenUser)

		if token.Vaild {
			_ = token.claims.(jwt.MapClaims)
		} else {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
