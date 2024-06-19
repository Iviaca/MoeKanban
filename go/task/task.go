package task

type status int

/* Task status enum */
const (
	Todo status = iota
	InProgress
	Done
)

/* define a Task */
type WorkingTask struct {
	Stat            status
	TaskTitle       string
	TaskDescription string
	Percentage      float64
}

func (t WorkingTask) FilterValue() string {
	return t.TaskTitle
}

func (t WorkingTask) Title() string {
	return t.TaskTitle
}

func (t WorkingTask) Description() string {
	return t.TaskDescription
}
