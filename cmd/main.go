package main

import (
	"log"
	"net/http"

	"github.com/WilliamKSilva/type-1v1/pkg/api"
	"github.com/WilliamKSilva/type-1v1/pkg/api/mysql"
	"github.com/WilliamKSilva/type-1v1/pkg/web"
)

func main() {
    db := mysql.Connect() 

    gameRepository := mysql.NewGameRepository(db)
    gameService := api.NewGameService(gameRepository)
    gameHandler := web.NewGameHandler(gameService) 

    http.HandleFunc("/bar", gameHandler.NewGameFunc)

    log.Fatal(http.ListenAndServe(":3000", nil))
}
