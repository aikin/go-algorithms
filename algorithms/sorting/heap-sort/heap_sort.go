package heap_sort

import "github.com/aikin/go-algorithms/data-structures/heap"

type Int int

func (x Int) Less(than heap.Item) bool {
	return x < than.(Int)
}

func Sort(nums []int) {
	h := heap.NewMinHeap()
	for i := 0; i < len(nums); i++ {
		h.Insert(Int(nums[i]))
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = int(h.Poll().(Int))
	}
}
