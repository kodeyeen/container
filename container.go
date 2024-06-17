package container

import "fmt"

type Container[E comparable] interface {
	Clear()
	Len() int

	fmt.Stringer
}
