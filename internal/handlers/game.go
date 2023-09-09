package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mbuchoff/hackathon_backend_230909/internal/dto"
	"github.com/mbuchoff/hackathon_backend_230909/internal/services/game"
)

func GameHandler(w http.ResponseWriter, r *http.Request) {
	// receive in the body the number of questions and the language of the questions
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		// allow CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")

		defer r.Body.Close() // Close the body when we're done
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(dto.ResponseError{Message: "Internal Server Error"})
			return
		}
		fmt.Println(string(body))

		// parse to the data type game options dto
		gameOptions := dto.GameOptions{}
		err = json.Unmarshal(body, &gameOptions)
		if err != nil {
			fmt.Println("Error unmarshaling game options:", err)
		}

		var gameResults []dto.Question

		for i := 0; i < gameOptions.Quantity; i++ {
			game, err := game.GameCreation(gameOptions.Language)
			if err != nil {
				fmt.Println("Error creating game:", err)
				return
			}
			gameResults = append(gameResults, game)
		}
		fmt.Println(gameResults)
		// Write the response body
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(dto.ResponseGame{Message: "OK", GameResult: dto.GameResult{Questions: gameResults}})

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
