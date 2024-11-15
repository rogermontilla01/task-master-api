package dtos

import "time"

type TaskDto struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Duration  string    `json:"duration,omitempty"`
	Skills    []string  `json:"skills"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}
