package main

import "fmt"

func platesBetweenCandles(s string, queries [][]int) []int {
	n, m := len(s), len(queries)
	ans, sums := make([]int, m), make([]int, n+1)
	list := []int{}
	for i := 0; i < n; i++ {
		if s[i] == '|' {
			list = append(list, i)
		}
		if s[i] == '*' {
			sums[i+1] = sums[i] + 1
		} else {
			sums[i+1] = sums[i]
		}
	}
	if len(list) == 0 {
		return ans
	}
	for i := 0; i < m; i++ {
		a, b := queries[i][0], queries[i][1]
		c, d := -1, -1
		l, r := 0, len(list)-1
		for l < r {
			mid := l + (r-l)/2
			if list[mid] >= a {
				r = mid
			} else {
				l = mid + 1
			}
		}
		if list[r] >= a {
			c = list[r]
		} else {
			continue
		}
		l, r = 0, len(list)-1
		for l < r {
			mid := l + r + 1>>1
			if list[mid] <= b {
				l = mid
			} else {
				r = mid - 1
			}
		}
		if list[r] <= b {
			d = list[r]
		} else {
			continue
		}
		if c <= d {
			ans[i] = sums[d+1] - sums[c]
		}
	}
	return ans
}

func main() {
	s := "**|**|***|"
	queries := [][]int{{2, 5}, {5, 9}}
	fmt.Println(platesBetweenCandles(s, queries))
}
