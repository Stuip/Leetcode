package main

import "fmt"

func longestConsecutive(nums []int) int {
	// 建立一个存储所有数的哈希表，同时起到去重功能
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	ans := 0
	// 遍历去重后的所有数字
	for num := range set {
		cur := num
		// 只有当 num-1 不存在时，才开始向后遍历 num+1，num+2，num+3......
		if !set[cur-1] {
			for set[cur+1] {
				cur++
			}
		}
		// [num, cur] 之间是连续的，数字有 cur - num + 1 个
		ans = max(ans, cur-num+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestConsecutive1(nums []int) int {
	// key表示num，value表示num最远到达的连续右边界
	mapNums := make(map[int]int)
	// 初始化每个num的右边界为自己
	for _, num := range nums {
		mapNums[num] = num
	}

	ans := 0
	for _, num := range nums {
		if _, ok := mapNums[num-1]; !ok {
			right := mapNums[num]
			for _, ok := mapNums[right+1]; ok; _, ok = mapNums[right+1] {
				right = mapNums[right+1]
			}
			mapNums[num] = right
			ans = max(ans, right-num+1)
		}
	}
	return ans
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive1(nums))
}
