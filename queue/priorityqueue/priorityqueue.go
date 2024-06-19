package priorityqueue

import (
	"container/heap"
)

type PriorityQueue[E comparable] struct {
	core *core[E]
}

// New creates a priority queue.
func New[E comparable](comparator Comparator[E]) *PriorityQueue[E] {
	pq := &PriorityQueue[E]{
		core: &core[E]{
			comparator: comparator,
		},
	}

	return pq
}

// Init initializes the queue with the given elements or clears it if no elements are provided. The complexity is O(n).
func (pq *PriorityQueue[E]) Init(elems ...E) {
	pq.core.elems = make([]E, 0, len(elems))
	pq.core.elems = append(pq.core.elems, elems...)
	heap.Init(pq.core)
}

// Enqueue pushes elements to the queue. The complexity is O(nlogn).
func (pq *PriorityQueue[E]) Enqueue(elems ...E) {
	for _, elem := range elems {
		heap.Push(pq.core, elem)
	}
}

// Dequeue pops an element from the queue. The complexity is O(nlogn).
func (pq *PriorityQueue[E]) Dequeue() (E, bool) {
	if len(pq.core.elems) == 0 {
		var e E
		return e, false
	}

	return heap.Pop(pq.core).(E), true
}

// Peek returns the first element without deleting it from the queue. The complexity is O(1).
func (pq *PriorityQueue[E]) Peek() (E, bool) {
	if pq.Len() == 0 {
		var e E
		return e, false
	}

	return pq.core.elems[0], true
}

func (pq *PriorityQueue[E]) Clear() {
	pq.core.elems = make([]E, 0)
}

// Len returns the number of elements in the queue. The complexity is O(1).
func (pq *PriorityQueue[E]) Len() int {
	return len(pq.core.elems)
}

func (pq *PriorityQueue[E]) String() string {
	return ""
}
