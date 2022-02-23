package main

func maxScoreIndices(nums []int) []int {
	numZerosInLeft, numOnesInRight := 0, 0
	for _, num := range nums {
		if num == 1 {
			numOnesInRight += 1
		}
	}

	maxScore, maxScoreRecord := numOnesInRight+numZerosInLeft, []int{0}

	for idx, num := range nums {
		if num == 0 {
			numZerosInLeft += 1
		} else if num == 1 {
			numOnesInRight -= 1
		}

		if score := numOnesInRight + numZerosInLeft; score > maxScore {
			maxScore = score
			maxScoreRecord = []int{idx + 1}
		} else if score == maxScore {
			maxScoreRecord = append(maxScoreRecord, idx+1)
		}

	}

	return maxScoreRecord
}
