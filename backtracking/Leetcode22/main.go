package main

import "fmt"

func generateParenthesis(n int) []string {
	ans := []string{}
	var backtrack func(left, right int, str string)
	backtrack = func(left, right int, str string) {
		if right > left {
			return
		}
		if left == n && right == n {
			ans = append(ans, str)
			return
		}
		if left < n {
			backtrack(left+1, right, str+"(")
		}
		if right < left {
			backtrack(left, right+1, str+")")
		}
	}
	backtrack(0, 0, "")
	return ans
}

func main() {
	fmt.Println(generateParenthesis(3))
}
