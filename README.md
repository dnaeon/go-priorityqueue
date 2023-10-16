# go-priorityqueue

[![Build Status](https://github.com/dnaeon/go-priorityqueue/actions/workflows/test.yaml/badge.svg)](https://github.com/dnaeon/go-priorityqueue/actions/workflows/test.yaml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/gopkg.in/dnaeon/go-priorityqueue.v1.svg)](https://pkg.go.dev/gopkg.in/dnaeon/go-priorityqueue.v1)
[![Go Report Card](https://goreportcard.com/badge/gopkg.in/dnaeon/go-priorityqueue.v1)](https://goreportcard.com/report/gopkg.in/dnaeon/go-priorityqueue.v1)
[![codecov](https://codecov.io/gh/dnaeon/go-priorityqueue/branch/v1/graph/badge.svg)](https://codecov.io/gh/dnaeon/go-priorityqueue)

A simple, generic implementation of [Priority
Queue](https://en.wikipedia.org/wiki/Priority_queue), based on
[container/heap](https://pkg.go.dev/container/heap).

## Installation

Executing the following command.

``` shell
go get -v gopkg.in/dnaeon/go-priorityqueue.v1
```

## Usage

``` go
package main

import (
	"fmt"

	pq "gopkg.in/dnaeon/go-priorityqueue.v1"
)

func main() {
	queue := pq.NewPriorityQueue[string, int64](pq.MinHeap)
	queue.Put("apple", 10)
	queue.Put("banana", 3)
	queue.Put("pear", 20)
	queue.Put("orange", 15)

	for !queue.IsEmpty() {
		item := queue.Get()
		fmt.Printf("%s: %d\n", item.Value, item.Priority)
	}
	// Output:
	// banana: 3
	// apple: 10
	// orange: 15
	// pear: 20
}
```

Make sure to check the included [test cases](./priority_queue_test.go) for
additional examples.

## Tests

Run the tests.

``` shell
make test
```

## License

`go-priorityqueue` is Open Source and licensed under the [BSD
License](http://opensource.org/licenses/BSD-2-Clause).
