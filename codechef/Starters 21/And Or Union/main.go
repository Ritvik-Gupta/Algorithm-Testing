package main

import (
	"fmt"
)

func main() {
	for testCase := readUint(); testCase > 0; testCase-- {
		n := readUint()
		arr := make([]uint, 0, n)
		for idx := n; idx > 0; idx-- {
			arr = append(arr, readUint())
		}

		cumulativeReversedOr, result := uint(0), uint(0)
		for i := n; i > 0; i-- {
			result |= arr[i-1] & cumulativeReversedOr
			cumulativeReversedOr |= arr[i-1]
		}

		fmt.Println(result)
	}
}

func readUint() (store uint) {
	fmt.Scan(&store)
	return
}
