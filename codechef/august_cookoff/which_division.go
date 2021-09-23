package august_cookoff

import "fmt"

func main() {
	var num_test_cases uint
	fmt.Scan(&num_test_cases)

	for i := uint(0); i < num_test_cases; i++ {
		var R uint
		fmt.Scan(&R)

		if R >= 2000 {
			fmt.Println(1)
		} else if R >= 1600 {
			fmt.Println(2)
		} else {
			fmt.Println(3)
		}
	}
}
