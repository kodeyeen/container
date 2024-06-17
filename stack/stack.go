package stack

import "github.com/kodeyeen/container"

type Stack[E any] interface {
	container.Container[E]
	Push(elems ...E)
	Pop() (elem E, ok bool)
	Peek() (elem E, ok bool)
}
