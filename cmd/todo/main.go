package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "github.com/PriyanshSharma1/CLI_TODO_APP/internal/todo"
)

const (
	todoFile = ".todos.json"
)

func Hello() string {
	return "Hello, World"
}
func main() {
	fmt.Println(Hello())
	add := flag.Bool("add", false, "Add a new todo")
	complete := flag.Int("complete", 0, "Mark a todo as completed")
	delete := flag.Int("delete", 0, "Delete a todo")
	list := flag.Bool("list", false, "List all todos")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading todos: %v\n", err)
		os.Exit(1)
	}

	switch {

	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting input: %v\n", err)
			os.Exit(1)
		}
		todos.Add(task)

	case *complete >= 1:
		err := todos.Complete(*complete)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error completing todo: %v\n", err)
			os.Exit(1)
		}

	case *delete >= 1:
		err := todos.Delete(*delete)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting todo: %v\n", err)
			os.Exit(1)
		}

	case *list:
		todos.Print()
	}

	if err := todos.Store(todoFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error storing todos: %v\n", err)
		os.Exit(1)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		return "", errors.New("unable to read input")
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := strings.TrimSpace(scanner.Text())
	if text == "" {
		return "", errors.New("empty todo is not allowed")
	}
	return text, nil
}
