package main

import (
	"sync"
)

// Todo represents our todo
type Todo struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	IsComplete bool   `json:"isDone"`
}

type TodoManager struct {
	todos []Todo
	m     sync.Mutex // we want all our operations to be atomic
}

func NewTodoManager() TodoManager {
	return TodoManager{
		todos: make([]Todo, 0),
		m:     sync.Mutex{},
	}
}


func (tm *TodoManager) GetAll() []Todo {
	return tm.todos
}