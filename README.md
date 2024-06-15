# container

A set of containers that I had to implement during solving Leetcode problems.
Also it serves as my personal data structures cheatsheet.

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

	"github.com/kodeyeen/container/queue/priorityqueue"
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
		item, _ := pq.Dequeue()

		fmt.Println(item)
	}
}
```
