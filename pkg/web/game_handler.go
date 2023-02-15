package web

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/WilliamKSilva/type-1v1/pkg/api"
)

type GameHandler struct {
    gameService api.GameServiceInterface 
}

func NewGameHandler (gameService api.GameServiceInterface) *GameHandler {
    return &GameHandler{gameService}
}

func (g *GameHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        g.NewGameFunc(w, r)
    }

    if r.Method == "PUT" {
        g.UpdateGameFunc(w, r)
    }
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

func (g *GameHandler) UpdateGameFunc (w http.ResponseWriter, r *http.Request) {

    updateGameData := &api.UpdateGameData{}
    id := r.URL.Query().Get("id")

    body := r.Body
    data, err := io.ReadAll(body)

    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte(err.Error()))
    }

    json.Unmarshal(data, updateGameData) 

    u64, err := strconv.Atoi(id)

    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte(err.Error()))
    }

    game, err := g.gameService.UpdateGame(uint(u64), *updateGameData)

    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte(err.Error()))
    }

    resp, err := json.Marshal(game)

    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte(err.Error()))
    }

    w.WriteHeader(200)
    w.Write(resp)
}
