package hashset

type HashSet[E comparable] map[E]struct{}

func New[E comparable](capacity int) HashSet[E] {
	return make(map[E]struct{}, capacity)
}

func (s HashSet[E]) Add(elems ...E) {
	for _, elem := range elems {
		s[elem] = struct{}{}
	}
}

func (s HashSet[E]) Remove(elems ...E) {
	for _, elem := range elems {
		delete(s, elem)
	}
}

func (s HashSet[E]) Contains(elem E) bool {
	_, ok := s[elem]
	return ok
}

func (s HashSet[E]) Union(other HashSet[E]) HashSet[E] {
	result := New[E](len(s) + len(other))

	for elem := range s {
		result.Add(elem)
	}

	for elem := range other {
		result.Add(elem)
	}

	return result
}

func (s HashSet[E]) Intersection(other HashSet[E]) HashSet[E] {
	result := New[E](len(s) + len(other))

	for elem := range other {
		if s.Contains(elem) {
			result.Add(elem)
		}
	}

	return result
}

func (s HashSet[E]) Clear() {
	clear(s)
}

func (s HashSet[E]) Len() int {
	return len(s)
}

func (s HashSet[E]) String() string {
	return ""
}
