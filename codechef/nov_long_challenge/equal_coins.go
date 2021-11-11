package nov_long_challenge

import "fmt"

func equal_coins() {
	var num_test_cases uint
	fmt.Scan(&num_test_cases)

	for i := uint(0); i < num_test_cases; i++ {
		var X, Y uint
		fmt.Scan(&X, &Y)
		if X%2 == 0 && Y%2 == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
