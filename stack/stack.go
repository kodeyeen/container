package stack

import "github.com/kodeyeen/container"

type Stack[E comparable] interface {
	Push(elems ...E)
	Pop() (elem E, ok bool)
	Peek() (elem E, ok bool)

	container.Container[E]
}
