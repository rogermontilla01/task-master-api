package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskEntity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `json:"title"`
	Duration  string             `json:"duration"`
	Skills    []string           `json:"skills"`
	Completed bool               `json:"completed"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	DeletedAt time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
}