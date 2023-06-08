package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	ans := []int{}
	m := make(map[int]int, len(nums1))
	for _, v := range nums1 {
		m[v]++
	}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			ans = append(ans, v)
			m[v]--
			if m[v] == 0 {
				delete(m, v)
			}
		}
	}
	return ans
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersect(nums1, nums2))
}
