package feed

import (
	"math/rand"
	"strings"
	"time"

	"github.com/mbuchoff/hackathon_backend_230909/internal/dto"
	"github.com/mmcdole/gofeed"
)

const rssURL = "https://www.ted.com/feeds/talks.rss"

// GetEnglishSentences returns the first four English sentences from the RSS feed
func GetEnglishSentences() (*dto.FeedResult, error) {

	// analyze the RSS feed
	parser := gofeed.NewParser()

	// parse the RSS feed
	feed, err := parser.ParseURL(rssURL)
	if err != nil {
		return nil, err
	}

	var englishSentences []string
	for _, item := range feed.Items {
		if isEnglishSentence(item.Title) {
			englishSentences = append(englishSentences, strings.Split(item.Title, "|")[0])
		}
	}

	// Embaralhe os índices das sentenças
	rand.Seed(time.Now().UnixNano())
	indexes := rand.Perm(len(englishSentences))

	// Selecione quatro sentenças aleatórias
	numSentences := 4
	if len(englishSentences) < numSentences {
		numSentences = len(englishSentences)
	}

	selectedSentences := make([]string, numSentences)
	for i := 0; i < numSentences; i++ {
		selectedSentences[i] = englishSentences[indexes[i]]
	}

	result := &dto.FeedResult{Sentences: selectedSentences}
	return result, nil
}

// isEnglishSentence verifies if a sentence is in English
func isEnglishSentence(sentence string) bool {
	for _, char := range sentence {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			return true
		}
	}
	return false
}
