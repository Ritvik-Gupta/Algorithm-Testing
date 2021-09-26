package main

import (
	"github.com/algorithm-testing/leetcode/all_o_one_data_structure"
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
	register:= all_o_one_data_structure.Constructor()
	register.Inc("a")
	register.Inc("b")
	register.Inc("b")
	register.Inc("c")
	register.Inc("c")
	register.Inc("c")
	
	register.Dec("b")
	register.Dec("b")

	println("\t", register.GetMinKey())
	
	register.Dec("a")
	
	println("\t", register.GetMaxKey())
	println("\t", register.GetMinKey())
}

/*
["AllOne","inc a","inc b","inc b","inc c","inc c","inc c","dec b", "dec b","getMinKey","dec a","getMaxKey","getMinKey"]

[null,null,null,null,null,null,null,null,null,"b",null,"c","b"]

[null,null,null,null,null,null,null,null,null,"a",null,"c","c"]
*/
