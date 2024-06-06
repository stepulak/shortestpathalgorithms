package shortestpath

import (
	"math/rand"
	"testing"
)

func TestBinaryHeap(t *testing.T) {
	bh := NewBinaryHeap(func(a, b int) bool {
		return a < b
	})

	// Binary heap construction
	if bh == nil || bh.Tree == nil || bh.Size() != 0 {
		t.Error("TestBinaryHeap: construction error")
	}

	// Check methods on empty heap
	if _, ok := bh.Head(); ok {
		t.Error("TestBinaryHeap: head should be nil")
	}
	if _, ok := bh.Pop(); ok {
		t.Error("TestBinaryHeap: pop unexpectedly succeeded")
	}

	// Push and check bubble up algorithm
	bh.Push(1)
	if val, ok := bh.Head(); val != 1 || !ok || bh.Size() != 1 {
		t.Error("TestBinaryHeap: head should be 1, size should be 1")
	}
	bh.Push(-1)
	if val, ok := bh.Head(); val != -1 || !ok || bh.Size() != 2 {
		t.Error("TestBinaryHeap: head should be -1, size should be 2")
	}
	bh.Push(-2)
	if val, ok := bh.Head(); val != -2 || !ok || bh.Size() != 3 {
		t.Error("TestBinaryHeap: head should be -2, size should be 3")
	}
	bh.Push(2)
	if val, ok := bh.Head(); val != -2 || !ok || bh.Size() != 4 {
		t.Error("TestBinaryHeap: head still should be -2, size should be 4")
	}

	findCmp := func(a, b int) bool { return a == b }
	if !bh.Contains(1, findCmp) || !bh.Contains(-1, findCmp) || !bh.Contains(-2, findCmp) || !bh.Contains(2, findCmp) {
		t.Error("TestBinaryHeap: unable to find elements present in binary heap")
	}

	// Pop and check bubble down algorithm
	if val, ok := bh.Pop(); val != -2 || !ok || bh.Size() != 3 {
		t.Error("TestBinaryHeap: pop should be -1, size should be 3")
	}
	if val, ok := bh.Pop(); val != -1 || !ok || bh.Size() != 2 {
		t.Error("TestBinaryHeap: pop should be -1, size should be 2")
	}
	if val, ok := bh.Head(); val != 1 || !ok {
		t.Error("TestBinaryHeap: head should be 1")
	}
	if val, ok := bh.Pop(); val != 1 || !ok || bh.Size() != 1 {
		t.Error("TestBinaryHeap: pop should be -1, size should be 1")
	}
	if val, ok := bh.Pop(); val != 2 || !ok || bh.Size() != 0 {
		t.Error("TestBinaryHeap: pop should be 2, size should be 0")
	}
	if _, ok := bh.Pop(); ok || bh.Size() != 0 {
		t.Error("TestBinaryHeap: pop should be empty, size should be 0")
	}

	if len(bh.Tree) != 0 {
		t.Error("TestBinaryHeap: length of heap's tree should be 0")
	}
}

func TestBinaryHeap_complex(t *testing.T) {
	type Value struct {
		value int
	}

	bh := NewBinaryHeap[Value](func(a, b Value) bool {
		return a.value <= b.value
	})

	// Push pseudorandom valuess
	for i := 0; i < 1000; i++ {
		bh.Push(Value{rand.Int()})
	}

	head, ok := bh.Head()
	if !ok {
		t.Error("TestBinaryHeap_complex: head is empty")
	}

	// Pop them all
	// They need to be sorted
	for bh.Size() > 0 {
		if val, ok := bh.Pop(); !ok || !bh.Comparator(head, val) {
			t.Error("TestBinaryHeap_complex: pop error")
		}
	}

	if len(bh.Tree) != 0 {
		t.Error("TestBinaryHeap_complex: length of heap's tree should be 0")
	}
}
