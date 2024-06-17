package queue

import "github.com/kodeyeen/container"

type Queue[T any] interface {
	container.Container[T]
	Enqueue(elems ...T)
	Dequeue() (elem T, ok bool)
	Peek() (elem T, ok bool)
}
