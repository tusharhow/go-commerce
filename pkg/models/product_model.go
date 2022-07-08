package models

import "time"

type Product struct {
	ID          string    `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string    `json:"title"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
