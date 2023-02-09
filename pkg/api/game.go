package api

type Game struct {
    ID uint `json:"id"` 
    PlayerOne string `json:"player_one"`
    PlayerTwo string `json:"player_two"`
    Text string `json:"text"`
    Timer uint `json:"timer"`
}

type newGameData struct {
    Timer uint `json:"timer"`
}
