package priority_queue

import (
	"github.com/aikin/go-algorithms/data-structures/heap"
)

type Element struct {
	Value    interface{}
	Priority int
}

func (x Element) Less(than heap.Item) bool {
	return x.Priority <= than.(Element).Priority
}

func NewElement(value interface{}, priority int) (i *Element) {
	return &Element{
		Value:    value,
		Priority: priority,
	}
}

type PriorityQueue struct {
	priorities heap.Heap
}

func NewMinHeapPriorityQueue() (pq *PriorityQueue) {
	return &PriorityQueue{priorities: *heap.NewMinHeap()}
}

func (pq *PriorityQueue) Add(el Element) {
	pq.priorities.Insert(heap.Item(el))
}

func (pq *PriorityQueue) Peek() (el Element) {
	return pq.priorities.Peek().(Element)
}

func (pq *PriorityQueue) Len() int {
	return pq.priorities.Len()
}
