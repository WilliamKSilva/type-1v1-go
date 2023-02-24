package main

import (
	"log"
	"net/http"

	"github.com/WilliamKSilva/type-1v1/pkg"
	"github.com/WilliamKSilva/type-1v1/pkg/api"
	"github.com/WilliamKSilva/type-1v1/pkg/api/mysql"
)

func main() {
	db := mysql.Connect()
    cacheMap := make(map[string]*api.GameState)

    gameHandler := pkg.MakeGameHandler(db, cacheMap) 

	http.Handle("/games", gameHandler)
    http.HandleFunc("/games/run", gameHandler.RunGameFunc)
    http.HandleFunc("/games/cache", gameHandler.StoreGameCache)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
