package main

import (
	"go-post-api-projects/database"
	"go-post-api-projects/logger"
	"go-post-api-projects/routes"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init(){
	err := godotenv.Load()
	if(err != nil){
		logger.AppLogger.Error.Println("Couldn't Load env file : ",err)
		os.Exit(1)
	}
}

func main() {
	router := gin.Default()
	database.ConnectDB()
	routes.SetRoutes(router)
	port := os.Getenv("PORT")
	if(port == ""){port = "8080"}
	router.Run(":"+port)
}