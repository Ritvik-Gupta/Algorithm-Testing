package rcb_and_playoffs

import "fmt"

const MAX_POINTS_PER_MATCH uint = 2

func main() {
	var numTestCases uint
	fmt.Scanln(&numTestCases)

	for ; numTestCases > 0; numTestCases-- {
		var currentPoints, requiredPoints, matchesLeft uint
		fmt.Scan(&currentPoints)
		fmt.Scan(&requiredPoints)
		fmt.Scan(&matchesLeft)

		if requiredPoints-currentPoints <= matchesLeft*MAX_POINTS_PER_MATCH {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
