package main

import (
	"fmt"
)

func main() {
	for testCase := readUint(); testCase > 0; testCase-- {
		num := readUint()
		fmt.Println("A", computeNumFactors(num))
		fmt.Println("Y", algorithm(num))
	}
}

func algorithm(num uint) uint {
	dp := make([]uint, 0, num+1)
	dp = append(dp, 0, 1, 1)

	for i := uint(3); i <= num; i++ {
		dp = append(dp, dp[i-1]+1-computeNumFactors(i))
	}

	return dp[num]
}

func computeNumFactors(num uint) uint {
	numFactors := uint(0)
	for i := uint(2); i <= num; i++ {
		if num%i == 0 {
			numFactors++
		}
	}
	return numFactors
}

func readUint() (store uint) {
	fmt.Scan(&store)
	return
}
