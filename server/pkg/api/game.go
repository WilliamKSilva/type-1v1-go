package api

import (
	"errors"
	"math/rand"

	"github.com/WilliamKSilva/type-1v1/pkg/ex"
)

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
    Winner string `json:"Winner"`
}

type NewGameData struct {
    PlayerOne string `json:"player_one"` 
}

type UpdateGameData struct {
    PlayerTwo string `json:"player_two"` 
    Status string `json:"status"`
    Winner string `json:"winner"`
}

type GameRepository interface {
    Create (game *Game) error
    Update (id uint, gameData UpdateGameData) (*Game, error)
    Find (id uint) (*Game, error)
}

type GameServiceInterface interface {
    NewGame(data NewGameData) (*Game, error)
    UpdateGame(id uint, data UpdateGameData) (*Game, error)
    RunGame(player string, id uint, text string) (*Game)
}

type gameService struct {
    repo GameRepository
    textService ex.TextServiceInterface
}

func NewGameService (repo GameRepository, textService ex.TextServiceInterface) *gameService {
    return &gameService{repo, textService}
}

func (g *gameService) NewGame(data NewGameData) (*Game, error) {
    if (data.PlayerOne == "") {
        return nil, errors.New("Player one name is required")
    }

    randomWordTrigger := ex.AvaiableWordTriggers[rand.Intn(4)]
    text, err := g.textService.GetRandomText(randomWordTrigger)

    if err != nil {
        return nil, errors.New("Internal Server Error") 
    }

    game := &Game{
        ID: 0,
        PlayerOne: data.PlayerOne,
        PlayerTwo: "",
        Status: Waiting,
        Text: text,
        Winner: "",
    }

    err = g.repo.Create(game) 

    if err != nil {
        return nil, err
    }

    return game, err
}

func (g *gameService) UpdateGame(id uint, data UpdateGameData) (*Game, error) {
    if data.Status == "" && data.PlayerTwo == "" {
        return nil, errors.New("Provide valid data to update an Game")
    }

    game, err := g.repo.Update(id, data)

    if err != nil {
        return nil, err
    }

    return game, nil
}

func (g *gameService) RunGame(player string, id uint, text string) (*Game) {
    game, err := g.repo.Find(id) 

    if err != nil {
        panic(err)
    }

    if text == game.Text {
       game.Status = Finished 
       game.Winner = player

       updateGameData := UpdateGameData{
           Status: Finished,
           Winner: player,
       }

       g.repo.Update(id, updateGameData)

       return game
    }

    return nil
}
