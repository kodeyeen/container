package set

import "github.com/kodeyeen/container"

type Set[E any] interface {
	container.Container[E]
	Add(elems ...E)
	Remove(elems ...E)
	Contains(elem E) bool
}
