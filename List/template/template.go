package template

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
