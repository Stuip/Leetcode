package main

import (
	"fmt"
	"sort"
)

// O(n^2)  超时
func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return n
	}
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0]
	})
	ans := 1
	dp := make([]int, n) // 考虑前i个信封,并以i结尾的最大值
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if check(envelopes, j, i) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(dp[i], ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func check(envelopes [][]int, i, j int) bool {
	return envelopes[i][0] < envelopes[j][0] && envelopes[i][1] < envelopes[j][1]
}

/**
这里为什么对信封排序是先按宽度递增排序,然后再按照长度递减排序呢?
	我可能思考,如果我们存在这样一个信封组[[1,1],[1,2], [1,3], [1,4]]
	如果长度也是递增排序的话,那么可以发现能够将四封信封依次装入,则是一个错误思维.
	对于这种情况最多能装入1个信封,所以长度需要按递减顺序排序
*/
func maxEnvelopes_opt(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] > b[1])
	})
	f := []int{}
	for _, e := range envelopes {
		h := e[1]
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	return len(f)
}

func main() {
	envelopes := [][]int{
		{5, 4}, {6, 4}, {6, 7}, {2, 3},
	}
	fmt.Println(maxEnvelopes_opt(envelopes))
}
