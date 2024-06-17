package queue

import "github.com/kodeyeen/container"

type Queue[T any] interface {
	container.Container[T]
	Enqueue(items ...T)
	Dequeue() (item T, ok bool)
	Peek() (item T, ok bool)
}
