package main

import "fmt"

/**
令状态f(i, j)表示将下标位置i到j的所有的元素合并所获得的价值的最大值,
那么 f(i, j) = max(f(i, k) + f(k+1, j) + cost), cost是将两组元素起来的代价
*/

const INF = 2e9

func main() {
	var n int
	fmt.Scanln(&n) // 读取第一行的整数 n

	stones := make([]int, n) // 创建一个长度为 n 的整数切片

	for i := 0; i < n; i++ {
		fmt.Scan(&stones[i]) // 逐行读取整数并保存到切片中
	}
	fmt.Println(combineStones(stones))
}

func combineStones(stones []int) int {
	n := len(stones)
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] += preSum[i-1] + stones[i-1]
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for length := 2; length <= n; length++ { // 区间长度必须大于等于2
		for left := 1; left+length-1 <= n; left++ {
			right := left + length - 1
			dp[left][right] = INF
			for k := left; k < right; k++ {
				dp[left][right] = min(dp[left][right], dp[left][k]+dp[k+1][right]+preSum[right]-preSum[left-1])
			}
		}
	}
	return dp[1][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
