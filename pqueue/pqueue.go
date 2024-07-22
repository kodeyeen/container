package pqueue

import (
	"github.com/kodeyeen/container"
	"github.com/kodeyeen/container/heap"
)

type PQueue[E comparable] struct {
	heap *heap.Heap[E]
}

// New creates a priority queue.
func New[E comparable](comparator container.Comparator[E]) *PQueue[E] {
	return &PQueue[E]{
		heap: heap.New(comparator),
	}
}

// Init initializes the queue with the given elements or clears it if no elements are provided. The complexity is O(n).
func (pq *PQueue[E]) Init(elems ...E) {
	pq.heap.Init(elems...)
}

// Enqueue pushes elements to the queue. The complexity is O(nlogn).
func (pq *PQueue[E]) Enqueue(elems ...E) {
	for _, elem := range elems {
		pq.heap.Push(elem)
	}
}

// Dequeue pops an element from the queue. The complexity is O(nlogn).
func (pq *PQueue[E]) Dequeue() (E, bool) {
	return pq.heap.Pop()
}

// Peek returns the first element without deleting it from the queue. The complexity is O(1).
func (pq *PQueue[E]) Peek() (E, bool) {
	return pq.heap.Peek()
}

func (pq *PQueue[E]) Clear() {
	pq.heap.Clear()
}

// Len returns the number of elements in the queue. The complexity is O(1).
func (pq *PQueue[E]) Len() int {
	return pq.heap.Len()
}

func (pq *PQueue[E]) String() string {
	return ""
}
