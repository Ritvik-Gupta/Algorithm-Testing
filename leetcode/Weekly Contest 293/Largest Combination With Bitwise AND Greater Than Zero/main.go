package main

func largestCombination(candidates []int) (maxCombinationSize int) {
	for bitPos, bitOffset := 0, 1; bitPos < 32; bitPos, bitOffset = bitPos+1, bitOffset<<1 {
		combinationSize := 0
		for _, num := range candidates {
			if num&bitOffset != 0 {
				combinationSize++
			}
		}
		if combinationSize > maxCombinationSize {
			maxCombinationSize = combinationSize
		}
	}
	return
}
