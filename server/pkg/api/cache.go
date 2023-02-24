package api

type GameState struct {
    GameID string `json:"gameId"`
    Player string `json:"player"`
    Text string `json:"text"`
}

type CacheService interface {
    Store (gameState *GameState) error
    Read (id uint) *GameState 
}

type cacheService struct {
    cacheMap map[string]*GameState
}

func (c *cacheService) Store (gameState *GameState) error {
    c.cacheMap[gameState.GameID] = gameState 

    return nil
}

func (c *cacheService) Read (id string) *GameState {
    gameState := c.cacheMap[id]

    return gameState
}
