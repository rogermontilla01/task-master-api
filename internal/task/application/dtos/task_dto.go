package dtos

import "time"

type TaskDto struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Duration time.Time `json:"createdAt,omitempty"`
	Skills   []string  `json:"skills"`
}
