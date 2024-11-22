package lib

import (
	"errors"
	"fmt"
	"todo-list/utils"
)

type TodoItem struct {
	completed bool
	item      string
	id        int
}

func MapTodo(item TodoItem, f func(TodoItem) TodoItem) TodoItem {
	return f(item)
}

func MapTodoInner(item TodoItem, f func(string) string) TodoItem {
	return MapTodo(item, func(a TodoItem) TodoItem {
		return TodoItem{
			completed: a.completed,
			item:      f(a.item),
			id:        a.id,
		}
	})
}

func ToggleCompleted(item TodoItem) TodoItem {
	return MapTodo(item, func(a TodoItem) TodoItem {
		return TodoItem{
			completed: !a.completed,
			item:      a.item,
			id:        a.id,
		}
	})
}

func Add(items []TodoItem, name string) []TodoItem {
	var newId int

	if utils.IsEmpty(items) {
		newId = 0
	} else {
		newId = utils.Last(items).id + 1
	}

	return append(items, TodoItem{
		completed: false,
		item:      name,
		id:        newId,
	})
}

func Search(items []TodoItem, id int) ([]TodoItem, error) {
	matching := utils.Filter(items, func(a TodoItem) bool { return a.id == id })

	if utils.IsEmpty(matching) {
		return matching, errors.New(fmt.Sprintf("No results for search: %v", id))
	}

	return matching, nil
}

func MapWithId(items []TodoItem, id int, f func(TodoItem) TodoItem) []TodoItem {
	return utils.Map(items, func(a TodoItem) TodoItem {
		if a.id == id {
			return f(a)
		}

		return a
	})
}

func ShowTodoItem(a TodoItem) string {
	var completedStr string

	if a.completed {
		completedStr = "[x]"
	} else {
		completedStr = "[ ]"
	}

	return fmt.Sprintf("%v %v | %v", a.id, completedStr, a.item)
}

func Default() TodoItem {
	return TodoItem{
		completed: true,
		item:      "default item",
		id:        0,
	}
}
