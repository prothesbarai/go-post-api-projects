package main

import (
	"go-post-api-projects/database"
	"go-post-api-projects/routes"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.ConnectDB()
	routes.SetRoutes(router)
	port := os.Getenv("PORT")
	if(port == ""){port = "8080"}
	router.Run(":"+port)
}