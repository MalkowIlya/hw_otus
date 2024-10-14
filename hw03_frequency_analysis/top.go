package main

import (
	"regexp"
	"sort"
	"strings"
)

type Words struct {
	word  string
	count int
}

var reg = regexp.MustCompile(`(-{2,})|([а-яА-Я]+-?[a-яА-Я]?)`)

func Top10(text string) []string {
	var result []string
	wordList := strings.Fields(text)
	words := make([]Words, 0, len(wordList))
	wordMap := make(map[string]int, len(wordList))

	for _, word := range wordList {
		if word := reg.FindString(word); word != "" {
			wordMap[strings.ToLower(word)]++
		}
	}

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

	if len(words) >= 10 {
		words = words[:10]

		for _, word := range words {
			result = append(result, word.word)
		}
	}

	return result
}
