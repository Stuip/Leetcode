package main

import "fmt"

var x, y, n int

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	dp := make([][]int, m*n)
	for i := 0; i < m*n; i++ {
		dp[i] = make([]int, maxMove)
	}
	return 0
}

func Index(x, y int) int {
	return x*n + y
}

func main() {
	m, n := 2, 2
	maxMove := 2
	startRow, startCol := 0, 0
	fmt.Println(findPaths(m, n, maxMove, startRow, startCol))
}
