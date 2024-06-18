package arraylist

import (
	"iter"
	"slices"
)

type ArrayList[E comparable] []E

func New[E comparable](length, capacity int) ArrayList[E] {
	return make([]E, length, capacity)
}

func (l *ArrayList[E]) Append(elems ...E) {
	*l = append(*l, elems...)
}

func (l *ArrayList[E]) Insert(i int, elems ...E) {
	*l = slices.Insert(*l, i, elems...)
}

func (l ArrayList[E]) Map(mapping func(E) E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, elem := range l {
			if !yield(mapping(elem)) {
				return
			}
		}
	}
}

func (l ArrayList[E]) Contains(elem E) bool {
	return slices.Contains(l, elem)
}

func (l *ArrayList[E]) Clear() {
	*l = make(ArrayList[E], 0)
}

func (l ArrayList[E]) Len() int {
	return len(l)
}
