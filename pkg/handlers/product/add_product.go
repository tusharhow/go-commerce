package product

import (
	"context"
	"time"
	db "github.com/tusharhow/go-commerce/pkg/db"
	mod "github.com/tusharhow/go-commerce/pkg/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddProduct(c *gin.Context) {
	var product mod.Product
	c.BindJSON(&product)
	collection := db.MGI.Db.Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	product.ID = primitive.NewObjectID().Hex()
	_, err := collection.InsertOne(ctx, product)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "status": "error"})
		return
	}
	if product.Title == "" || product.Price == 0 || product.Description == "" {
		c.JSON(400, gin.H{"error": "Invalid Product", "status": "error"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "message": "Product added successfully", "product": product})

}