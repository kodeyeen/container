package arraystack

type ArrayStack[T any] struct {
	elems []T
}

func New[T any](maxLen int) *ArrayStack[T] {
	return &ArrayStack[T]{
		elems: make([]T, 0, maxLen),
	}
}

func (s *ArrayStack[T]) Len() int {
	return len(s.elems)
}

func (s *ArrayStack[T]) Push(elems ...T) {
	s.elems = append(s.elems, elems...)
}

func (s *ArrayStack[T]) Pop() (T, bool) {
	if s.Len() == 0 {
		var t T
		return t, false
	}

	lastElem := s.elems[s.Len()-1]
	s.elems = s.elems[:s.Len()-1]
	return lastElem, true
}

func (s *ArrayStack[T]) Peek() (T, bool) {
	if len(s.elems) == 0 {
		var t T
		return t, false
	}

	return s.elems[len(s.elems)-1], true
}
