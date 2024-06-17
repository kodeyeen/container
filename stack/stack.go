package stack

import "github.com/kodeyeen/container"

type Stack[T any] interface {
	container.Container[T]
	Push(elems ...T)
	Pop() (elem T, ok bool)
	Peek() (elem T, ok bool)
}
