package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ConstructLinkedList(nums []int) *ListNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	dummy := &ListNode{-1, nil}
	cur := dummy
	for i := 0; i < n; i++ {
		cur.Next = &ListNode{nums[i], nil}
		cur = cur.Next
	}
	return dummy.Next
}

func PrintLinkedList(l *ListNode) []int {
	res := []int{}
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}


 type MinHeap struct {
    arr []*ListNode
}

func getParentIndex(index int) int { return (index - 1) / 2 }

func getLeftIndex(index int) int { return index * 2 + 1 }
func getRightIndex(index int) int { return index * 2 + 2 }

func NewHeap() *MinHeap {
    return &MinHeap{}
}

func (h *MinHeap) push(value *ListNode) {
    h.arr = append(h.arr, value)
    h.heapifyUp()
} 

func (h *MinHeap) pop() *ListNode {
    n := h.size()
    if n == 0 {
        return nil
    }
    node := h.arr[0]
    lastIdx := n - 1
    h.arr[0] = h.arr[lastIdx]
    h.arr = h.arr[:lastIdx]
    h.heapifyDown()
    return node
}

func (h *MinHeap) heapifyUp() {
    index := h.size() - 1
    for index > 0 {
        parentIdx := getParentIndex(index)
        if h.arr[index].Val < h.arr[parentIdx].Val {
            h.arr[parentIdx], h.arr[index] = h.arr[index], h.arr[parentIdx]
            index = parentIdx
        } else {
            break
        }
    }
}

func (h *MinHeap) heapifyDown() {
    index := 0
    n := h.size()
    for {
        left := getLeftIndex(index)
        right := getRightIndex(index)

        smallIndex := index
        if left < n && h.arr[left].Val < h.arr[smallIndex].Val {
            smallIndex = left
        }

        if right < n && h.arr[right].Val < h.arr[smallIndex].Val {
            smallIndex = right
        }
        
        if smallIndex == index {
            break
        }

        h.arr[smallIndex], h.arr[index] = h.arr[index], h.arr[smallIndex]

        index = smallIndex
    }
}


func (h *MinHeap) size() int {
    return len(h.arr)
}


func mergeKLists(lists []*ListNode) *ListNode {
    mh := NewHeap()
    for _, list := range lists {
        cur := list
        for cur != nil {
			next := cur.Next
			cur.Next = nil
            mh.push(cur)
            cur = next
        }
    }

    dummy := &ListNode{-1, nil}
    cur := dummy
    for mh.size() != 0 {
        cur.Next = mh.pop()
        cur = cur.Next
    }
    return dummy.Next
}

func main() {
	nums := [][]int{
		{-3,-3,-2},
		{-3,-3,-2,-2},
	}

	lists := []*ListNode{}
	for i:=0;i<len(nums);i++{
		lists = append(lists, ConstructLinkedList(nums[i]))
	}
	fmt.Println(PrintLinkedList(mergeKLists(lists)))
}