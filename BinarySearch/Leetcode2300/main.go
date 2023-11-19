package main

import (
	"fmt"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	m, n := len(spells), len(potions)
	sort.Ints(potions)
	ans := make([]int, m)
	for i := 0; i < m; i++ {
		left, right := 0, n-1
		for left < right {
			mid := (left + right) >> 1
			if int64(potions[mid]*spells[i]) >= success {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if int64(potions[right]*spells[i]) >= success {
			ans[i] = n - right
		}
	}
	return ans
}

func main() {
	spells := []int{5, 1, 3}
	potions := []int{1, 2, 3, 4, 5}
	var success int64 = 7
	fmt.Println(successfulPairs(spells, potions, success))
}
