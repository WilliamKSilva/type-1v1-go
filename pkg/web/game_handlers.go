package web

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/WilliamKSilva/type-1v1/pkg/api"
)

type GameHandler struct {
    gameService api.GameService 
}

func (g *GameHandler) NewGameHandler (w http.ResponseWriter, r *http.Request) {
    gameData := &api.NewGameData{}
    body := r.Body
    jsonData, err := io.ReadAll(body)

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write(nil) 
    }

    err = json.Unmarshal(jsonData, gameData) 

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write(nil) 
    }

    game, err := g.gameService.NewGame(*gameData)

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
    }

    response, err := json.Marshal(*game)

    w.WriteHeader(http.StatusOK)
    w.Write(response)
}
