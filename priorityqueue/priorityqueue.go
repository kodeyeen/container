package priorityqueue

import (
	"github.com/kodeyeen/container/heap"
	"github.com/kodeyeen/container/internal/cmp"
)

type PriorityQueue[E comparable] struct {
	heap *heap.Heap[E]
}

// New creates a priority queue.
func New[E comparable](comparator cmp.Comparator[E]) *PriorityQueue[E] {
	return &PriorityQueue[E]{
		heap: heap.New(comparator),
	}
}

// Init initializes the queue with the given elements or clears it if no elements are provided. The complexity is O(n).
func (pq *PriorityQueue[E]) Init(elems ...E) {
	pq.heap.Init(elems...)
}

// Enqueue pushes elements to the queue. The complexity is O(nlogn).
func (pq *PriorityQueue[E]) Enqueue(elems ...E) {
	for _, elem := range elems {
		pq.heap.Push(elem)
	}
}

// Dequeue pops an element from the queue. The complexity is O(nlogn).
func (pq *PriorityQueue[E]) Dequeue() (E, bool) {
	return pq.heap.Pop()
}

// Peek returns the first element without deleting it from the queue. The complexity is O(1).
func (pq *PriorityQueue[E]) Peek() (E, bool) {
	return pq.heap.Peek()
}

func (pq *PriorityQueue[E]) Clear() {
	pq.heap.Clear()
}

// Len returns the number of elements in the queue. The complexity is O(1).
func (pq *PriorityQueue[E]) Len() int {
	return pq.heap.Len()
}

func (pq *PriorityQueue[E]) String() string {
	return ""
}
