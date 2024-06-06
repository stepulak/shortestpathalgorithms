package shortestpath

type CompareFunc[TValue any] func(a, b TValue) bool

type BinaryHeap[TValue any] struct {
	Tree       []TValue
	Comparator CompareFunc[TValue]
}

func NewBinaryHeap[TValue any](cmp CompareFunc[TValue]) *BinaryHeap[TValue] {
	return &BinaryHeap[TValue]{
		Tree:       make([]TValue, 0, 1),
		Comparator: cmp,
	}
}

func (h *BinaryHeap[TValue]) Size() int {
	return len(h.Tree)
}

func (h *BinaryHeap[TValue]) Contains(valueToFind TValue, cmp CompareFunc[TValue]) bool {
	for _, value := range h.Tree {
		if cmp(value, valueToFind) {
			return true
		}
	}
	return false
}

// Push value to the heap
func (h *BinaryHeap[TValue]) Push(value TValue) {
	h.Tree = append(h.Tree, value)
	h.bubbleUp(h.Size() - 1)
}

// Pop head (root) value
func (h *BinaryHeap[TValue]) Pop() (TValue, bool) {
	if h.Size() == 0 {
		var empty TValue
		return empty, false
	}
	head := h.Tree[0]
	h.Tree[0] = h.Tree[h.Size()-1]
	h.Tree = h.Tree[:h.Size()-1]
	h.bubbleDown(0)
	return head, true
}

func (h *BinaryHeap[TValue]) Head() (TValue, bool) {
	if h.Size() == 0 {
		var empty TValue
		return empty, false
	}
	return h.Tree[0], true
}

func (h *BinaryHeap[TValue]) bubbleDown(index int) {
	for index < h.Size() {
		parent := index
		left := getLeftChildIndex(index)
		right := getRightChildIndex(index)
		goLeft := left < h.Size() && h.Comparator(h.Tree[left], h.Tree[index])
		goRight := right < h.Size() && h.Comparator(h.Tree[right], h.Tree[index])

		if goLeft && goRight {
			if h.Comparator(h.Tree[right], h.Tree[left]) {
				index = right
			} else {
				index = left
			}
		} else if goLeft {
			index = left
		} else if goRight {
			index = right
		} else {
			break
		}
		h.swap(parent, index)
	}
}

func (h *BinaryHeap[TValue]) bubbleUp(index int) {
	parent := getParentIndex(index)
	for parent >= 0 && parent < index && h.Comparator(h.Tree[index], h.Tree[parent]) {
		h.swap(index, parent)
		index = parent
		parent = getParentIndex(parent)
	}
}

func (h *BinaryHeap[TValue]) swap(indexA, indexB int) {
	h.Tree[indexA], h.Tree[indexB] = h.Tree[indexB], h.Tree[indexA]
}

func getParentIndex(index int) int {
	return (index - 1) / 2
}

func getLeftChildIndex(index int) int {
	return index*2 + 1
}

func getRightChildIndex(index int) int {
	return index*2 + 2
}
