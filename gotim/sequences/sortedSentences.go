package sequences

import (
	"strconv"
	"strings"
)

func sortSentences(sentence string) string {
	if sentence == "" {
		return ""
	}
	sentenceSplit := strings.Split(sentence, " ")
	sortedSentenceSplit := make([]string, len(sentenceSplit))
	numberMap := map[string]interface{} {
		"1": nil,
		"2": nil,
		"3": nil,
		"4": nil,
		"5": nil,
		"6": nil,
		"7": nil,
		"8": nil,
		"9": nil,
	}

	for _, word := range sentenceSplit {
		for i := 0; i < len(word); i++ {
			letter := word[i:i+1]
			if _, ok := numberMap[letter]; ok {
				index, _ := strconv.ParseInt(letter, 10, 8)
				sortedSentenceSplit[index-1] = word
			}
		}
	}

	return strings.Join(sortedSentenceSplit, " ")
}