package priority_queue

import (
	"errors"

	"github.com/aikin/go-algorithms/data-structures/heap"
	"github.com/aikin/go-algorithms/data-structures/queue"
)

type Element struct {
	Value    interface{}
	Priority int
}

func (x Element) Less(than heap.Item) bool {
	return x.Priority < than.(Element).Priority
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

func (pq *PriorityQueue) Len() int {
	return pq.priorities.Len()
}

func (pq *PriorityQueue) Add(el Element) {
	pq.priorities.Insert(heap.Item(el))
}

func (pq *PriorityQueue) Peek() (el Element) {
	return pq.priorities.Peek().(Element)
}

func (pq *PriorityQueue) Poll() (el Element) {
	return pq.priorities.Poll().(Element)
}

func (pq *PriorityQueue) ChangePriority(val interface{}, priority int) error {
	if pq.Len() == 0 {
		return errors.New("Priority queue is empty")
	}

	container := queue.NewQueue()

	popped := pq.Poll()
	for val != popped.Value {
		if pq.Len() == 0 {
			return errors.New("Element not found")
		}
		container.Enqueue(popped)
		popped = pq.Poll()
	}

	popped.Priority = priority
	pq.priorities.Insert(popped)

	for container.Len() > 0 {
		pq.priorities.Insert(container.Dequeue().(heap.Item))
	}
	return nil
}
