package models

var ProductModels struct{
	Name string `form:"name" binding:"required,min=3,max=100"`
	Price float64 `form:"price" binding:"required,gt=0"`
}