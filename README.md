# container

This is an extended version of the Go's builtin container package.  
There are already similar ones, for example, [gods](https://github.com/emirpasic/gods), but I don't like it for being too Java-ish.

## Installation

```shell
go get github.com/kodeyeen/container
```

## Quickstart

```go
package main

import (
	"cmp"
	"fmt"

	"github.com/kodeyeen/container/priorityqueue"
)

func main() {
	pq := priorityqueue.New(cmp.Compare[int])

	pq.Init(44, 22, 11, 33)
	pq.Enqueue(-55)

	min, ok := pq.Peek()
	if ok {
		fmt.Printf("Minimum value is %d\n", min)
	}

	for pq.Len() > 0 {
		elem, _ := pq.Dequeue()

		fmt.Println(elem)
	}
}
```
