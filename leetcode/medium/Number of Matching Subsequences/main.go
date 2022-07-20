package medium

func numMatchingSubseq(S string, words []string) (num int) {
	waitingWords := map[rune][]string{' ': words}

	for _, ch := range " " + S {
		advancingWords := waitingWords[ch]
		delete(waitingWords, ch)

		for i := 0; i < len(advancingWords); i++ {
			word := &advancingWords[i]
			if len(*word) == 0 {
				num++
			} else {
				ch := rune((*word)[0])
				waitingWords[ch] = append(waitingWords[ch], (*word)[1:])
			}
		}
	}
	return
}
