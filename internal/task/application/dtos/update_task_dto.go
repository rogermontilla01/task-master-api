package dtos

type UpdateTaskDto struct {
	Title    *string   `json:"title"`
	Duration *string   `json:"duration"`
	Skills   *[]string `json:"skills"`
	Status   *string   `json:"status"`
}
