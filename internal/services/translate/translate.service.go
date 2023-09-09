package translate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	apiTranslationURL = "https://api.cognitive.microsofttranslator.com/translate?api-version=3.0&from=en&to="
	subscriptionKey   = "138aff8d95b748468016c66926bb09c7"
	location          = "eastus"
)

type TranslationResponse struct {
	Translations []struct {
		Text string `json:"text"`
		To   string `json:"to"`
	} `json:"translations"`
}

func TranslateText(text string, language string) (string, error) {

	payload := []struct {
		Text string `json:"Text"`
	}{
		{Text: text},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("Error marshaling JSON payload: %v", err)
	}

	req, err := http.NewRequest("POST", apiTranslationURL+language, bytes.NewReader(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("Error creating HTTP request: %v", err)
	}

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Ocp-Apim-Subscription-Key", subscriptionKey)
	req.Header.Add("Ocp-Apim-Subscription-Region", location)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response body: %v", err)
	}

	var translationResponse []TranslationResponse
	err = json.Unmarshal(body, &translationResponse)
	if err != nil {
		return "", fmt.Errorf("Error unmarshaling JSON response: %v", err)
	}

	if len(translationResponse) > 0 && len(translationResponse[0].Translations) > 0 {
		fmt.Println("Translation:", translationResponse[0].Translations[0].Text)
		return translationResponse[0].Translations[0].Text, nil
	}

	return "", fmt.Errorf("Translation not found in the response")
}
