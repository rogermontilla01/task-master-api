package dtos

import "time"

type AssignmentDto struct {
	ID         string    `json:"id"`
	TaskID     string    `json:"taskId"`
	EmployeeID string    `json:"employeeId,omitempty"`
	Duration   string    `json:"duration"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
	DeletedAt  time.Time `json:"deletedAt,omitempty"`
}
