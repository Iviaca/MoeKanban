package main_model

import (
	"MoeKanbanTui/task"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

/*define a main model*/
type mainModel struct {
	lst []list.Model
	err error
}

/*initial main model*/
func (m *mainModel) initialModel(width, height int) {
	lst := make([]list.Model, 3)

	myTask01 := task.WorkingTask{
		Stat:            task.Todo, // 使用 task 包中定义的常量
		TaskTitle:       "Finish Go project",
		TaskDescription: "Complete the Go project for the client.",
		Percentage:      0.25,
	}

	myTask02 := task.WorkingTask{
		Stat:            task.Todo, // 使用 task 包中定义的常量
		TaskTitle:       "Finish Go project",
		TaskDescription: "Complete the Go project for the client.",
		Percentage:      0.25,
	}

	myTask03 := task.WorkingTask{
		Stat:            task.Todo, // 使用 task 包中定义的常量
		TaskTitle:       "Finish Go project",
		TaskDescription: "Complete the Go project for the client.",
		Percentage:      0.25,
	}

	//submit to change
	lst[0] = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	lst[0].Title = "To Do"
	lst[0].SetItems([]list.Item{myTask01, myTask02, myTask03})
}

/*implementation of bubbletea*/
func (m mainModel) Init() tea.Cmd {

}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

}

func (m mainModel) View() string {

}
