package tasklist

import "fmt"

// List ...
type List struct {
	Tasks []*Task
}

// Add ...
func (l *List) Add(t *Task) {
	l.Tasks = append(l.Tasks, t)
}

// Remove ...
func (l *List) Remove(index int) {
	l.Tasks = append(l.Tasks[:index], l.Tasks[index+1:]...)
}

// PrintAll ...
func (l *List) PrintAll() {
	for _, t := range l.Tasks {
		fmt.Println("name:", t.Name)
		fmt.Println("description:", t.Description)
		fmt.Println("completed:", t.Completed)
	}
}

// PrintCompleted ...
func (l *List) PrintCompleted() {
	for _, t := range l.Tasks {
		if t.Completed {
			fmt.Println("name:", t.Name)
			fmt.Println("description:", t.Description)
			fmt.Println("completed:", t.Completed)
		}
	}
}

// Task ...
type Task struct {
	Name        string
	Description string
	Completed   bool
}

func (t *Task) markAsCompleted() {
	t.Completed = true
}

func (t *Task) updateDescription(description string) {
	t.Description = description
}

func (t *Task) updateName(name string) {
	t.Name = name
}
