package main

import (
	"fmt"
)

type MinHeap struct {
	arr []int
}

func NewMaxHeap() *MinHeap {
	return &MinHeap{}
}

func getParentIndex(index int) int { return (index - 1) / 2 }
func getLeftIndex(index int) int   { return index*2 + 1 }
func getRightIndex(index int) int  { return index*2 + 2 }

func (h *MinHeap) push(value int) {
	h.arr = append(h.arr, value)
	h.heapifyUp()
}

func (h *MinHeap) pop() int {
	n := len(h.arr)
	if n == 0 {
		return -1
	}
	node := h.arr[0]
	lastIdx := n - 1
	h.arr[0] = h.arr[lastIdx]
	h.arr = h.arr[:lastIdx]

	h.heapifyDown()

	return node
}

func (h *MinHeap) heapifyDown() {
	index := 0
	n := len(h.arr)
	for {
		left := getLeftIndex(index)
		right := getRightIndex(index)

		smallestIdx := index
		if left < n && h.arr[left] < h.arr[smallestIdx] {
			smallestIdx = left
		}

		if right < n && h.arr[right] < h.arr[smallestIdx] {
			smallestIdx = right
		}

		if smallestIdx == index {
			break
		}

		h.arr[index], h.arr[smallestIdx] = h.arr[smallestIdx], h.arr[index]
		index = smallestIdx
	}

}

func (h *MinHeap) heapifyUp() {
	index := len(h.arr) - 1

	for index > 0 {
		parentIdx := getParentIndex(index)
		if h.arr[index] < h.arr[parentIdx] {
			h.arr[index], h.arr[parentIdx] = h.arr[parentIdx], h.arr[index]
			index = parentIdx
		} else {
			break
		}
	}
}

func (h *MinHeap) size() int {
	return len(h.arr)
}

func main() {
	mh := NewMaxHeap()
	mh.push(4)
	mh.push(2)
	mh.push(8)
	mh.push(5)
	mh.push(1)

	res := []int{}
	for mh.size() != 0 {
		res = append(res, mh.pop())
	}
	fmt.Println(res)

}
