package main

import "fmt"

const mod = 1000000007

var cache [][]int

func countRoutes(locations []int, start int, finish int, fuel int) int {
	// cache[i][fuel] 表示从位置i出发，当前剩余的油量为fuel的前提下，到达目标位置的「路径数量」。
	n := len(locations)
	cache = make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, fuel+1)
		for j := 0; j <= fuel; j++ {
			cache[i][j] = -1
		}
	}
	return dfs(locations, start, finish, fuel)
}

func dfs(ls []int, u, end, fuel int) int {
	if cache[u][fuel] != -1 {
		return cache[u][fuel]
	}

	n := len(ls)

	// 第一种情况：油量为0且不在目标位置
	if fuel == 0 && u != end {
		cache[u][fuel] = 0
		return 0
	}

	// 第二种情况是油量不为0,但是无法到达任何位置
	hasNext := false
	for i := 0; i < n; i++ {
		if i != u {
			need := abs(ls[i] - ls[u])
			if fuel >= need {
				hasNext = true
				break
			}
		}
	}
	if fuel != 0 && !hasNext {
		cache[u][fuel] = boolToInt(u == end)
	}

	sum := boolToInt(u == end)
	for i := 0; i < n; i++ {
		if i != u {
			need := abs(ls[i] - ls[u])
			if fuel >= need {
				sum += dfs(ls, i, end, fuel-need)
				sum %= mod
			}
		}
	}
	cache[u][fuel] = sum
	return sum
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func countRoutes_DP(locations []int, start int, finish int, fuel int) int {
	n := len(locations)
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, fuel+1)
	}
	for i := 0; i <= fuel; i++ {
		memo[finish][i] = 1
	}

	for cur := 0; cur <= fuel; cur++ {
		for i := 0; i < n; i++ {
			for k := 0; k < n; k++ {
				if i != k {
					need := abs(locations[i] - locations[k])
					if cur >= need {
						memo[i][cur] += memo[k][cur-need]
						memo[i][cur] %= mod
					}
				}
			}
		}
	}
	return memo[start][fuel]
}

func main() {
	locations := []int{2, 3, 6, 8, 4}
	start := 1
	finish := 3
	fuel := 5
	fmt.Println(countRoutes_DP(locations, start, finish, fuel))
}
