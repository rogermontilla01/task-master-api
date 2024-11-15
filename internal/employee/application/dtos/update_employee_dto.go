package dtos

type UpdateEmployeeDto struct {
	ID             *string   `json:"id,omitempty"`
	Name           *string   `json:"name,omitempty"`
	Skills         *[]string `json:"skills,omitempty"`
	AvailableHours *string   `json:"hours,omitempty"`
}
