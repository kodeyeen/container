package stack

type Stack[E any] struct {
	elems []E
}

func New[E any](capacity int) *Stack[E] {
	return &Stack[E]{
		elems: make([]E, 0, capacity),
	}
}

func (s *Stack[E]) Push(elems ...E) {
	s.elems = append(s.elems, elems...)
}

func (s *Stack[E]) Pop() (E, bool) {
	if len(s.elems) == 0 {
		var e E
		return e, false
	}

	lastElem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return lastElem, true
}

func (s *Stack[E]) Peek() (E, bool) {
	if len(s.elems) == 0 {
		var e E
		return e, false
	}

	return s.elems[len(s.elems)-1], true
}

func (s *Stack[E]) Clear() {
	s.elems = make([]E, 0)
}

func (s *Stack[E]) Len() int {
	return len(s.elems)
}

func (s *Stack[E]) String() string {
	return ""
}
