package container

type Container[T any] interface {
	Size() int
}
