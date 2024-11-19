package todo

type Todo struct {
	ID        int
	Task      string
	Completed bool
}

type TodoList struct {
	todos []Todo
}

func (tl *TodoList) AddTask(task string) {
	id := len(tl.todos) + 1
	tl.todos = append(tl.todos, Todo{ID: id, Task: task, Completed: false})
}

func (tl *TodoList) RemoveTask(id int) {
	for i, todo := range tl.todos {
		if todo.ID == id {
			tl.todos = append(tl.todos[:i], tl.todos[i+1:]...)
			break
		}
	}
}

func (tl *TodoList) CompleteTask(id int) {
	for i, todo := range tl.todos {
		if todo.ID == id {
			tl.todos[i].Completed = true
			break
		}
	}
}

func (tl *TodoList) GetTasks() []Todo {
	return tl.todos
}
