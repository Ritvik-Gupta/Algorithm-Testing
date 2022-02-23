package main

import (
	"fmt"
)

type empty = struct{}

type Pair struct {
	a, b uint
}

func main() {
	n, m := readUint(), readUint()
	containedInRow := make(map[uint]map[uint]empty)
	deniedPairs := make(map[Pair]empty)
	w := make([]uint, 0, n)

	for i := uint(0); i < n; i++ {
		for j := uint(0); j < m; j++ {
			val := readUint()
			if _, ok := containedInRow[val]; !ok {
				containedInRow[val] = make(map[uint]empty)
			}
			for k := range containedInRow[val] {
				deniedPairs[Pair{k, i}] = empty{}
			}
			containedInRow[val][i] = empty{}
		}
		w = append(w, readUint())
	}

	minVal := ^uint(0)
	for i := uint(0); i < n; i++ {
		for j := i + 1; j < n; j++ {
			if _, ok := deniedPairs[Pair{i, j}]; !ok {
				if val := w[i] + w[j]; val < minVal {
					minVal = val
				}
			}
		}
	}
	if minVal == ^uint(0) {
		fmt.Println(-1)
	} else {
		fmt.Println(minVal)
	}
}

func readUint() (store uint) {
	fmt.Scan(&store)
	return
}
