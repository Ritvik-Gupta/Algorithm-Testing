package main

import "fmt"

func main() {
	var numTestCases uint

	fmt.Scan(&numTestCases)
	for testCase := numTestCases; testCase > 0; testCase-- {
		var n uint
		fmt.Scan(&n)

		var elm int
		uniqueElements := make(map[int]struct{}, n)
		for i := n; i > 0; i-- {
			fmt.Scan(&elm)
			if _, containsElm := uniqueElements[elm]; !containsElm {
				uniqueElements[elm] = struct{}{}
			} else if _, containsElm := uniqueElements[-elm]; !containsElm {
				uniqueElements[-elm] = struct{}{}
			}
		}

		fmt.Println(len(uniqueElements))
	}
}
