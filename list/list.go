package list

import "github.com/kodeyeen/container"

type List[E comparable] interface {
	Prepend(elems ...E)
	Append(elems ...E)
	Contains(elem E) bool

	container.Container[E]
}
