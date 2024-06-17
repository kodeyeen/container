package arraystack

type ArrayStack[E any] struct {
	elems []E
}

func New[E any](maxLen int) *ArrayStack[E] {
	return &ArrayStack[E]{
		elems: make([]E, 0, maxLen),
	}
}

func (s *ArrayStack[E]) Len() int {
	return len(s.elems)
}

func (s *ArrayStack[E]) Push(elems ...E) {
	s.elems = append(s.elems, elems...)
}

func (s *ArrayStack[E]) Pop() (E, bool) {
	if s.Len() == 0 {
		var e E
		return e, false
	}

	lastElem := s.elems[s.Len()-1]
	s.elems = s.elems[:s.Len()-1]
	return lastElem, true
}

func (s *ArrayStack[E]) Peek() (E, bool) {
	if len(s.elems) == 0 {
		var e E
		return e, false
	}

	return s.elems[len(s.elems)-1], true
}
