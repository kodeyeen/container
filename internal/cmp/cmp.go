package cmp

type Comparator[T any] func(x, y T) int
