package arraystack

type ArrayStack[E any] struct {
	elems []E
}

func New[E any](capacity int) *ArrayStack[E] {
	return &ArrayStack[E]{
		elems: make([]E, 0, capacity),
	}
}

func (s *ArrayStack[E]) Push(elems ...E) {
	s.elems = append(s.elems, elems...)
}

func (s *ArrayStack[E]) Pop() (E, bool) {
	if len(s.elems) == 0 {
		var e E
		return e, false
	}

	lastElem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return lastElem, true
}

func (s *ArrayStack[E]) Peek() (E, bool) {
	if len(s.elems) == 0 {
		var e E
		return e, false
	}

	return s.elems[len(s.elems)-1], true
}

func (s *ArrayStack[E]) Clear() {
	s.elems = make([]E, 0)
}

func (s *ArrayStack[E]) Len() int {
	return len(s.elems)
}

func (s *ArrayStack[E]) String() string {
	return ""
}
