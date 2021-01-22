package heap

type Item interface {
	Less(than Item) bool
}

type Int int

func (x Int) Less(than Item) bool {
	return x < than.(Int)
}

type Heap struct {
	data []Item
	min  bool
}

func New() *Heap {
	return &Heap{
		data: make([]Item, 0),
	}
}

func NewMaxHeap() *Heap {
	return &Heap{
		data: make([]Item, 0),
		min:  false,
	}
}

func NewMinHeap() *Heap {
	return &Heap{
		data: make([]Item, 0),
		min:  true,
	}
}

func (h *Heap) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *Heap) Len() int {
	return len(h.data)
}

func (h *Heap) Get(n int) Item {
	return h.data[n]
}

func (h *Heap) Insert(n Item) {
	h.data = append(h.data, n)
	h.siftUp()
	return
}

func (h *Heap) Poll() (el Item) {
	if h.Len() == 0 {
		return
	}

	el = h.data[0]
	last := h.data[h.Len()-1]
	if (h.Len()) == 1 {
		h.data = nil
		return
	}

	h.data = append([]Item{last}, h.data[1:h.Len()-1]...)
	h.siftDown()

	return
}

func (h *Heap) Peek() (el Item) {
	if h.Len() == 0 {
		return
	}
	return h.data[0]
}

func (h *Heap) siftUp() {
	for i, parent := h.Len()-1, h.Len()-1; i > 0; i = parent {
		parent = i >> 1
		if h.Less(h.Get(i), h.Get(parent)) {
			h.data[parent], h.data[i] = h.data[i], h.data[parent]
		} else {
			break
		}
	}
}

func (h *Heap) siftDown() {
	for i, child := 0, 1; i < h.Len() && i<<1+1 < h.Len(); i = child {
		child = i<<1 + 1

		if child+1 <= h.Len()-1 && h.Less(h.Get(child+1), h.Get(child)) {
			child++
		}

		if h.Less(h.Get(i), h.Get(child)) {
			break
		}
		h.data[i], h.data[child] = h.data[child], h.data[i]
	}
}

func (h *Heap) Less(parent, child Item) bool {
	if h.min {
		return parent.Less(child)
	}
	return child.Less(parent)
}
