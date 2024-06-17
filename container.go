package container

type Container[T any] interface {
	Len() int
}
