package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssignmentEntity struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TaskID     primitive.ObjectID `bson:"taskId" json:"taskId"`
	EmployeeID primitive.ObjectID `bson:"employeeId" json:"employeeId"`
	Duration   string             `bson:"duration" json:"duration"`
	CreatedAt  time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt  time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	DeletedAt  time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
}
