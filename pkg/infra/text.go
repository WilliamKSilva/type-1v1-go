package infra

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WordResult struct {
	Word  string `json:"word"`
	Score uint   `json:"score"`
}

var AvaiableWordTriggers = [4]string{"apple", "cow", "computer", "games"}

func GetRelatedWords(trigger string) ([]string, error) {
	var words []string

    wordsResult := &[]WordResult{} 

	resp, err := http.Get(fmt.Sprintf("https://api.datamuse.com/words?rel_trg=%s", trigger))

    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)

    json.Unmarshal(body, wordsResult)

    for i, w := range *wordsResult {
       words[i] = w.Word 
    }

    return words, nil
}
