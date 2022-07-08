package product

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/tusharhow/go-commerce/pkg/db"
	mod "github.com/tusharhow/go-commerce/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	
)

func GetProducts(c *gin.Context) {
	query := bson.D{{
		
	}}
	collection := db.MGI.Db.Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, query)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var results []*mod.Product
	for cur.Next(ctx) {
		var elem mod.Product
		err := cur.Decode(&elem)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		results = append(results, &elem)

	}
	c.JSON(200, results)

}