package models

type ProductModel struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
}