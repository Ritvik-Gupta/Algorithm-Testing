package main

import (
	"fmt"
)

func main() {
	for testCase := readUint(); testCase > 0; testCase-- {
		n, x := readUint(), readUint()
		records := make(map[uint]uint, n)

		for i := uint(0); i < n; i++ {
			val := readUint()
			records[val]++
		}

		for loopAgain := true; loopAgain; loopAgain = false {
			for key := range records {
				aKey := key * x
				if _, ok := records[aKey]; ok {
					records[key]--
					if records[key] == 0 {
						delete(records, key)
					}
					records[aKey]--
					if records[aKey] == 0 {
						delete(records, aKey)
					}
					loopAgain = true
					continue
				}

				bKey := key / x
				if bKey*x != key {
					continue
				}
				if _, ok := records[bKey]; ok {
					records[key]--
					if records[key] == 0 {
						delete(records, key)
					}
					records[bKey]--
					if records[key] == 0 {
						delete(records, bKey)
					}
					loopAgain = true
					continue
				}
			}
		}

		count := uint(0)
		for _, freq := range records {
			count += freq
		}

		fmt.Println(count)
	}
}

func readUint() (store uint) {
	fmt.Scan(&store)
	return
}
