package vector

import (
	"iter"
	"slices"
)

type Vector[E comparable] []E

func New[E comparable](length, capacity int) Vector[E] {
	return make([]E, length, capacity)
}

func (l *Vector[E]) Prepend(elems ...E) {
	*l = slices.Insert(*l, 0, elems...)
}

func (l *Vector[E]) Append(elems ...E) {
	*l = append(*l, elems...)
}

func (l *Vector[E]) Insert(i int, elems ...E) {
	*l = slices.Insert(*l, i, elems...)
}

func (l Vector[E]) Map(mapping func(E) E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, elem := range l {
			if !yield(mapping(elem)) {
				return
			}
		}
	}
}

func (l Vector[E]) Contains(elem E) bool {
	return slices.Contains(l, elem)
}

func (l Vector[E]) ContainsFunc(elem E, f func(E) bool) bool {
	return slices.ContainsFunc(l, f)
}

func (l *Vector[E]) Delete(i, j int) {
	*l = slices.Delete(*l, i, j)
}

func (l *Vector[E]) DeleteFunc(del func(E) bool) {
	*l = slices.DeleteFunc(*l, del)
}

func (l *Vector[E]) Clear() {
	*l = make(Vector[E], 0)
}

func (l Vector[E]) Len() int {
	return len(l)
}
