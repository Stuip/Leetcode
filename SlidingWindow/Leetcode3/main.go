package main

import "fmt"

func lengthOfLongestSubstring1(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	var freq [127]int
	ans, left, right := 0, 0, -1
	for left < n {
		if right+1 < n && freq[s[right+1]] == 0 {
			freq[s[right+1]] += 1
			right = right + 1
		} else {
			freq[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

/**哈系表+滑动窗口*/
func lengthOfLongestSubstring2(s string) int {
	left, right, ans := 0, 0, 0
	n := len(s)
	indexes := make(map[byte]int, n)
	for left < n {
		if idx, ok := indexes[s[left]]; ok && idx >= right {
			right = idx + 1
		}
		indexes[s[left]] = left
		left += 1
		ans = max(ans, left-right)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstring1("asdaafafsafafcdvsdfafasfdsadgsadgfwae"))
}
