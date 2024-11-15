package dtos

type UpdateEmployeeDto struct {
	ID             *string   `json:"id,omitempty"`
	Name           *string   `json:"name,omitempty"`
	Skills         *[]string `json:"skills,omitempty"`
	AvailableHours *float64  `json:"hours,omitempty"`
	AvailableDays  *float64  `json:"days,omitempty"`
}
