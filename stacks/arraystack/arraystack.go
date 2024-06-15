package arraystack

type ArrayStack[T any] struct {
	items []T
}

func New[T any](maxLen int) *ArrayStack[T] {
	return &ArrayStack[T]{
		items: make([]T, 0, maxLen),
	}
}

func (s *ArrayStack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var t T
		return t, false
	}

	return s.items[len(s.items)-1], true
}

func (s *ArrayStack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *ArrayStack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var t T
		return t, false
	}

	lastItem := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return lastItem, true
}

func (s *ArrayStack[T]) Len() int {
	return len(s.items)
}
