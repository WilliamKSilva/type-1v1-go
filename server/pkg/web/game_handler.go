package web

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/WilliamKSilva/type-1v1/pkg/api"
	"github.com/gorilla/websocket"
)

type GameHandler struct {
    gameService api.GameServiceInterface 
}

type RunGameData struct {
    Player string `json:"player"` 
    Text string `json:"text"` 
    Id string `json:"id"`
}

type SocketResponse struct {
    Message string `json:"message"`
}

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
}

func NewGameHandler (gameService api.GameServiceInterface) *GameHandler {
    return &GameHandler{gameService}
}

func (g *GameHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")

    fmt.Println(r)

    if r.Method == "POST" {
        g.NewGameFunc(w, r)
        return
    }

    if r.Method == "PUT" {
        g.UpdateGameFunc(w, r)
        return
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
        return
    }

    json.Unmarshal(data, updateGameData) 

    u64, err := strconv.Atoi(id)

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
        return
    }

    game, err := g.gameService.UpdateGame(uint(u64), *updateGameData)

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
        return
    }

    resp, err := json.Marshal(game)

    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(resp)
}


func (g *GameHandler) RunGameFunc(w  http.ResponseWriter, r *http.Request) {
    ctx := context.Background()
    ctx, cancel := context.WithTimeout(ctx, time.Minute * 4)

    defer cancel()

    conn, err := upgrader.Upgrade(w, r, nil)

    if err != nil {
        log.Println(err)
        return
    }

    g.SocketMessageReceiver(ctx, conn)

    return
}

func (g *GameHandler) SocketMessageReceiver (ctx context.Context, conn *websocket.Conn) {
    for {
        select {
        case <- ctx.Done():
            json := &SocketResponse{Message: "time expired"}
            conn.WriteJSON(json)
            conn.Close()

            return
        default:
            runGameData := &RunGameData{} 
            conn.ReadJSON(runGameData)

            convertedId, _ := strconv.Atoi(runGameData.Id)

            game := g.gameService.RunGame(runGameData.Player, uint(convertedId), runGameData.Text)

            if game != nil {
                conn.WriteJSON(game)
                conn.Close()

                return
            }
        }
    } 
}
