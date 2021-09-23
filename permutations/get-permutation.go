package permutations

import "strings"

func factorial(n int) int {
	fact := 1
	for i := 2; i <= n; i++ {
		fact *= i
	}
	return fact
}

func FindKthPermutation(nums []rune, k int) string {
	if len(nums) == 0 {
		return ""
	}

	numSize, permutation := len(nums), &strings.Builder{}
	i, headPermutations := 0, factorial(numSize-1)

	for ; i < numSize && k-headPermutations > 0; i++ {
		k -= headPermutations
	}

	if i == numSize {
		panic("Range out of Bounds of Total Possible Permutations")
	}
	permutation.WriteRune(nums[i])
	permutation.WriteString(FindKthPermutation(append(nums[0:i], nums[i+1:]...), k))
	return permutation.String()
}

func GetPermutation(n int, k int) string {
	nums := []rune{}
	for i, char := 0, '1'; i < n; i, char = i+1, char+1 {
		nums = append(nums, char)
	}
	return FindKthPermutation(nums, k)
}
