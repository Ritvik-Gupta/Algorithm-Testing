package nov_long_challenge

import "fmt"

func weird_palindrome_making() {
	var num_test_cases uint
	fmt.Scan(&num_test_cases)

	for i := uint(0); i < num_test_cases; i++ {
		var N uint
		fmt.Scan(&N)

		var x uint
		numOdd := uint(0)
		for i := uint(0); i < N; i++ {
			fmt.Scan(&x)
			if x%2 != 0 {
				numOdd++
			}
		}

		fmt.Println(numOdd / 2)

	}
}
