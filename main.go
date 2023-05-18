package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOk, gin.H{"data": "App up and running"})
	})

	
}