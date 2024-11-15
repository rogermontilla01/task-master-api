package dtos

import "time"

type EmployeeDto struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Skills         []string  `json:"skills"`
	AvailableHours string    `json:"hours"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
	UpdatedAt      time.Time `json:"updatedAt,omitempty"`
	DeletedAt      time.Time `json:"deletedAt,omitempty"`
}
