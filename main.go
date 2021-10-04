package main

import (
	"github.com/algorithm-testing/leetcode/design_circular_deque"
)

// import (
// 	"github.com/algorithm-testing/graph"
// )

// func main() {
// 	graph := graph.CreateGraph()

// 	graph.AddNode("A")
// 	graph.AddNode("B")
// 	graph.AddNode("C")

// 	graph.Connect("A", "B", 2)
// 	graph.Connect("B", "C", 1)
// 	graph.Connect("A", "C", 4)

// 	graph.Print()
// }

func main() {
	deque := design_circular_deque.Constructor(4)
	
	deque.Traverse()
	println(deque.InsertFront(9), "\n")
	
	deque.Traverse()
	println(deque.DeleteLast(), "\n")
	
	deque.Traverse()
	println(deque.GetRear(), "\n")
	
	deque.Traverse()
	println(deque.GetFront(), "\n")
	
	deque.Traverse()
	println(deque.GetFront(), "\n")
	
	deque.Traverse()
	println(deque.DeleteFront(), "\n")
	
	deque.Traverse()
	println(deque.InsertFront(6), "\n")
	
	deque.Traverse()
	println(deque.InsertLast(5), "\n")

	deque.Traverse()	
	println(deque.InsertFront(9), "\n")

	deque.Traverse()	
	println(deque.GetFront(), "\n")
	
	deque.Traverse()	
	println(deque.InsertFront(6), "\n")

	
	deque.Traverse()	
}

/*
["getRear","getFront","getFront","deleteFront","insertFront","insertLast","insertFront","getFront","insertFront"]
[[],[],[],[],[6],[5],[9],[],[6]]
*/
