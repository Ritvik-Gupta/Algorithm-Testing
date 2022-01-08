package main

func longestPalindrome(words []string) int {
	selfPalindromes, pairPalindromes := 0, 0
	uniqueWords := make(map[string]int)

	for _, word := range words {
		reversedWord := reverse(word)

		if prevCounter, ok := uniqueWords[reversedWord]; ok {
			if prevCounter == 1 {
				delete(uniqueWords, reversedWord)
			} else {
				uniqueWords[reversedWord]--
			}
			pairPalindromes++

			if word == reversedWord {
				selfPalindromes--
			}
		} else {
			uniqueWords[word]++

			if word == reversedWord {
				selfPalindromes++
			}
		}
	}

	result := pairPalindromes * 4
	if selfPalindromes > 0 {
		result += 2
	}
	return result
}

func reverse(word string) string {
	return string(word[1]) + string(word[0])
}
