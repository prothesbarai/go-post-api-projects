package controllers

import (
	"go-post-api-projects/database"
	"go-post-api-projects/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetProducts(ginCtx *gin.Context){
	rows, err := database.DB.Query("SELECT id,name,price FROM products")
	if(err != nil){
		ginCtx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	defer rows.Close()

	var products []models.ProductModel

	for rows.Next(){
		var p models.ProductModel
		rows.Scan(&p.Id,&p.Name,&p.Price)
		products = append(products, p)
	}

	ginCtx.JSON(http.StatusOK,products)
}