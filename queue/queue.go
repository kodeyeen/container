package queue

import "github.com/kodeyeen/container"

type Queue[E any] interface {
	Enqueue(elems ...E)
	Dequeue() (elem E, ok bool)
	Peek() (elem E, ok bool)
	container.Container[E]
}
