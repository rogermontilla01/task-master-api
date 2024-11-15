package enums

type TaskStatusError string

const (
	TaskAlreadyInProgress TaskStatusError = "task already in progress"
	TaskAlreadyCompleted  TaskStatusError = "task already completed"
	TaskAlreadyCanceled   TaskStatusError = "task already canceled"
	NoEmployeeAvailable   TaskStatusError = "not employee available"
)
