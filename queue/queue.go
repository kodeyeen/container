package queue

import "github.com/kodeyeen/container/slice"

type Queue[E comparable] struct {
	elems slice.Slice[E]
}

func New[E comparable](maxLen int) *Queue[E] {
	return &Queue[E]{
		elems: make(slice.Slice[E], 0, maxLen),
	}
}

func (q *Queue[E]) Enqueue(elems ...E) {
	q.elems.Append(elems...)
}

func (q *Queue[E]) Dequeue() (E, bool) {
	if q.elems.Len() == 0 {
		var e E
		return e, false
	}

	elem := q.elems[0]
	q.elems = q.elems[1:]
	return elem, true
}

func (q *Queue[E]) Peek() (E, bool) {
	if q.elems.Len() == 0 {
		var e E
		return e, false
	}

	return q.elems[0], true
}
