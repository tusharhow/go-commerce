package product

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/tusharhow/go-commerce/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteProductByID(c *gin.Context) {
	productId := c.Param("id")
	collection := db.MGI.Db.Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"_id": productId}
	defer cancel()
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"status": "success", "message": "Product deleted successfully"})

}