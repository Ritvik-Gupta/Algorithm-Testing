package main

import "fmt"

func main() {
	for testCase := readUint(); testCase > 0; testCase-- {
		nitin, sobhagya := readUint(), readUint()

		for _, extraCoins := range [2]uint{readUint(), readUint()} {
			if nitin >= sobhagya {
				sobhagya += extraCoins
			} else {
				nitin += extraCoins
			}
		}

		if nitin >= sobhagya {
			fmt.Printf("%c\n", 'N')
		} else {
			fmt.Printf("%c\n", 'S')
		}
	}
}

func readUint() (store uint) {
	fmt.Scan(&store)
	return
}
