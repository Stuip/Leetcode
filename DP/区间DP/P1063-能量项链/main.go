package main

import "fmt"

func main() {
	//var n int
	//fmt.Scan(&n)
	//energy := make([]int, n)
	//for i := 0; i < n; i++ {
	//	fmt.Scan(&energy[i])
	//}
	energy := []int{2, 3, 5, 10}
	fmt.Println(neckLace(energy))

}

func neckLace(energy []int) int {
	energy = append(energy, energy...)
	n := len(energy)
	sum := make([]int, n)
	sum[0] = energy[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + energy[i]
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for length := 2; length <= n; length++ {
		for left := 1; left+length-1 <= n; left++ {
			right := left + length - 1
			fmt.Println(left, right)
		}
	}
	return -1
}
