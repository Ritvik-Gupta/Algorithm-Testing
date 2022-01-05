package main

import "fmt"

func main() {
	for testCase := readUint(); testCase > 0; testCase-- {
		n, k := readUint(), readUint()
		var s string
		fmt.Scan(&s)

		s += "0"

		xorValue := 0
		for i := uint(0); i < n-k; i++ {
			if s[i] == '1' {
				xorValue ^= 1
			}
		}

		sum := 0
		for i := uint(0); i < k; i++ {
			if s[i+n-k+1] == '1' {
				xorValue ^= 1
			}
			sum += xorValue
			if s[i] == '1' {
				xorValue ^= 1
			}
		}

		fmt.Println(sum)
	}
}

func readUint() (store uint) {
	fmt.Scan(&store)
	return
}
