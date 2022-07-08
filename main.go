package main

import (
	"log"

	"github.com/gin-gonic/gin"
	db "github.com/tusharhow/go-commerce/pkg/db"
	cart "github.com/tusharhow/go-commerce/pkg/handlers/cart"
	hand "github.com/tusharhow/go-commerce/pkg/handlers/product"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.New()
	r.POST("addProduct", hand.AddProduct)
	r.GET("getProducts", hand.GetProducts)
	r.PUT("updateProduct/:id", hand.UpdateProductByID)
	r.DELETE("deleteProduct/:id", hand.DeleteProductByID)
	r.POST("addCart", cart.AddCart)
	log.Fatal(r.Run(":8080"))
}
