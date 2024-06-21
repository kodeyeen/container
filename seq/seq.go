package list

import "github.com/kodeyeen/container"

type Seq[E comparable] interface {
	Prepend(elems ...E)
	Append(elems ...E)
	Contains(elem E) bool

	container.Container[E]
}
