package priorityqueue

// core is used to implement sort.Interface and heap.Interface.
// The purpose of this struct is to hide implementation from the user.
type core[T any] struct {
	items      []T
	comparator Comparator[T]
}

func (c *core[T]) Len() int {
	return len(c.items)
}

func (c *core[T]) Less(i, j int) bool {
	return c.comparator(c.items[i], c.items[j]) == -1
}

func (c *core[T]) Swap(i, j int) {
	c.items[i], c.items[j] = c.items[j], c.items[i]
}

func (c *core[T]) Push(x any) {
	c.items = append(c.items, x.(T))
}

func (c *core[T]) Pop() any {
	n := len(c.items)
	item := c.items[n-1]
	c.items = c.items[0 : n-1]
	return item
}
