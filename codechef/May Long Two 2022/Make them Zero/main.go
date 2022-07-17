package main

import (
	"fmt"
)

func main() {
	for testCase := readUint(); testCase > 0; testCase-- {
		arrSize := readUint()
		arrBitwiseOr := uint(0)
		for i := uint(0); i < arrSize; i++ {
			arrBitwiseOr |= readUint()
		}

		fmt.Println(countOneBits(arrBitwiseOr))
	}
}

func countOneBits(num uint) (numOnes uint) {
	bitOffset := uint(1)
	for i := 0; i < 32; i++ {
		if (num & bitOffset) != 0 {
			numOnes++
		}
		bitOffset <<= 1
	}
	return
}

func readUint() (store uint) {
	fmt.Scan(&store)
	return
}
