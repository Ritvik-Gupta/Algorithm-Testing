package main

type frequencyCollection = [26]int

func removeAnagrams(words []string) []string {
	result := make([]string, 0)
	prevWordCollection := frequencyCollection{}

	for _, word := range words {
		wordCollection := frequencyCollection{}

		for _, char := range word {
			wordCollection[char-'a']++
		}

		for i := 0; i < 26; i++ {
			if prevWordCollection[i] != wordCollection[i] {
				result = append(result, word)
				prevWordCollection = wordCollection
				break
			}
		}
	}
	return result
}
