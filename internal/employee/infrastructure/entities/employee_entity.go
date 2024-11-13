package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CronjobEntity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	DeletedAt time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
}
