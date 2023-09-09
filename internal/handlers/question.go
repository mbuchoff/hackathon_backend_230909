package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mbuchoff/hackathon_backend_230909/internal/dto"
	"github.com/mbuchoff/hackathon_backend_230909/internal/services/translate"
)

func AnswerQuestion(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Set the content type header
		w.Header().Set("Content-Type", "application/json")
		// allow CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Get the request body
		defer r.Body.Close() // Close the body when we're done
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(dto.ResponseError{Message: "Internal Server Error"})
			return
		}
		fmt.Println(string(body))
		// extract the phrase from the request body and attach it to the payload and send to transtate the text
		textToBeTranslated := dto.Payload{Text: string(body)}

		// translate the phrase using our function translateTexts and return the translated phrase
		translatedPhrase, err := translate.TranslateText(textToBeTranslated.Text, "pt")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(dto.ResponseError{Message: "Internal Server Error"})
			return
		}
		fmt.Println("Translated phrase:", translatedPhrase)

		// Write the response body
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(dto.Response{Message: translatedPhrase})

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(dto.ResponseError{Message: "Method not allowed"})
	}
}
