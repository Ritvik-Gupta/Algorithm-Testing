package main

import (
	"fmt"

	"github.com/algorithm-testing/leetcode/clone_graph"
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

type Node = clone_graph.Node

func printGraph(node *Node) {
	nodeQueue := make(chan *Node, 5)
	nodeQueue <- node
	visitedNodes := make(map[*Node]struct{})
	visitedNodes[node] = struct{}{}

iterationOverGraph:
	for {
		select {
		case node := <-nodeQueue:
			fmt.Printf("Node[%d] -> %p\n", node.Val, node)
			fmt.Printf("Neighbours :\n")

			for idx, neighbour := range node.Neighbors {
				fmt.Printf("\t%d. Node[%d] -> %p\n", idx+1, neighbour.Val, neighbour)
				if _, ok := visitedNodes[neighbour]; !ok {
					visitedNodes[neighbour] = struct{}{}
					nodeQueue <- neighbour
				}
			}
			fmt.Println()
		default:
			break iterationOverGraph
		}
	}
	close(nodeQueue)
}

func main() {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node3, node1}
	node3.Neighbors = []*Node{node4, node2}
	node4.Neighbors = []*Node{node1, node3}

	clonedNode := clone_graph.CloneGraph(node1)

	printGraph(node1)
	println("\n\n")
	printGraph(clonedNode)
}
