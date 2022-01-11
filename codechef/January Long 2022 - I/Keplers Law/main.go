package main

import (
	"fmt"
)

func main() {
	for testCase := readUint(); testCase > 0; testCase-- {
		t1, t2, r1, r2 := readUint(), readUint(), readUint(), readUint()

		if t1*t1*r2*r2*r2 == t2*t2*r1*r1*r1 {
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
