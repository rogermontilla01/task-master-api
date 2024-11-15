package dtos

type UpdateTaskDto struct {
	Title     *string   `json:"title"`
	Duration  *string   `json:"duration"`
	Skills    *[]string `json:"skills"`
	Completed *bool     `json:"completed"`
}
