package main

import "fmt"

func main() {
	// for i := uint(0); i < 10; i++ {
	calculate(90)
	// }

	for testCase := readUint(); testCase > 0; testCase-- {
		fmt.Println(smallestDivisor(readUint()))
	}
}

func calculate(n uint) {
	for i := uint(2); i <= n; i++ {
		if n%i == 0 {
			m := n / i
			c := 0
			for j := uint(1); j <= m; j++ {
				if m%j == 0 {
					c++
				}
			}

			fmt.Println(n, i, m, c)
		}
	}
}

func smallestDivisor(n uint) uint {
	for divisor := uint(2); divisor <= n; divisor++ {
		if n%divisor == 0 {
			return divisor
		}
	}
	return 0
}

func readUint() (store uint) {
	fmt.Scan(&store)
	return
}
