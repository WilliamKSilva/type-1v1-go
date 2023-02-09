package api

import "errors"

type Game struct {
    ID uint `json:"id"` 
    PlayerOne string `json:"player_one"`
    PlayerTwo string `json:"player_two"`
    Text string `json:"text"`
    Timer uint `json:"timer"`
}

type newGameData struct {
    PlayerOne string `json:"player_one"` 
    Timer uint `json:"timer"`
}

type GameRepository interface {
    Create (game *Game) error
}

type gameService struct {
    repo GameRepository
}

func (g *gameService) NewGame(data newGameData) (*Game, error) {
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
        Timer: data.Timer,
        Text: "",
    }

    err := g.repo.Create(game) 

    if err != nil {
        return nil, err
    }

    return game, err
}
