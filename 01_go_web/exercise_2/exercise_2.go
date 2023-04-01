package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Product struct{
	ID int 	`json:"id"`
	Name string `json:"name"`
	Quantity int `json:"quantity"`
	CodeValue string `json:"code_value"`
	IsPublished bool `json:"is_published"`
	Expiration string `json:"expiration"`
	Price float64 `json:"price"`
}

var products []Product

func readJson(path string) error{
	var err error
	var file *os.File

	file, err = os.Open(path)
	if err != nil{
		return errors.New("file not found")
	}

	defer file.Close()

	byteValue,_ := io.ReadAll(file)

	err = json.Unmarshal(byteValue, &products)
	if err != nil{
		errors.New("not load json")
	}
	return nil
}

func getAll(c *gin.Context)  {
	c.JSON(http.StatusOK, products)
}


func getById(c *gin.Context)  {
	id := c.Param("id")
	IdProduct,err := strconv.Atoi(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
	}
	for _, v := range products {
		if v.ID == IdProduct {
			return c.JSON(http.StatusOK, v)
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not Found"})
}

func getProductsByPrice(c *gin.Context)  {
	price := c.Query("price")
	priceProduct, err:= strconv.ParseFloat(price, 32)
	if err != nil{
		return c.JSON(http.StatusBadRequest,err)
	}

	var filteredProducts []Product
	for _, v := range products {
		if v.Price > priceProduct{
			filteredProducts =	append(filteredProducts, v)
		}
	}

	c.IndentedJSON(http.StatusOK,filteredProducts)

}


func main()  {
	var err error
	err = readJson("./products.json")
	if err != nil{
		panic("file not found")
	}
	router := gin.Default()
	router.GET("/products",getAll)
	router.GET("/products/:id",getById)
	router.GET("/products/search",getProductsByPrice)
	router.GET("/ping", func (c *gin.Context)  {
		c.string(http.StatusOK,"pong")
	})
	if err = router.Run(":8080"); err != nil{
		panic(err)
	}
}