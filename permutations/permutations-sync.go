package permutations

func permuteRecursive(prevNums, leftNums []int) [][]int {
	if len(leftNums) == 0 {
		return [][]int{prevNums}
	}
	result := [][]int{}

	for i, stored := range leftNums {
		newLeftNums := make([]int, i)
		copy(newLeftNums, leftNums[:i])

		result = append(
			result,
			permuteRecursive(append(prevNums, stored), append(newLeftNums, leftNums[i+1:]...))...,
		)
	}
	return result
}

func PermuteSync(nums []int) [][]int {
	return permuteRecursive([]int{}, nums)
}
