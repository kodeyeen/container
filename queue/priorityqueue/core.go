package priorityqueue

// core is used to implement sort.Interface and heap.Interface.
// The purpose of this struct is to hide implementation from the user.
type core[T any] struct {
	elems      []T
	comparator Comparator[T]
}

func (c *core[T]) Len() int {
	return len(c.elems)
}

func (c *core[T]) Less(i, j int) bool {
	return c.comparator(c.elems[i], c.elems[j]) == -1
}

func (c *core[T]) Swap(i, j int) {
	c.elems[i], c.elems[j] = c.elems[j], c.elems[i]
}

func (c *core[T]) Push(x any) {
	c.elems = append(c.elems, x.(T))
}

func (c *core[T]) Pop() any {
	n := len(c.elems)
	elem := c.elems[n-1]
	c.elems = c.elems[0 : n-1]
	return elem
}
