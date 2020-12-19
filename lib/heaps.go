package lib

import "fmt"

type Heap interface {
	Print()
	Push(item *HeapItem)
	Pop() *HeapItem
	Len() int
	Empty() bool
	CheckHeapOrder() bool
	Peek() *HeapItem
	Remove(priority int)
}

var _ Heap = &MinHeap{}
var _ Heap = &MaxHeap{}

type HeapItem struct {
	Item     interface{}
	Priority int
}

// a minheap implementation
type MinHeap struct {
	heap []*HeapItem
	n    int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (h *MinHeap) Empty() bool {
	return h.Len() == 0
}

func (h *MinHeap) Len() int {
	return h.n
}

func (h *MinHeap) less(a, b int) bool {
	return h.heap[a].Priority < h.heap[b].Priority
}

func (h *MinHeap) swap(a, b int) {
	h.heap[a], h.heap[b] = h.heap[b], h.heap[a]
}

func (h *MinHeap) Push(s *HeapItem) {
	if h.n == 0 {
		h.heap = make([]*HeapItem, 2)
		h.heap[1] = s
		h.n++
		return
	}

	// this is where I fell into the trap of always appending to the heap slice, which may actually be much longer than
	// h.N  Don't forget to re-use positions in the slice.
	h.n++
	if h.n < len(h.heap) {
		h.heap[h.n] = s
	} else {
		h.heap = append(h.heap, s)
	}

	h.swim(h.n)
}

func (h *MinHeap) swim(n int) {
	for n > 1 {
		parent := n / 2
		// if n is less than the parent, swap with parent
		if h.less(n, parent) {
			h.swap(n, parent)
			n /= 2
		} else {
			break
		}
	}
}

func (h *MinHeap) sink(n int) {
	for {
		child1, child2 := 2*n, 2*n+1
		if child1 > h.n {
			return
		}

		if child2 > h.n {
			if h.less(child1, n) {
				h.swap(child1, n)
			}
			return
		}

		if h.less(child1, child2) {
			// promote child1
			h.swap(child1, n)
			n = child1
		} else {
			// promote child2
			h.swap(child2, n)
			n = child2
		}
	}
}

func (h *MinHeap) Pop() *HeapItem {
	if h.Len() == 0 {
		return nil
	}

	if h.Len() == 1 {
		h.n = 0
		nodeToReturn := h.heap[1]
		h.heap = []*HeapItem{}
		return nodeToReturn
	}

	nodeToReturn := h.heap[1]
	h.heap[1] = h.heap[h.n]
	h.n--
	h.sink(1)
	return nodeToReturn
}

func (h *MinHeap) Peek() *HeapItem {
	if h.Len() == 0 {
		return nil
	}

	return h.heap[1]
}

func (h *MinHeap) CheckHeapOrder() bool {
	for i := h.n; i > 1; i-- {
		if !h.less(i/2, i) {
			return false
		}
	}

	return true
}

func (h *MinHeap) Print() {
	fmt.Printf("MinHeap\n")
	for i := 1; i <= h.n; i++ {
		fmt.Printf("%d: %+v\n", i, h.heap[i])
	}
}

// Removes one entry with the given priority
func (h *MinHeap) Remove(priority int) {
	for i := 1; i <= h.n; i++ {
		if h.heap[i].Priority == priority {
			h.swap(i, h.n)
			h.heap[h.n] = nil
			h.n--
			h.sink(i)
			return
		}
	}
}

// a minheap implementation
type MaxHeap struct {
	heap []*HeapItem
	n    int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Empty() bool {
	return h.Len() == 0
}

func (h *MaxHeap) Len() int {
	return h.n
}

func (h *MaxHeap) less(a, b int) bool {
	return h.heap[a].Priority > h.heap[b].Priority
}

func (h *MaxHeap) swap(a, b int) {
	h.heap[a], h.heap[b] = h.heap[b], h.heap[a]
}

func (h *MaxHeap) Push(s *HeapItem) {
	if h.n == 0 {
		h.heap = make([]*HeapItem, 2)
		h.heap[1] = s
		h.n++
		return
	}

	h.n++
	if h.n < len(h.heap) {
		h.heap[h.n] = s
	} else {
		h.heap = append(h.heap, s)
	}
	h.swim(h.n)
}

func (h *MaxHeap) swim(n int) {
	for n > 1 {
		parent := n / 2
		// if n is greater than the parent, move it up
		if h.less(n, parent) {
			h.swap(n, parent)
			n /= 2
		} else {
			break
		}
	}
}

func (h *MaxHeap) sink(n int) {
	for {
		child1, child2 := 2*n, 2*n+1
		if child1 > h.n {
			return
		}

		if child2 > h.n {
			if h.less(child1, n) {
				h.swap(child1, n)
			}
			return
		}

		if h.less(child1, child2) {
			// promote child1
			h.swap(child1, n)
			n = child1
		} else {
			// promote child2
			h.swap(child2, n)
			n = child2
		}
	}
}

func (h *MaxHeap) Pop() *HeapItem {
	if h.Len() == 0 {
		return nil
	}

	if h.Len() == 1 {
		h.n = 0
		nodeToReturn := h.heap[1]
		h.heap = []*HeapItem{}
		return nodeToReturn
	}

	nodeToReturn := h.heap[1]
	h.heap[1] = h.heap[h.n]
	h.n--
	h.sink(1)
	return nodeToReturn
}

func (h *MaxHeap) Peek() *HeapItem {
	if h.Len() == 0 {
		return nil
	}

	return h.heap[1]
}

func (h *MaxHeap) CheckHeapOrder() bool {
	for i := h.n; i > 1; i-- {
		if !h.less(i/2, i) {
			return false
		}
	}

	return true
}

func (h *MaxHeap) Print() {
	fmt.Printf("MaxHeap\n")
	for i := 1; i <= h.n; i++ {
		fmt.Printf("%d: %+v\n", i, h.heap[i])
	}
}

// Removes one entry with the given priority
func (h *MaxHeap) Remove(priority int) {
	for i := 1; i <= h.n; i++ {
		if h.heap[i].Priority == priority {
			h.swap(i, h.n)
			h.heap[h.n] = nil
			h.n--
			h.sink(i)
			return
		}
	}
}

