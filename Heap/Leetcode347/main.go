package main

import (
	"container/heap"
	"fmt"
)

type Element struct {
	x         int
	frequency int
}

type MinHeap []Element

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].frequency < h[j].frequency }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num] += 1
	}

	ans := []int{}
	h := new(MinHeap)

	for k, v := range m {
		h.Push(Element{k, -v})
	}
	heap.Init(h)
	for k > 0 && h.Len() > 0 {
		val := heap.Pop(h).(Element)
		ans = append(ans, val.x)
		k--
	}
	return ans
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
