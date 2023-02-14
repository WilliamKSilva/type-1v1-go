package api

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockGameRepository struct {
    mock.Mock
}
type mockTextService struct {
    mock.Mock
}

func (m *mockGameRepository) Create (game *Game) error {
   args := m.Called(game)

   return args.Error(1) 
}

func (m *mockTextService) GetRandomText (trigger string) (string, error) {
   args := m.Called(trigger)

   return args.String(0), args.Error(1)
}

func TestShouldThrowIfPlayerOneNameIsMissing(t *testing.T) {
    repo := new(mockGameRepository)
    textService := new(mockTextService)

    repo.On("Create", mock.Anything).Return(errors.New("Player one name is required"))
    repo.On("GetRandomText", mock.Anything).Return("bla bla bla bla", nil)

    g := gameService{repo, textService}

    gameData := NewGameData{}

    want := errors.New("Player one name is required")
    _, err := g.NewGame(gameData)

    if err.Error() != want.Error() {
        t.Errorf("Expected: %s, got %s", want, err)
    } 
}

func TestShouldThrowIfTextServiceThrows(t *testing.T) {
    repo := &mockGameRepository{}
    g := gameService{repo}

    gameData := NewGameData{
        PlayerOne: "test",
    }

    game, _ := g.NewGame(gameData)
}

func TestShouldReturnAGameOnSuccess(t *testing.T) {
    repo := &mockGameRepository{}
    g := gameService{repo}

    gameData := NewGameData{
        PlayerOne: "test",
    }

    game, _ := g.NewGame(gameData)
    want := &Game{
        ID: 1,
        PlayerOne: "test",
        Status: Waiting,
        Text: "bla bla bla bla",
    }

    if game.ID != want.ID {
        t.Errorf("Expected: %d, got %d", want.ID, game.ID)
    } 
}

