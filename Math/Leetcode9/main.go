package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reverse, num := 0, x
	for num != 0 {
		reverse = reverse*10 + num%10
		num /= 10
	}
	return reverse == x
}

func main() {
	fmt.Println(isPalindrome(math.MaxInt))
}
