package web

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/WilliamKSilva/type-1v1/pkg/api"
	"github.com/gorilla/websocket"
)

type GameHandler struct {
    gameService api.GameServiceInterface 
}

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
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
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
    }

    json.Unmarshal(data, updateGameData) 

    u64, err := strconv.Atoi(id)

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
    }

    game, err := g.gameService.UpdateGame(uint(u64), *updateGameData)

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
    }

    resp, err := json.Marshal(game)

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
    }

    w.WriteHeader(http.StatusOK)
    w.Write(resp)
}

type runGameData struct {
    Player string `json:"player"` 
    Text string `json:"text"` 
    Id string `json:"id"`
}

func (g *GameHandler) RunGameFunc(w  http.ResponseWriter, r *http.Request) {
    ctx := context.Background()
    ctx, cancel := context.WithTimeout(ctx, time.Minute * 3)
    defer cancel()

    clientData := &runGameData{}
    
    conn, err := upgrader.Upgrade(w, r, nil)

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
    }

    conn.ReadJSON(clientData)

    u64, err := strconv.Atoi(clientData.Id)

    gamechan := make(chan *api.Game)
    go g.gameService.RunGame(clientData.Player, uint(u64), clientData.Text, gamechan)

    select {
    case <- gamechan:
        json, _ := json.Marshal(gamechan)

        w.WriteHeader(http.StatusOK)
        w.Write(json)
    case <- ctx.Done():
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Timeout exceeded"))
    }

}
