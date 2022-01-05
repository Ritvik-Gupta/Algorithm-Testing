package main

import "fmt"

const CODE = "code"
const CHEF = "chef"

func main() {
	for testCase := readInt(); testCase > 0; testCase-- {
		if checkForCodeAndChef() {
			fmt.Println("AC")
		} else {
			fmt.Println("WA")
		}
	}
}

func checkForCodeAndChef() bool {
	n := readInt()
	var s string
	fmt.Scan(&s)

	codeOccurence, chefOccurence := -1, -1
	for idx := 0; idx < n && (codeOccurence == -1 || chefOccurence == -1); idx++ {
		if s[idx] == 'c' {
			if s[idx:idx+4] == CODE && codeOccurence == -1 {
				codeOccurence = idx
			} else if s[idx:idx+4] == CHEF && chefOccurence == -1 {
				if codeOccurence == -1 {
					return false
				}
				chefOccurence = idx
			}
		}
	}

	return codeOccurence < chefOccurence
}

func readInt() (store int) {
	fmt.Scan(&store)
	return
}
