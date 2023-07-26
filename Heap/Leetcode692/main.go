package main

import (
	"container/heap"
	"fmt"
)

type Element struct {
	word      string
	frequency int
}

type PriorityQueue []Element

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].frequency < pq[j].frequency {
		return true
	}
	if pq[i].frequency == pq[j].frequency {
		if pq[i].word < pq[j].word {
			return true
		}
	}
	return false
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Element))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	ans := []string{}
	pq := new(PriorityQueue)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	for k, v := range m {
		pq.Push(Element{k, -v})
	}
	heap.Init(pq)
	for k > 0 && pq.Len() > 0 {
		val := heap.Pop(pq).(Element)
		ans = append(ans, val.word)
		k--
	}
	return ans
}

func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 3
	fmt.Println(topKFrequent(words, k))
}
