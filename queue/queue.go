package queue

import "github.com/kodeyeen/container"

type Queue[E any] interface {
	container.Container[E]
	Enqueue(elems ...E)
	Dequeue() (elem E, ok bool)
	Peek() (elem E, ok bool)
}
