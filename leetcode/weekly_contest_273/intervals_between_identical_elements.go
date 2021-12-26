package weekly_contest_273

func getDistances(arr []int) []int64 {
	frequency := make(map[int]int)
	idxHistory := make(map[int]int64)

	for idx, elm := range arr {
		idxHistory[elm] += int64(idx)
		frequency[elm] += 1
	}

	results := make([]int64, 0, len(arr))
	for idx, elm := range arr {
		cumulativeResult := idxHistory[elm] - int64(frequency[elm])*int64(idx)
		if cumulativeResult < 0 {
			cumulativeResult *= -1
		}
		results = append(results, cumulativeResult)

		frequency[elm] -= 2
		idxHistory[elm] -= int64(idx) * 2
	}
	return results
}
