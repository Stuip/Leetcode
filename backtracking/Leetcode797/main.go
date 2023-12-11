package main

import "fmt"


func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	ans, path := [][]int{}, []int{0}
	var backtrack func(startIdx int)
	backtrack = func(startIdx int) {
		if startIdx == n-1 {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for _, v := range graph[startIdx] {
			path = append(path, v)
			backtrack(v)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return ans
}

func main() {
	graph := [][]int{
		{1,2},{3},{3},{},
	}
	fmt.Println(allPathsSourceTarget(graph))
}