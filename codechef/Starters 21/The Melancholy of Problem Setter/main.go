package main

import (
	"fmt"
)

const RED = 'R'
const GREEN = 'G'
const BLUE = 'B'

func main() {
	for testCase := readUint(); testCase > 0; testCase-- {
		if testRGBTreeConstruction() {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func testRGBTreeConstruction() bool {
	n := readUint()
	var nodeColors string
	fmt.Scan(&nodeColors)

	redNodesHaveBlueConnections := make(map[uint]bool)
	greenNodesHaveBlueConnections := make(map[uint]bool)

	for i := uint(1); i < n; i++ {
		u, v := readUint()-1, readUint()-1

		if nodeColors[v] == BLUE {
			u, v = v, u
		}

		if nodeColors[u] == BLUE {
			switch nodeColors[v] {
			case BLUE:
				return false
			case RED:
				if redNodesHaveBlueConnections[v] {
					return false
				}
				redNodesHaveBlueConnections[v] = true
			case GREEN:
				if greenNodesHaveBlueConnections[v] {
					return false
				}
				greenNodesHaveBlueConnections[v] = true
			}
		}
	}
	return true
}

func readUint() (store uint) {
	fmt.Scan(&store)
	return
}
