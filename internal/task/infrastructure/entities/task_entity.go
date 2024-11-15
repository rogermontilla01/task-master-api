package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskEntity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Duration  string             `bson:"duration" json:"duration"`
	Skills    []string           `bson:"skills" json:"skills"`
	Status    string             `bson:"status" json:"status"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	DeletedAt time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
}
