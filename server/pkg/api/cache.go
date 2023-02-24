package api

import (
	"sync"
)

type GameState struct {
    GameID string `json:"gameId"`
    Player string `json:"player"`
    Text string `json:"text"`
}

type CacheService interface {
    Store (gameState *GameState)
    Read (id string) *GameState
    ReadAll () map[string]*GameState
}

type cacheService struct {
    mu sync.Mutex
    v map[string]*GameState
}

func (c *cacheService) Store (game *GameState) {
    c.mu.Lock()

    c.v[game.GameID] = game

    c.mu.Unlock()
}

func (c *cacheService) Read (id string) *GameState {
    c.mu.Lock()

    defer c.mu.Unlock()

    return c.v[id]
}

func (c *cacheService) ReadAll () map[string]*GameState {
    return c.v
}

func NewCacheService (cacheMap map[string]*GameState) *cacheService {
    return &cacheService{v: cacheMap} 
}
