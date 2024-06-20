package arraylist

import (
	"iter"
	"slices"
)

type ArrayList[E comparable] []E

func New[E comparable](length, capacity int) ArrayList[E] {
	return make([]E, length, capacity)
}

func (l *ArrayList[E]) Prepend(elems ...E) {
	*l = slices.Insert(*l, 0, elems...)
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

func (l ArrayList[E]) ContainsFunc(elem E, f func(E) bool) bool {
	return slices.ContainsFunc(l, f)
}

func (l *ArrayList[E]) Delete(i, j int) {
	*l = slices.Delete(*l, i, j)
}

func (l *ArrayList[E]) DeleteFunc(del func(E) bool) {
	*l = slices.DeleteFunc(*l, del)
}

func (l *ArrayList[E]) Clear() {
	*l = make(ArrayList[E], 0)
}

func (l ArrayList[E]) Len() int {
	return len(l)
}
