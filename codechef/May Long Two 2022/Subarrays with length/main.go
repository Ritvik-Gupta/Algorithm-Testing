package main

import (
	"fmt"
)

func main() {
	for testCase := readUint(); testCase > 0; testCase-- {
		arrSize := readUint()
		arr := make([]uint, 0, arrSize)
		for i := uint(0); i < arrSize; i++ {
			arr = append(arr, readUint())
		}
	}
}

func readUint() (store uint) {
	fmt.Scan(&store)
	return
}
