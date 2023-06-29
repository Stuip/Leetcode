package main

import (
	"fmt"
)

/*
时间复杂度为O(n), 空间复杂度为O（n）
*/
func maxProfit(prices []int) int {
	n := len(prices)
	minPrices := make([]int, n)
	minPrices[0] = prices[0]
	profit := make([]int, n)
	for i := 1; i < n; i++ {
		if prices[i] < minPrices[i-1] {
			minPrices[i] = prices[i]
		} else {
			minPrices[i] = minPrices[i-1]
		}
		profit[i] = max(profit[i-1], prices[i]-minPrices[i])
	}
	return profit[n-1]
}

func maxProfit1(prices []int) int {
	n := len(prices)
	minPrice := prices[0]
	profit := make([]int, n)
	for i := 1; i < n; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		profit[i] = max(profit[i-1], prices[i]-minPrice)
	}
	return profit[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit1(prices))
}
