package lib

import (
	"errors"
	"fmt"
	"strconv"
	"todo-list/utils"
)

func HandleInput(todoItems []TodoItem, input string) ([]TodoItem, error) {
	words := utils.Words(input)

	id, getIdFailed := utils.Get(words, 1)
	command, getCommandFailed := utils.Get(words, 0)

	if getCommandFailed != nil {
		return todoItems, errors.New("Failed to get a command")
	}

	parsedId, convertIdFailed := strconv.Atoi(id)

	if command == "add" {
		return Add(todoItems, utils.Unwords(words[0:])), nil
	}

	if getIdFailed != nil || convertIdFailed != nil {
		return todoItems, errors.New("Failed to get an ID")
	}

	switch command {
	case "remove":
		return utils.Filter(todoItems, func(a TodoItem) bool { return a.id != parsedId }), nil
	case "complete":
		return MapWithId(todoItems, parsedId, func(a TodoItem) TodoItem {
			return TodoItem{
				completed: !a.completed,
				item:      a.item,
				id:        a.id,
			}
		}), nil
	case "edit":
		return MapWithId(todoItems, parsedId, func(a TodoItem) TodoItem {
			return TodoItem{
				completed: a.completed,
				item:      utils.Unwords(words[2:]),
				id:        a.id,
			}
		}), nil
	}

	return todoItems, errors.New("Invalid Command" + command)
}

func GetInput() string {
	var i string

	fmt.Scanln(&i)

	return i
}
