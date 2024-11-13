package dtos

import "time"

type EmployeeDto struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Skills         []string  `json:"skills"`
	AvailableHours float64   `json:"hours"`
	AvailableDays  float64   `json:"days"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
	UpdatedAt      time.Time `json:"updatedAt,omitempty"`
	DeletedAt      time.Time `json:"deletedAt,omitempty"`
}
