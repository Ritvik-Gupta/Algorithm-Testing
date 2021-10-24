package length_of_last_word

func lengthOfLastWord(s string) int {
	wordLen := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			wordLen++
		} else if wordLen != 0 {
			break
		}
	}
	return wordLen
}
