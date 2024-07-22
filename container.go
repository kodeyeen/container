package container

import (
	"fmt"
)

type Container[E any] interface {
	Clear()
	Len() int

	fmt.Stringer
}

type Map[K comparable, V any] interface {
	Get(key K) (val V, ok bool)
	Set(key K, val V)
	Delete(key K)

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

type Stack[E any] interface {
	Push(elems ...E)
	Pop() (elem E, ok bool)
	Peek() (elem E, ok bool)

	Container[E]
}
