package api

import (
	"errors"
	"testing"
)


type mockGameRepository struct {}

func (m *mockGameRepository) Create (game *Game) error {
    game.ID = 1 
    game.Text = "bla bla bla bla"
    return nil
}

func TestShouldThrowIfPlayerOneNameIsMissing(t *testing.T) {
    repo := &mockGameRepository{}
    g := gameService{repo}

    gameData := NewGameData{
        Timer: 120,
    }

    want := errors.New("Player one name is required")
    _, err := g.NewGame(gameData)

    if err.Error() != want.Error() {
        t.Errorf("Expected: %s, got %s", want, err)
    } 
}

func TestShouldThrowIfTimerIsMissing(t *testing.T) {
    repo := &mockGameRepository{}
    g := gameService{repo}

    gameData := NewGameData{
        PlayerOne: "test",
    }

    want := errors.New("Timer is required")
    _, err := g.NewGame(gameData)

    if err.Error() != want.Error() {
        t.Errorf("Expected: %s, got %s", want, err)
    } 
}

func TestShouldReturnAGameOnSuccess(t *testing.T) {
    repo := &mockGameRepository{}
    g := gameService{repo}

    gameData := NewGameData{
        PlayerOne: "test",
        Timer: 160,
    }

    game, _ := g.NewGame(gameData)
    want := &Game{
        ID: 1,
        PlayerOne: "test",
        Timer: 160,
        Status: Waiting,
        Text: "bla bla bla bla",
    }

    if game.ID != want.ID {
        t.Errorf("Expected: %d, got %d", want.ID, game.ID)
    } 
}

