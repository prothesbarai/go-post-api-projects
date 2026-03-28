package controllers

import (
	"go-post-api-projects/database"
	"go-post-api-projects/logger"
	"go-post-api-projects/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var products []models.ProductModel

func CreateProducts(ginCtx *gin.Context) {
	// >>> Bearer token 
	authHeader := ginCtx.GetHeader("Authorization")
	if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer "{
		ginCtx.JSON(http.StatusUnauthorized,gin.H{"error" : "Unauthorized. Bearer token missing"})
		logger.AppLogger.Error.Println("Unauthorized. Bearer token missing")
		return
	}

	token := authHeader[7:]
	if token != "angkan#prothes#shreyasi" {
		ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Invalid token"})
		logger.AppLogger.Error.Println("Unauthorized. Invalid token")
		return
	}

	/// >>> JSON body parse
	var input struct{
		Name  string  `form:"name" binding:"required"`
		Price float64 `form:"price" binding:"required"`
	}


	if err := ginCtx.ShouldBind(&input); err != nil{
		ginCtx.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
		logger.AppLogger.Error.Println("Input Data Type Problem : ",err)
		return
	}


	addProduct, err := database.DB.Exec("INSERT INTO products (name, price) VALUES (?,?)",input.Name,input.Price)
	if (err != nil) {
		ginCtx.JSON(http.StatusInternalServerError,gin.H{"error" : err.Error()})
		logger.AppLogger.Error.Println("Product Not Insert : ",err)
		return
	}

	id,_:= addProduct.LastInsertId()


	ginCtx.JSON(http.StatusOK,gin.H{
		"message": "Product created successfully",
		"id": id,
	})

}
