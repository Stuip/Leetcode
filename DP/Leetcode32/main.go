package main

import "fmt"

// 使用动态规划
func longestValidParentheses_DP(s string) int {
	maxAns := 0
	n := len(s)
	dp := make([]int, n) // 记录以i结尾的字符的最长有效括号的长度
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					// 第2个索引之后
					dp[i] = dp[i-2] + 2
				} else {
					// 即"()"
					dp[i] = 2
				}
				//     i-dp[i-1]>0
				// '....(......)'
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		maxAns = max(maxAns, dp[i])
	}
	return maxAns
}


func longestValidParentheses_Stack(s string) int {
	maxAns := 0
	n, stack := len(s), []int{}
	stack = append(stack, -1)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}
	return maxAns
}


func longestValidParentheses_TwoPointer(s string) int {
	left, right, maxLength := 0, 0, 0
	n := len(s)
	for i:=0;i<n;i++{
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2 * right)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i:=n-1;i>=0;i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2 * left)
		} else if right > left {
			left, right = 0, 0
		}
	}
	return maxLength
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "()(())"
	fmt.Println(longestValidParentheses_DP(s))
	fmt.Println(longestValidParentheses_Stack(s))
	fmt.Println(longestValidParentheses_TwoPointer(s))
}
