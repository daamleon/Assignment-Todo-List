package todo_test

import (
	"unit-test-adam/todo"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TodoList", func() {
	var (
		todoList todo.TodoList
	)

	BeforeEach(func() {
		todoList = todo.TodoList{}
	})

	Describe("AddTask", func() {
		It("should add a task to the list", func() {
			todoList.AddTask("Learn Go")
			Expect(todoList.GetTasks()).To(HaveLen(1))
			Expect(todoList.GetTasks()[0].Task).To(Equal("Learn Go"))
		})
	})

	Describe("RemoveTask", func() {
		It("should remove a task from the list", func() {
			todoList.AddTask("Learn Go")
			todoList.RemoveTask(1)
			Expect(todoList.GetTasks()).To(BeEmpty())
		})
	})

	Describe("CompleteTask", func() {
		It("should mark a task as completed", func() {
			todoList.AddTask("Learn Go")
			todoList.CompleteTask(1)
			Expect(todoList.GetTasks()[0].Completed).To(BeTrue())
		})
	})
})
