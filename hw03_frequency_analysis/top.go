package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type Words struct {
	word  string
	count int
}

var reg = regexp.MustCompile(`(-{2,})|([а-яА-Яa-zA-Z]+-?[a-яА-Яa-zA-Z]?)`)

func Top10(text string) []string {
	result := make([]string, 0, 10)
	wordList := strings.Fields(text)
	wordMap := make(map[string]int, len(wordList))

	for _, word := range wordList {
		if word := reg.FindString(word); word != "" {
			wordMap[strings.ToLower(word)]++
		}
	}
	words := make([]Words, 0, len(wordMap))

	for word, count := range wordMap {
		words = append(words, Words{word: word, count: count})
	}

	sort.Slice(words, func(i, j int) bool {
		if words[i].count > words[j].count {
			return true
		}
		if words[i].count == words[j].count {
			return words[i].word < words[j].word
		}
		return false
	})

	splice := len(words)
	if splice > 10 {
		splice = 10
	}
	words = words[:splice]
	for _, word := range words {
		result = append(result, word.word)
	}

	return result
}
