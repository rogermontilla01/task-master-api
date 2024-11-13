package dtos

import "time"

type Assignment struct {
	ID            string    `json:"id"`
	TaskID        string    `json:"taskId"`
	EmployeeID    string    `json:"employeeId"`
	Date          time.Time `json:"date"`
	Status        string    `json:"status"`
	AssignedHours float64   `json:"assignedHours"`
}
