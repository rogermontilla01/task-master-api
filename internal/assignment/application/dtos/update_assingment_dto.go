package dtos

type UpdateAssignmentDto struct {
	TaskID     *string `json:"taskId"`
	EmployeeID *string `json:"employeeId"`
	Duration   *string `json:"duration"`
}
