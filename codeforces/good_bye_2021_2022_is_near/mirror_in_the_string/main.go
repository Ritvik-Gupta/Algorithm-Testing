package main

import (
	"fmt"
	"strings"
)

func main() {
	var numTestCases uint
	fmt.Scan(&numTestCases)

	for testCase := numTestCases; testCase > 0; testCase-- {
		var n int
		fmt.Scan(&n)
		var word string
		fmt.Scan(&word)

		i := 0
		for ; i < n-1 && word[i] > word[i+1]; i++ {
		}

		result := strings.Builder{}
		result.Grow((i + 1) * 2)
		result.WriteString(word[0 : i+1])
		for ; i >= 0; i-- {
			result.WriteByte(word[i])
		}

		fmt.Println(result.String())
	}
}
