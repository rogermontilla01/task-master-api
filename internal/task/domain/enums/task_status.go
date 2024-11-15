package enums

type TaskStatus string

const (
	InProgress TaskStatus = "inProgress"
	Completed  TaskStatus = "completed"
	Canceled   TaskStatus = "canceled"
)
