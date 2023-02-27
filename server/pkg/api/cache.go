package api

import (
	"sync"
)

type GameState struct {
	GameID string `json:"gameId"`
	Player string `json:"player"`
	Text   string `json:"text"`
}

type States struct {
	gameStates *GameState
}

type CacheService interface {
	Store(gameState *GameState)
	Read(id string, player string) *GameState
	ReadAll(id string) map[string]*GameState 
}

type cacheService struct {
	mu sync.Mutex
	v  map[string]map[string]*GameState
}

func (c *cacheService) Store(game *GameState) {
	c.mu.Lock()

    if c.v[game.GameID] == nil {
        c.v[game.GameID] = make(map[string]*GameState)
    }

    c.v[game.GameID][game.Player] = game

	c.mu.Unlock()
}

func (c *cacheService) Read(id string, player string) *GameState {
	c.mu.Lock()

	defer c.mu.Unlock()

	states := c.v[id]

    for _, v := range states {
       if v.Player == player {
           return v
       }
    }

    return nil
}

func (c *cacheService) ReadAll(id string) map[string]*GameState {
	states := c.v[id]
	return states
}

func NewCacheService(cacheMap map[string]map[string]*GameState) *cacheService {
	return &cacheService{v: cacheMap}
}
