package main

import "fmt"

const N = 3

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	ans := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	mat := [][]int{
		{1, 1, 1},
		{1, 0, 0},
		{0, 1, 0},
	}
	k := n - 2
	for k != 0 {
		if k&1 != 0 {
			ans = mul(ans, mat)
		}
		mat = mul(mat, mat)
		k >>= 1
	}
	return ans[0][0] + ans[0][1]
}

func mul(a, b [][]int) [][]int {
	res := make([][]int, N)
	for i := 0; i < N; i++ {
		res[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			res[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j] + a[i][2]*b[2][j]
		}
	}
	return res
}

func main() {
	n := 4
	fmt.Println(tribonacci(n))
}
