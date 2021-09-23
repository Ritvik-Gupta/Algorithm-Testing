package main

import (
	"github.com/algorithm-testing/leetcode/time_needed_to_inform_all_employees"
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
	println(time_needed_to_inform_all_employees.NumOfMinutes(
		8,
		0,
		[]int{-1, 5, 0, 6, 7, 0, 0, 0},
		[]int{89, 0, 0, 0, 0, 523, 241, 519},
	))
}
