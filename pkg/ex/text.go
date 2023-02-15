package ex 

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TextServiceInterface interface {
	GetRandomText(trigger string) (string, error)
}

type WordResult struct {
	Word  string `json:"word"`
	Score uint   `json:"score"`
}

var AvaiableWordTriggers = [4]string{"apple", "cow", "computer", "games"}

type textService struct{}

func NewTextService () *textService {
    return &textService{}
}

func (t *textService) GetRandomText(trigger string) (string, error) {
	var text string

	wordsResult := &[]WordResult{}

	resp, err := http.Get(fmt.Sprintf("https://api.datamuse.com/words?rel_trg=%s", trigger))

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	json.Unmarshal(body, wordsResult)

	for _, r := range *wordsResult {
		text = text + fmt.Sprintf(" %s", r.Word)
	}

	return text, nil
}
