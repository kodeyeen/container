package list

import "github.com/kodeyeen/container"

type List[E comparable] interface {
	Append(elems ...E)
	Contains(elem E) bool
	container.Container[E]
}
