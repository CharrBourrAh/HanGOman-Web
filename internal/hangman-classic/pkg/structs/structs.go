package structs

type HangManData struct {
	Word         string   `json:"word"`     // Word composed of '_', ex: H_ll_
	ToFind       string   `json:"toFind"`   // Final word chosen by the program at the beginning. It is the word to find
	Attempts     int      `json:"attempts"` // Number of attempts left
	StrAttempts  string   `json:"strAttempts"`
	WordFile     string   `json:"wordFile"`
	AlreadyTried []string `json:"alreadyTried"`
	Nickname     string   `json:"nickname"`
	Input        string   `json:"input"`
}
