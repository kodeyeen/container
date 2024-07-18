package container

import "fmt"

type Container interface {
	Clear()
	Len() int

	fmt.Stringer
}

type Map[K comparable, V any] interface {
	Get(key K) (val V, ok bool)
	Set(key K, val V)
	Delete(key K)

	Container
}

type Set[E comparable] interface {
	Add(elems ...E)
	Remove(elems ...E)
	Contains(elem E) bool

	Container
}

type Queue[E any] interface {
	Enqueue(elems ...E)
	Dequeue() (elem E, ok bool)
	Peek() (elem E, ok bool)

	Container
}

type Stack[E any] interface {
	Push(elems ...E)
	Pop() (elem E, ok bool)
	Peek() (elem E, ok bool)

	Container
}
