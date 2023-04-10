package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

type Product struct{
	Id int 				`json:"id"` 
	Name string 		`json:"name"`
	Quantity int 		`json:"quantity"`
	CodeValue string 	`json:"code_value"`
	IsPublished bool 	`json:"is_published"`
	Expiration string 	`json:"expiration"`
	Price float64 		`json:"price"`

}

type Request struct{
	Id int 				`json:"id"` 
	Name string `json:"name"`
	Quantity int `json:"quantity"`
	CodeValue string `json:"code_value"`
	IsPublished bool `json:"is_published"`
	Expiration string `json:"expiration"`
	Price float64 `json:"price"`
}

var products  []Request
var lastId int


func main(){
	server:= gin.Default()



	server.POST("/create", func (c *gin.Context)  {
		var request Request

		if err := c.ShouldBindJSON(&request); err != nil{
			c.JSON(400, gin.H{
				"err":err.Error(),
			})
			return
		}
		fmt.Println(request)
		
		request.Id=4
		products = append(products, request)
		c.JSON(201,request)
	})

	server.GET("/products", func (c *gin.Context)  {
		c.JSON(200, &products)
	})

	server.Run()
}


