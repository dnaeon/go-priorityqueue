// Copyright (c) 2023 Marin Atanasov Nikolov <dnaeon@gmail.com>
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
//
//   1. Redistributions of source code must retain the above copyright
//      notice, this list of conditions and the following disclaimer.
//   2. Redistributions in binary form must reproduce the above copyright
//      notice, this list of conditions and the following disclaimer in the
//      documentation and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package priorityqueue

import (
	"container/heap"
)

// HeapKind specifies the heap kind - min or max
type HeapKind int

const (
	// A heap which yields max-value items
	MinHeap HeapKind = iota

	// A heap which yields min-value items
	MaxHeap
)

// Item represents an item from the priority queue.
type Item[T any, V int64 | float64] struct {
	// The value associated with the item
	Value T

	// The priority of the item
	Priority V

	// The index is needed by update and is maintained by the
	// heap.Interface methods.
	index int
}

// PriorityQueue is a priority queue implementation based
// container/heap
type PriorityQueue[T any, V int64 | float64] struct {
	items []*Item[T, V]
	kind  HeapKind
}

// NewPriorityQueue creates a new priority queue, containing items of
// type T with priority V.
func NewPriorityQueue[T any, V int64 | float64](kind HeapKind) *PriorityQueue[T, V] {
	pq := &PriorityQueue[T, V]{
		items: make([]*Item[T, V], 0),
		kind:  kind,
	}
	heap.Init(pq)

	return pq
}

// Len implements sort.Interface
func (pq PriorityQueue[T, V]) Len() int {
	return len(pq.items)
}

// Less implements sort.Interface
func (pq PriorityQueue[T, V]) Less(i, j int) bool {
	if pq.kind == MinHeap {
		return pq.items[i].Priority < pq.items[j].Priority
	}

	return pq.items[i].Priority > pq.items[j].Priority
}

// Swap implements sort.Interface
func (pq PriorityQueue[T, V]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].index = i
	pq.items[j].index = j
}

// Push implements heap.Interface
func (pq *PriorityQueue[T, V]) Push(x any) {
	n := len(pq.items)
	item := x.(*Item[T, V])
	item.index = n
	pq.items = append(pq.items, item)
}

// Pop implements heap.Interface
func (pq *PriorityQueue[T, V]) Pop() any {
	old := pq.items
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	pq.items = old[0 : n-1]
	return item
}

// Put adds a value with the given priority to the priority queue
func (pq *PriorityQueue[T, V]) Put(value T, priority V) {
	n := len(pq.items)
	item := &Item[T, V]{
		Value:    value,
		Priority: priority,
		index:    n,
	}
	pq.items = append(pq.items, item)
	heap.Fix(pq, n)
}

// Get returns the next item from the priority queue
func (pq *PriorityQueue[T, V]) Get() *Item[T, V] {
	item := heap.Pop(pq)
	return item.(*Item[T, V])
}

// IsEmpty returns a boolean indicating whether the priority queue is
// empty or not
func (pq *PriorityQueue[T, V]) IsEmpty() bool {
	return pq.Len() == 0
}
