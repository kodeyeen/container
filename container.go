package container

import "fmt"

type Container[E any] interface {
	Clear()
	Len() int

	fmt.Stringer
}
