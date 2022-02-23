package main

import "sort"

func binarySearch(val int, nums []int, lowerIdx, upperIdx int) int {
	for lowerIdx <= upperIdx {
		if midIdx := (lowerIdx + upperIdx) / 2; nums[midIdx] < val {
			lowerIdx = midIdx + 1
		} else if nums[midIdx] > val {
			upperIdx = midIdx - 1

		} else {
			return midIdx
		}
	}
	return -1
}

func findFinalValue(nums []int, original int) int {
	sort.Ints(nums)

	lowerIdx, upperIdx := 0, len(nums)-1
	for {
		if idx := binarySearch(original, nums, lowerIdx, upperIdx); idx == -1 {
			return original
		} else {
			original *= 2
			lowerIdx = idx + 1
		}
	}
}
