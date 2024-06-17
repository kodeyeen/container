package priorityqueue

import (
	"container/heap"
)

type PriorityQueue[T any] struct {
	core *core[T]
}

// New creates a priority queue.
func New[T any](comparator Comparator[T]) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{
		core: &core[T]{
			comparator: comparator,
		},
	}

	return pq
}

// Init initializes the queue with the given elements or clears it if no elements are provided. The complexity is O(n).
func (pq *PriorityQueue[T]) Init(elems ...T) {
	pq.core.elems = make([]T, 0, len(elems))
	pq.core.elems = append(pq.core.elems, elems...)
	heap.Init(pq.core)
}

// Len returns the number of elements in the queue. The complexity is O(1).
func (pq *PriorityQueue[T]) Len() int {
	return len(pq.core.elems)
}

// Enqueue pushes elements to the queue. The complexity is O(nlogn).
func (pq *PriorityQueue[T]) Enqueue(elems ...T) {
	for _, elem := range elems {
		heap.Push(pq.core, elem)
	}
}

// Dequeue pops an element from the queue. The complexity is O(nlogn).
func (pq *PriorityQueue[T]) Dequeue() (T, bool) {
	if pq.Len() == 0 {
		var t T
		return t, false
	}

	return heap.Pop(pq.core).(T), true
}

// Peek returns the first element without deleting it from the queue. The complexity is O(1).
func (pq *PriorityQueue[T]) Peek() (T, bool) {
	if pq.Len() == 0 {
		var t T
		return t, false
	}

	return pq.core.elems[0], true
}
