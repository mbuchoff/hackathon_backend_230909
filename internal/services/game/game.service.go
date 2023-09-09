package game

import (
	"fmt"
	"math/rand"

	"github.com/mbuchoff/hackathon_backend_230909/internal/dto"
	"github.com/mbuchoff/hackathon_backend_230909/internal/services/feed"
	"github.com/mbuchoff/hackathon_backend_230909/internal/services/translate"
)

// logic to game

// receive the language and than call the sentences service

func GameCreation(language string) (dto.Question, error) {
	sentences, err := feed.GetEnglishSentences()
	if err != nil {
		fmt.Println("Error getting sentences:", err)
		return dto.Question{}, err
	}
	correctOption := rand.Intn(4)

	untranslatedOption := sentences.Sentences[correctOption]

	for i := 0; i < 4; i++ {
		translatedOption, err := translate.TranslateText(sentences.Sentences[i], language)
		if err != nil {
			fmt.Println("Error translating sentence:", err)
			return dto.Question{}, err
		}
		sentences.Sentences[i] = translatedOption

	}

	question := dto.Question{
		Question: untranslatedOption,
		Choices:  sentences.Sentences,
		Answer:   correctOption,
	}
	return question, nil
}
