package web

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/WilliamKSilva/type-1v1/pkg/api"
)

type GameHandler struct {
    gameService api.GameServiceInterface 
}

func NewGameHandler (gameService api.GameServiceInterface) *GameHandler {
    return &GameHandler{gameService}
}

func (g *GameHandler) NewGameFunc (w http.ResponseWriter, r *http.Request) {
    gameData := &api.NewGameData{}
    jsonData, err := io.ReadAll(r.Body)

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write(nil) 
    }

    err = json.Unmarshal(jsonData, gameData) 

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error())) 
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
