package main

import (
	"github.com/algorithm-testing/leetcode/valid_anagram"
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
	println(valid_anagram.IsAnagram("anagram", "nakgaram"))
}
