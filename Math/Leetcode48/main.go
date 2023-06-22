package main

import "fmt"

/**
先沿着对角线翻转，再沿着中线翻转
*/
func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i < j {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		left, right := 0, len(matrix[i])-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left += 1
			right -= 1
		}
	}

	fmt.Println(matrix)

}

func main() {
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}
