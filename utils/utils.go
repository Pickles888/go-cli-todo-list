package utils

import (
	"errors"
	"fmt"
	"strings"
)

func Words(s string) []string {
	return strings.Split(s, " ")
}

func Unwords(words []string) string {
	return strings.Join(words[1:], " ")
}

func Map[T, B any](items []T, f func(T) B) []B {
	var arr []B

	for _, item := range items {
		arr = append(arr, f(item))
	}

	return arr
}

func Filter[T any](items []T, f func(T) bool) []T {
	var arr []T

	for _, item := range items {
		if f(item) {
			arr = append(arr, item)
		}
	}

	return arr
}

func Last[T any](items []T) T {
	return items[len(items)-1]
}

func IsEmpty[T any](items []T) bool { return len(items) == 0 }

func Get[T any](items []T, index int) (T, error) {
	var empty T

	if len(items)-1 < index {
		return empty, errors.New(fmt.Sprintf("%v is larger than the length of the array", index))
	}

	return items[index], nil
}
