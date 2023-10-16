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

package priorityqueue_test

import (
	"testing"

	pq "gopkg.in/dnaeon/go-priorityqueue.v1"
)

func TestMinPriorityQueue(t *testing.T) {
	queue := pq.NewPriorityQueue[string, int64](pq.MinHeap)
	queue.Put("apple", 10)
	queue.Put("banana", 3)
	queue.Put("pear", 20)
	queue.Put("orange", 15)

	want := []struct {
		value    string
		priority int64
	}{
		{"banana", 3},
		{"apple", 10},
		{"orange", 15},
		{"pear", 20},
	}

	i := 0
	for !queue.IsEmpty() {
		item := queue.Get()
		if item.Value != want[i].value || item.Priority != want[i].priority {
			t.Fatalf("want %q with priority %d, got %q with priority %d", want[i].value, want[i].priority, item.Value, item.Priority)
		}
		i += 1
	}
}

func TestMaxPriorityQueue(t *testing.T) {
	queue := pq.NewPriorityQueue[string, int64](pq.MaxHeap)
	queue.Put("apple", 10)
	queue.Put("banana", 3)
	queue.Put("pear", 20)
	queue.Put("orange", 15)

	want := []struct {
		value    string
		priority int64
	}{
		{"pear", 20},
		{"orange", 15},
		{"apple", 10},
		{"banana", 3},
	}

	i := 0
	for !queue.IsEmpty() {
		item := queue.Get()
		if item.Value != want[i].value || item.Priority != want[i].priority {
			t.Fatalf("want %q with priority %d, got %q with priority %d", want[i].value, want[i].priority, item.Value, item.Priority)
		}
		i += 1
	}
}
