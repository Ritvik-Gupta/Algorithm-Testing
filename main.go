package main

import (
	"container/heap"
	"fmt"

	valid_anagram "github.com/algorithm-testing/leetcode/easy/valid_anagram"
)

type PQElement struct {
	x, y int
}

func (elm *PQElement) String() string {
	return fmt.Sprintf("(%d, %d)", elm.x, elm.y)
}

type PriorityQueue []*PQElement

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].x == pq[j].x {
		return pq[i].y > pq[j].y
	}
	return pq[i].x > pq[j].x
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(element interface{}) {
	*pq = append(*pq, element.(*PQElement))
}

func (pq *PriorityQueue) Pop() interface{} {
	element := (*pq)[pq.Len()-1]
	*pq = (*pq)[0 : pq.Len()-1]
	return element
}

func main() {
	println(valid_anagram.IsAnagram("anagram", "nakgaram"))

	pq := &PriorityQueue{}
	heap.Init(pq)

	for i := 0; i < 10; i++ {
		var x, y int
		fmt.Print("Enter X and Y :\t")
		fmt.Scan(&x)
		fmt.Scan(&y)
		heap.Push(pq, &PQElement{x: x, y: y})
	}

	for i := 0; i < 10; i++ {
		fmt.Println(heap.Pop(pq))
	}
}
