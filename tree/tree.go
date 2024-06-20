package tree

import "github.com/kodeyeen/container"

type Tree[E comparable] interface {
	container.Container[E]
}
