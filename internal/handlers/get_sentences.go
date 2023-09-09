package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mbuchoff/hackathon_backend_230909/internal/dto"
	"github.com/mbuchoff/hackathon_backend_230909/internal/services/feed"
)

func GetEnglishSentencesHandler(w http.ResponseWriter, r *http.Request) {

	result, err := feed.GetEnglishSentences()
	if err != nil {
		fmt.Println("Error getting English sentences:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.FeedResult{Error: err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(result)
}
