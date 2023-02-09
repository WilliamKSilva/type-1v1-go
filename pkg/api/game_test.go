package api

import (
	"errors"
	"testing"
)


type mockGameRepository struct {}

func (m *mockGameRepository) Create (game *Game) error {
    return nil
}

func TestShouldThrowIfPlayerOneNameIsMissing(t *testing.T) {
    repo := &mockGameRepository{}
    g := gameService{repo}

    gameData := newGameData{
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

    gameData := newGameData{
        PlayerOne: "test",
    }

    want := errors.New("Timer is required")
    _, err := g.NewGame(gameData)

    if err.Error() != want.Error() {
        t.Errorf("Expected: %s, got %s", want, err)
    } 
}
