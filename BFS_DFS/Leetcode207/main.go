package main

import "fmt"

// BFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	indegrees := make([]int, numCourses)
	adjacency := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		adjacency[i] = make([]int, 0)
	}
	for _, pre := range prerequisites {
		indegrees[pre[0]] += 1
		adjacency[pre[1]] = append(adjacency[pre[1]], pre[0])
	}
	queue := []int{}
	for idx, c := range indegrees {
		if c == 0 {
			queue = append(queue, idx)
		}
	}
	for len(queue) != 0 {
		pre := queue[0]
		queue = queue[1:]
		numCourses -= 1
		for _, c := range adjacency[pre] {
			indegrees[c] -= 1
			if indegrees[c] == 0 {
				queue = append(queue, c)
			}
		}
	}
	return numCourses == 0
}

// DFS
func canFinishDFS(numCourses int, prerequisites [][]int) bool {
	indegrees := make([]int, numCourses)
	adjacency := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		adjacency[i] = make([]int, 0)
	}
	for _, pre := range prerequisites {
		indegrees[pre[0]] += 1
		adjacency[pre[1]] = append(adjacency[pre[1]], pre[0])
	}
	queue := []int{}
	for idx, c := range indegrees {
		if c == 0 {
			queue = append(queue, idx)
		}
	}
	for len(queue) != 0 {
		pre := queue[0]
		queue = queue[1:]
		numCourses -= 1
		for _, c := range adjacency[pre] {
			indegrees[c] -= 1
			if indegrees[c] == 0 {
				queue = append(queue, c)
			}
		}
	}
	return numCourses == 0
}

func main() {
	numCourses := 2
	prerequisites := [][]int{
		{1, 0},
		{0, 1},
	}
	fmt.Println(canFinish(numCourses, prerequisites))
}
