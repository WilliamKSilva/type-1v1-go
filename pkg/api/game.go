package api

import "errors"

const (
    Waiting string = "waiting"
    Started string = "started"
    Finished string = "finished"
)

type Game struct {
    ID uint `json:"id"` 
    PlayerOne string `json:"player_one"`
    PlayerTwo string `json:"player_two"`
    Status string `json:"status"`
    Text string `json:"text"`
    Timer uint `json:"timer"`
}

type NewGameData struct {
    PlayerOne string `json:"player_one"` 
    Timer uint `json:"timer"`
}

type GameRepository interface {
    Create (game *Game) error
}

type GameServiceInterface interface {
    NewGame(data NewGameData) (*Game, error)
}

type gameService struct {
    repo GameRepository
}

func NewGameService (repo GameRepository) *gameService {
    return &gameService{repo}
}

func (g *gameService) NewGame(data NewGameData) (*Game, error) {
    if (data.PlayerOne == "") {
        return nil, errors.New("Player one name is required")
    }

    if (data.Timer == 0) {
        return nil, errors.New("Timer is required")
    }

    game := &Game{
        ID: 0,
        PlayerOne: data.PlayerOne,
        PlayerTwo: "",
        Status: Waiting,
        Timer: data.Timer,
        Text: "",
    }

    err := g.repo.Create(game) 

    if err != nil {
        return nil, err
    }

    return game, err
}
