package routes

import (
	"go-post-api-projects/controllers"
	"github.com/gin-gonic/gin"
)

func SetRoutes(routers *gin.Engine) {
	/// >>> For FavIcon
	routers.GET("/favicon.ico",func(ctx *gin.Context) {ctx.Status(204)})
	routers.GET("/products",controllers.GetProducts)
}