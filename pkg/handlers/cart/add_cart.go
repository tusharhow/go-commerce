package cart

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/tusharhow/go-commerce/pkg/db"
	mod "github.com/tusharhow/go-commerce/pkg/models"
)

func AddCart(c *gin.Context) {
	var cart mod.Cart
	c.BindJSON(&cart)
	collection := db.MGI.Db.Collection("carts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, cart)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"status": "success", "message": "Cart added successfully", "cart": cart})
}
