package hashset

import "iter"

type HashSet[E comparable] struct {
	elems map[E]struct{}
}

func New[E comparable](capacity int) *HashSet[E] {
	return &HashSet[E]{
		elems: make(map[E]struct{}, capacity),
	}
}

func (s *HashSet[E]) Add(elems ...E) {
	for _, elem := range elems {
		s.elems[elem] = struct{}{}
	}
}

func (s *HashSet[E]) Remove(elems ...E) {
	for _, elem := range elems {
		delete(s.elems, elem)
	}
}

func (s *HashSet[E]) Contains(elem E) bool {
	_, ok := s.elems[elem]
	return ok
}

func (s *HashSet[E]) Union(other *HashSet[E]) *HashSet[E] {
	result := New[E](s.Len() + other.Len())

	for elem := range s.elems {
		result.Add(elem)
	}

	for elem := range other.elems {
		result.Add(elem)
	}

	return result
}

func (s *HashSet[E]) Intersection(other *HashSet[E]) *HashSet[E] {
	result := New[E](s.Len() + other.Len())

	for elem := range other.elems {
		if s.Contains(elem) {
			result.Add(elem)
		}
	}

	return result
}

func (s *HashSet[E]) Difference(other *HashSet[E]) *HashSet[E] {
	result := New[E](s.Len())

	for elem := range s.elems {
		if !other.Contains(elem) {
			result.Add(elem)
		}
	}

	return result
}

func (s *HashSet[E]) IsSubSet(other *HashSet[E]) bool {
	for elem := range s.elems {
		if !other.Contains(elem) {
			return false
		}
	}

	return true
}

func (s *HashSet[E]) IsSuperSet(other *HashSet[E]) bool {
	for elem := range other.elems {
		if !s.Contains(elem) {
			return false
		}
	}

	return true
}

func (s *HashSet[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		for elem := range s.elems {
			if !yield(elem) {
				return
			}
		}
	}
}

func (s *HashSet[E]) Clear() {
	clear(s.elems)
}

func (s *HashSet[E]) Len() int {
	return len(s.elems)
}

func (s *HashSet[E]) String() string {
	return ""
}
