package dto

type Response struct {
	Message string `json:"message"`
}

type ResponseError struct {
	Message string `json:"message"`
}

type Payload struct {
	// TODO: Phrases to be translated
	Text string `json:"Text"`
}

type FeedResult struct {
	Sentences []string `json:"sentences"`
	Error     string   `json:"error,omitempty"`
}

type GameOptions struct {
	Language string `json:"language"`
	Quantity int    `json:"quantity"`
}

type GameResult struct {
	Questions []Question `json:"questions"`
}

type Question struct {
	Question string   `json:"question"`
	Choices  []string `json:"choices"`
	Answer   int      `json:"answer"` // index of the correct answer
}

type ResponseGame struct {
	Message    string     `json:"message"`
	GameResult GameResult `json:"game_result"`
}
