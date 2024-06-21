package slice

import (
	"iter"
	"slices"
)

type Slice[E comparable] []E

func New[E comparable](length, capacity int) Slice[E] {
	return make([]E, length, capacity)
}

func (s *Slice[E]) Prepend(elems ...E) {
	*s = slices.Insert(*s, 0, elems...)
}

func (s *Slice[E]) Append(elems ...E) {
	*s = append(*s, elems...)
}

func (s *Slice[E]) Insert(i int, elems ...E) {
	*s = slices.Insert(*s, i, elems...)
}

func (s Slice[E]) Map(mapping func(E) E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, elem := range s {
			if !yield(mapping(elem)) {
				return
			}
		}
	}
}

func (s Slice[E]) Contains(elem E) bool {
	return slices.Contains(s, elem)
}

func (s Slice[E]) ContainsFunc(elem E, f func(E) bool) bool {
	return slices.ContainsFunc(s, f)
}

func (s *Slice[E]) Delete(i, j int) {
	*s = slices.Delete(*s, i, j)
}

func (s *Slice[E]) DeleteFunc(del func(E) bool) {
	*s = slices.DeleteFunc(*s, del)
}

func (s *Slice[E]) Clear() {
	*s = make(Slice[E], 0)
}

func (s Slice[E]) Len() int {
	return len(s)
}
