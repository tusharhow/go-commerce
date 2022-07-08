package product

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/tusharhow/go-commerce/pkg/db"
	mod "github.com/tusharhow/go-commerce/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)


func UpdateProductByID(c *gin.Context) {
	var product mod.Product
	productId := c.Param("id")
	c.BindJSON(&product)
	collection := db.MGI.Db.Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"_id": productId}
	update := bson.D{
		{"$set", bson.D{
			{"title", product.Title},
			{"description", product.Description},
			{"created_at", time.Now()},
		}},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"status": "success", "message": "Product updated successfully", "product": product})

}