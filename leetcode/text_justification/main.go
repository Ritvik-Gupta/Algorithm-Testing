package text_justification

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	result := make([]string, 0, len(words))

	lineStartIdx, lineLength := 0, 0
	for wordIdx, word := range words {
		if lineLength+len(word) > maxWidth {
			result = append(
				result,
				fullyJustifyLine(words[lineStartIdx:wordIdx], maxWidth, lineLength+lineStartIdx-wordIdx),
			)
			lineStartIdx, lineLength = wordIdx, 0
		}
		lineLength += len(word) + 1
	}
	result = append(
		result,
		leftJustifyLine(words[lineStartIdx:], maxWidth, lineLength+lineStartIdx-len(words)),
	)

	return result
}

const SPACE = ' '

func leftJustifyLine(words []string, maxWidth, totalWordsLen int) string {
	var result strings.Builder
	result.Grow(maxWidth)
	totalSpacesLeft := maxWidth - totalWordsLen

	for i := 0; i < len(words)-1; i++ {
		result.WriteString(words[i])
		result.WriteByte(SPACE)
		totalSpacesLeft -= 1
	}

	result.WriteString(words[len(words)-1])
	for i := 0; i < totalSpacesLeft; i++ {
		result.WriteByte(SPACE)
	}

	return result.String()
}

func fullyJustifyLine(words []string, maxWidth, totalWordsLen int) string {
	var result strings.Builder
	result.Grow(maxWidth)
	totalSpacesLeft := maxWidth - totalWordsLen

	for idx, word := range words {
		result.WriteString(word)
		totalWordsLeft := len(words) - 1 - idx

		spacesTaken := totalSpacesLeft
		if totalWordsLeft != 0 {
			spacesTaken = ceilDivision(totalSpacesLeft, totalWordsLeft)
		}

		for i := 0; i < spacesTaken; i++ {
			result.WriteByte(SPACE)
		}
		totalSpacesLeft -= spacesTaken
	}

	return result.String()
}

func ceilDivision(a, b int) int {
	result := a / b
	if a%b == 0 {
		return result
	} else {
		return result + 1
	}
}
