package plus_one

func plusOne(digits []int) []int {
	idx, hasCarry := len(digits)-1, true
	for hasCarry {
		if idx == -1 {
			digits = append([]int{1}, digits...)
			hasCarry = false
		} else if digits[idx] == 9 {
			digits[idx] = 0
		} else {
			digits[idx]++
			hasCarry = false
		}
		idx--
	}
	return digits
}
