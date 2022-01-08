package main

const SPACE = ' '

func capitalizeTitle(title string) string {
	titleBytes := []byte(title)
	titleBytes = append(titleBytes, SPACE)

	prevWordStartFrom := 0
	totalChars := 0
	for i := 0; i < len(titleBytes); i++ {
		if titleBytes[i] == SPACE {
			if totalChars > 2 {
				titleBytes[prevWordStartFrom] -= 32
			}
			prevWordStartFrom = i + 1
			totalChars = 0
			continue
		}

		if 65 <= titleBytes[i] && titleBytes[i] <= 90 {
			titleBytes[i] += 32
		}
		totalChars++
	}

	return string(titleBytes[:len(titleBytes)-1])
}
