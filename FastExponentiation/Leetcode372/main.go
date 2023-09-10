package main

import "fmt"

const MOD = 1337

func superPow(a int, b []int) int {
	return dfs(a, b, len(b)-1)
}

func dfs(a int, b []int, u int) int {
	if u == -1 {
		return -1
	}
	return qpow(dfs(a, b, u-1), 10) * qpow(a, b[u]) % MOD
}

func qpow(a, b int) int {
	ans := 1
	a %= MOD
	for b != 0 {
		if (b & 1) != 0 {
			ans = ans * a % MOD
		}
		a = a * a % MOD
		b >>= 1
	}
	return ans
}

func main() {
	a := 2
	b := []int{1, 0}
	fmt.Println(superPow(a, b))
}
