package stack

import "github.com/kodeyeen/container"

type Stack[T any] interface {
	container.Container
	Push(item T)
	Pop() (item T, ok bool)
	Peek() (item T, ok bool)
}
