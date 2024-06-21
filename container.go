package container

import "fmt"

type Container[E any] interface {
	Clear()
	Len() int

	fmt.Stringer
}

type Dict[K comparable, V any] interface {
	Get(key K) (val V, ok bool)
	Set(key K, val V)

	Container[V]
}

type Set[E comparable] interface {
	Add(elems ...E)
	Remove(elems ...E)
	Contains(elem E) bool

	Container[E]
}

type Queue[E any] interface {
	Enqueue(elems ...E)
	Dequeue() (elem E, ok bool)
	Peek() (elem E, ok bool)

	Container[E]
}
