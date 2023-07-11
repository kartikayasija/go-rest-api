package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	Title     string              `json:"title" bson:"title" validate:"required"`
	Content   string              `json:"content" bson:"content" validate:"required"`
	CreatedAt primitive.Timestamp `json:"created_at" bson:"createdAt"`
	UpdatedAt primitive.Timestamp `json:"updated_at" bson:"updatedAt"`
}
