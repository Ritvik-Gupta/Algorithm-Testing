package main

func minSwaps(nums []int) int {
	windowSize := 0
	for _, num := range nums {
		if num == 1 {
			windowSize++
		}
	}

	if windowSize == 0 {
		return 0
	}

	numZeroes := 0
	window := make([]int, 0, windowSize)
	for i := 0; i < windowSize; i++ {
		window = append(window, nums[i])
		if nums[i] == 0 {
			numZeroes++
		}
	}
	minZeroes := numZeroes

	for i := 0; i < len(nums)+1; i++ {
		if minZeroes > numZeroes {
			minZeroes = numZeroes
		}

		if window[0] == 0 {
			numZeroes--
		}
		window = window[1:]
		window = append(window, nums[(i+windowSize)%len(nums)])
		if window[len(window)-1] == 0 {
			numZeroes++
		}
	}

	return minZeroes
}
