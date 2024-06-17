package container

type Container[E any] interface {
	Len() int
}
