package priorityqueue

import (
	"container/heap"
)

type PriorityQueue[T any] struct {
	core *core[T]
}

// New creates a queue.
func New[T any](comparator Comparator[T]) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{
		core: &core[T]{
			comparator: comparator,
		},
	}

	return pq
}

// Init initializes the queue with the given items or clears it if no items provided. The complexity is O(n).
func (pq *PriorityQueue[T]) Init(items ...T) {
	pq.core.items = make([]T, 0, len(items))
	pq.core.items = append(pq.core.items, items...)
	heap.Init(pq.core)
}

// Enqueue pushes items to the queue. The complexity is O(nlogn).
func (pq *PriorityQueue[T]) Enqueue(items ...T) {
	for _, item := range items {
		heap.Push(pq.core, item)
	}
}

// Dequeue pops an item from the queue. The complexity is O(nlogn).
func (pq *PriorityQueue[T]) Dequeue() T {
	return heap.Pop(pq.core).(T)
}

// Peek returns the first item without deleting it from the queue. The complexity is O(1).
func (pq *PriorityQueue[T]) Peek() T {
	return pq.core.items[0]
}

// Len returns the number of items in the queue. The complexity is O(1).
func (pq *PriorityQueue[T]) Len() int {
	return len(pq.core.items)
}
