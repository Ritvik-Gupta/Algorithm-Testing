package main

import (
	"fmt"
)

const SET_BIT = uint32(1)

func main() {
	for testCase := readUint(); testCase > 0; testCase-- {
		nums, i := bitwiseCreateFromXor(readUint(), readUint()), 0
		for ; i < len(nums)-1; i++ {
			fmt.Print(nums[i], " ")
		}
		fmt.Print(nums[i])
	}
}

func bitwiseCreateFromXor(totalNums, xorValue uint32) []uint32 {
	if totalNums == 1 {
		return []uint32{xorValue}
	}

	nums := make([]uint32, totalNums)
	for i := uint32(0); xorValue != 0 || i < totalNums; i++ {
		offsetBit := SET_BIT << i
		hasEvenOnes := (xorValue & SET_BIT) != SET_BIT

		if hasEvenOnes {
			nums[(i+1)%totalNums] += offsetBit
		}
		nums[i%totalNums] += offsetBit

		xorValue >>= SET_BIT
	}
	return nums
}

func readUint() (store uint32) {
	fmt.Scan(&store)
	return
}
