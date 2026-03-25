package routes

import (
	"go-post-api-projects/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(route *gin.Engine) {
	route.GET("/products",controllers.GetProducts)

	// >>> For Igoner Fav Icons
	route.GET("/favicon.ico",func(ctx *gin.Context) {ctx.Status(204)})
}