package nov_long_challenge

import "fmt"

func main() {
	var num_test_cases uint
	fmt.Scan(&num_test_cases)

	for i := uint(0); i < num_test_cases; i++ {
		var X, Y, A, B, K uint
		fmt.Scan(&X, &Y, &A, &B, &K)
		petrol, diesel := X+K*A, Y+K*B
		if petrol > diesel {
			fmt.Println("DIESEL")
		} else if petrol < diesel {
			fmt.Println("PETROL")
		} else {
			fmt.Println("SAME PRICE")
		}
	}
}
