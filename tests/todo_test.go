package todo_test

import (
	"testing"

	"github.com/PriyanshSharma1/CLI_TODO_APP/internal/todo"
)

func TestAdd(t *testing.T) {
	todos := make(todo.Todos, 0)

	todos.Add("Test Task 1")
	todos.Add("Test Task 2")

	if len(todos) != 2 {
		t.Errorf("Expected todos length to be 2, got %d", len(todos))
	}
}

func TestComplete(t *testing.T) {
	todos := todo.Todos{
		{Task: "Task 1", Done: false},
		{Task: "Task 2", Done: false},
	}

	err := todos.Complete(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !todos[0].Done {
		t.Errorf("Expected first task to be marked as done")
	}
}

func TestDelete(t *testing.T) {
	todos := todo.Todos{
		{Task: "Task 1", Done: false},
		{Task: "Task 2", Done: false},
	}

	err := todos.Delete(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(todos) != 1 {
		t.Errorf("Expected todos length to be 1 after deletion, got %d", len(todos))
	}
}
