package main

import (
	"fmt"
)

func main() {
	for testCase := readUint(); testCase > 0; testCase-- {
		_, startX, startY := readUint(), readUint(), readUint()
		fmt.Println((startX + startY) & 1)
	}
}

func readUint() (store uint) {
	fmt.Scan(&store)
	return
}
