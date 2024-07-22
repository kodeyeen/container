package heap

import (
	"github.com/kodeyeen/container"
)

type Heap[E any] struct {
	elems      []E
	comparator container.Comparator[E]
}

func New[E any](comparator container.Comparator[E]) *Heap[E] {
	return &Heap[E]{
		elems:      make([]E, 0),
		comparator: comparator,
	}
}

func (h *Heap[E]) Init(elems ...E) {
	n := len(h.elems)
	h.elems = make([]E, 0, n)
	h.elems = append(h.elems, elems...)

	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}

func (h *Heap[E]) Push(elems ...E) {
	for _, elem := range elems {
		h.elems = append(h.elems, elem)
		h.up(len(h.elems) - 1)
	}
}

func (h *Heap[E]) Pop() (E, bool) {
	n := len(h.elems)

	if n == 0 {
		var e E
		return e, false
	}

	h.swap(0, n-1)

	h.down(0, n-1)

	elem := h.elems[n-1]
	h.elems = h.elems[0 : n-1]
	return elem, true
}

func (h *Heap[E]) Peek() (E, bool) {
	if len(h.elems) == 0 {
		var e E
		return e, false
	}

	return h.elems[0], true
}

func (h *Heap[E]) less(i, j int) bool {
	return h.comparator(h.elems[i], h.elems[j]) == -1
}

func (h *Heap[E]) swap(i, j int) {
	h.elems[i], h.elems[j] = h.elems[j], h.elems[i]
}

func (h *Heap[E]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.less(j, i) {
			break
		}
		h.swap(i, j)
		j = i
	}
}

func (h *Heap[E]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.less(j, i) {
			break
		}
		h.swap(i, j)
		i = j
	}
	return i > i0
}

func (h *Heap[E]) Clear() {
	h.elems = make([]E, 0)
}

func (h *Heap[E]) Len() int {
	return len(h.elems)
}

func (h *Heap[E]) String() string {
	return ""
}
