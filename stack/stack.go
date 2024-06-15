package stack

import "github.com/kodeyeen/container"

type Stack[T any] interface {
	container.Container
	Peek() (item T, ok bool)
	Push(item T)
	Pop() (item T, ok bool)
}
