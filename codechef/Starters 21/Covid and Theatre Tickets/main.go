package main

import (
	"fmt"
	"math"
)

func main() {
	for testCase := readFloat(); testCase > 0; testCase-- {
		numRows, numSeats := readFloat(), readFloat()
		fmt.Println(math.Ceil(numRows/2.) * math.Ceil(numSeats/2.))
	}
}

func readFloat() (store float64) {
	fmt.Scan(&store)
	return
}
