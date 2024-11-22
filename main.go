package main

import (
	"fmt"
	"log"
	"strings"
	"todo-list/lib"
	"todo-list/utils"
)

func main() {
	log.SetPrefix("Error: ")
	log.SetFlags(0)

	var todoItems []lib.TodoItem

	todoItems = append(todoItems, lib.Default())

	for {
		input := lib.GetInput()

		todoItems, err := lib.HandleInput(todoItems, input)

		fmt.Println(
			strings.Join(
				utils.Map(todoItems, func(a lib.TodoItem) string { return lib.ShowTodoItem(a) }),
				"\n",
			),
		)

		if err != nil {
			log.Println(err)
		}
	}
}
