package main

import "fmt"

func nonHatedPlayerGroups(numPlayers int, a []int, b []int) int {
	closestHatedPlayers := make([]int, 0, numPlayers)
	for i := 0; i < numPlayers; i++ {
		closestHatedPlayers = append(closestHatedPlayers, numPlayers)
	}

	for i := 0; i < len(a); i++ {
		x, y := a[i], b[i]
		if x > y {
			x, y = y, x
		}

		if closestHatedPlayers[x-1] > y-1 {
			closestHatedPlayers[x-1] = y - 1
		}
	}

	totalGroups := 0
	minHatedPlayer := numPlayers

	for i := numPlayers - 1; i >= 0; i-- {
		if minHatedPlayer > closestHatedPlayers[i] {
			minHatedPlayer = closestHatedPlayers[i]
		}
		totalGroups += minHatedPlayer - i
	}

	return totalGroups
}

func main() {
	numPlayers := readInt()
	m := readInt()
	a, b := make([]int, 0, m), make([]int, 0, m)

	for i := 0; i < m; i++ {
		a = append(a, readInt())
		b = append(b, readInt())
	}

	fmt.Printf("%d", nonHatedPlayerGroups(numPlayers, a, b))
}

func readInt() (store int) {
	fmt.Scan(&store)
	return
}
