package main

import (
	"fmt"
)

func main() {
	for testCase := readUint(); testCase > 0; testCase-- {
		readUint()
		var str string
		fmt.Scan(&str)
		strSize := len(str)

		numDiff := 0
		for i := (strSize - 1) / 2; i >= 0; i-- {
			if str[i] != str[strSize-1-i] {
				numDiff++
			}
		}

		fmt.Println((numDiff + 1) / 2)
	}
}

func readUint() (store uint) {
	fmt.Scan(&store)
	return
}
