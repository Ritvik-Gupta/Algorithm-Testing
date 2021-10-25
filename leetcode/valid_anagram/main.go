package valid_anagram

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charFreq := make(map[rune]uint)
	for _, ch := range s {
		charFreq[ch]++
	}

	for _, ch := range t {
		if freq, ok := charFreq[ch]; !ok {
			return false
		} else if freq == 1 {
			delete(charFreq, ch)
		} else {
			charFreq[ch] = freq - 1
		}
	}

	return len(charFreq) == 0
}
