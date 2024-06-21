package set

import (
	"iter"
)

type Set[E comparable] struct {
	elems map[E]struct{}
}

func New[E comparable](capacity int) *Set[E] {
	return &Set[E]{
		elems: make(map[E]struct{}, capacity),
	}
}

func (s *Set[E]) Add(elems ...E) {
	for _, elem := range elems {
		s.elems[elem] = struct{}{}
	}
}

func (s *Set[E]) Remove(elems ...E) {
	for _, elem := range elems {
		delete(s.elems, elem)
	}
}

func (s *Set[E]) Contains(elem E) bool {
	_, ok := s.elems[elem]
	return ok
}

func (s *Set[E]) Union(other *Set[E]) *Set[E] {
	result := New[E](s.Len() + other.Len())

	for elem := range s.elems {
		result.Add(elem)
	}

	for elem := range other.elems {
		result.Add(elem)
	}

	return result
}

func (s *Set[E]) Intersection(other *Set[E]) *Set[E] {
	result := New[E](s.Len() + other.Len())

	for elem := range other.elems {
		if s.Contains(elem) {
			result.Add(elem)
		}
	}

	return result
}

func (s *Set[E]) Difference(other *Set[E]) *Set[E] {
	result := New[E](s.Len())

	for elem := range s.elems {
		if !other.Contains(elem) {
			result.Add(elem)
		}
	}

	return result
}

func (s *Set[E]) IsSubSet(other *Set[E]) bool {
	for elem := range s.elems {
		if !other.Contains(elem) {
			return false
		}
	}

	return true
}

func (s *Set[E]) IsSuperSet(other *Set[E]) bool {
	for elem := range other.elems {
		if !s.Contains(elem) {
			return false
		}
	}

	return true
}

func (s *Set[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		for elem := range s.elems {
			if !yield(elem) {
				return
			}
		}
	}
}

func (s *Set[E]) Clear() {
	clear(s.elems)
}

func (s *Set[E]) Len() int {
	return len(s.elems)
}

func (s *Set[E]) String() string {
	return ""
}
