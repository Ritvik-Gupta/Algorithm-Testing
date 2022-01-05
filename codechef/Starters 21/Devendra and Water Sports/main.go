package main

import "fmt"

func main() {
	for testCase := readUint(); testCase > 0; testCase-- {
		moneyLeft := readUint() - readUint()
		totalCost := readUint() + readUint() + readUint()
		if moneyLeft >= totalCost {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func readUint() (store uint) {
	fmt.Scan(&store)
	return
}
