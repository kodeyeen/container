package set

import "github.com/kodeyeen/container"

type Set[E comparable] interface {
	Add(elems ...E)
	Remove(elems ...E)
	Contains(elem E) bool
	container.Container[E]
}
