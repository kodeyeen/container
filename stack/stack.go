package stack

import "github.com/kodeyeen/container"

type Stack[T any] interface {
	container.Container
	Push(items ...T)
	Pop() (item T, ok bool)
	Peek() (item T, ok bool)
}
