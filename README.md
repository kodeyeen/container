# priorityqueue

A lightweight generic priority queue implementation in Go based on the example shown in [heap package documentation](https://pkg.go.dev/container/heap#example-package-PriorityQueue).
The goal is to provide a clean user friendly container and encapsulate the usage of the heap package.

## Installation

```shell
go get github.com/kodeyeen/priorityqueue
```

## Quickstart

```go
package main

import (
	"cmp"
	"fmt"

	"github.com/kodeyeen/priorityqueue"
)

func main() {
	pq := priorityqueue.New(cmp.Compare[int])

	pq.Init(44, 22, 11, 33)
	pq.Enqueue(-55)

	fmt.Printf("Minimum value is %d\n", pq.Peek())

	for pq.Len() > 0 {
		item := pq.Dequeue()

		fmt.Println(item)
	}
}
```
