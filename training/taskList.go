package main

import (
	"bufio"
	"fmt"
	"os"
)

type taskList struct {
	name  string
	tasks []*task
}

func NewTaskList(name string, t *task) *taskList {
	return &taskList{
		name: name,
		tasks: []*task{
			t,
		},
	}
}
func (tl *taskList) printTaskList() {
	fmt.Println("task list name: %s", tl.name)
	for i := 0; i < len(tl.tasks); i++ {
		fmt.Println("TASK:", i)
		fmt.Println(tl.tasks[i])
	}
}
func (t *taskList) appendTask(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) removeTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

type task struct {
	name        string
	description string
	completed   bool
}

func NewTask(name string, description string) *task {
	return &task{
		name:        name,
		description: description,
		completed:   false,
	}
}
func (t *task) String() string {
	return fmt.Sprintf("\n name: %s, \n description: %s, \n completed: %t", t.name, t.description, t.completed)
}
func (t *task) markAsCompleted() {
	t.completed = true
}

func (t *task) updateName(name string) {
	t.name = name
}

func (t *task) updateDescription(description string) {
	t.description = description
}
func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
func inputTask() *task {
	fmt.Println("add task name")
	name := readInput()
	fmt.Println("add task description")
	description := readInput()
	return NewTask(name, description)
}

func inputTaskList() *taskList {
	fmt.Println("add list name")
	name := readInput()
	fmt.Println("add first Taks")
	firstTask := inputTask()
	return NewTaskList(name, firstTask)
}

func main() {
	tl := inputTaskList()
	tl.printTaskList()
}
