package main

import (
	"fmt"
)

type idendity struct{}

func findTotalPrimeFactors(num uint) uint {
	possibleFactor := uint(2)
	allPrimeFactors := make(map[uint]idendity)
	for num > 1 {
		if num%possibleFactor == 0 {
			num /= possibleFactor
			allPrimeFactors[possibleFactor] = idendity{}
		} else {
			possibleFactor++
		}
	}
	return uint(len(allPrimeFactors))
}

func main() {
	for testCase := readUint(); testCase > 0; testCase-- {
		width, height := readUint(), readUint()
		maxColumnDivision := findTotalPrimeFactors(height)

		for ; maxColumnDivision > 0 && width%maxColumnDivision != 0; maxColumnDivision-- {
		}

		fmt.Println(maxColumnDivision)
	}
}

func readUint() (store uint) {
	fmt.Scan(&store)
	return
}
