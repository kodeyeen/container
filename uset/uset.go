package uset

import (
	"iter"
)

type USet[E comparable] struct {
	elems map[E]struct{}
}

func New[E comparable](capacity int) *USet[E] {
	return &USet[E]{
		elems: make(map[E]struct{}, capacity),
	}
}

func (s *USet[E]) Add(elems ...E) {
	for _, elem := range elems {
		s.elems[elem] = struct{}{}
	}
}

func (s *USet[E]) Remove(elems ...E) {
	for _, elem := range elems {
		delete(s.elems, elem)
	}
}

func (s *USet[E]) Contains(elem E) bool {
	_, ok := s.elems[elem]
	return ok
}

func (s *USet[E]) Union(other *USet[E]) *USet[E] {
	result := New[E](s.Len() + other.Len())

	for elem := range s.elems {
		result.Add(elem)
	}

	for elem := range other.elems {
		result.Add(elem)
	}

	return result
}

func (s *USet[E]) Intersection(other *USet[E]) *USet[E] {
	result := New[E](s.Len() + other.Len())

	for elem := range other.elems {
		if s.Contains(elem) {
			result.Add(elem)
		}
	}

	return result
}

func (s *USet[E]) Difference(other *USet[E]) *USet[E] {
	result := New[E](s.Len())

	for elem := range s.elems {
		if !other.Contains(elem) {
			result.Add(elem)
		}
	}

	return result
}

func (s *USet[E]) IsSubSet(other *USet[E]) bool {
	for elem := range s.elems {
		if !other.Contains(elem) {
			return false
		}
	}

	return true
}

func (s *USet[E]) IsSuperSet(other *USet[E]) bool {
	for elem := range other.elems {
		if !s.Contains(elem) {
			return false
		}
	}

	return true
}

func (s *USet[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		for elem := range s.elems {
			if !yield(elem) {
				return
			}
		}
	}
}

func (s *USet[E]) Clear() {
	clear(s.elems)
}

func (s *USet[E]) Len() int {
	return len(s.elems)
}

func (s *USet[E]) String() string {
	return ""
}
