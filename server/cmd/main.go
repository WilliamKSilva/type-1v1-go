package main

import (
	"log"
	"net/http"

	"github.com/WilliamKSilva/type-1v1/pkg"
	"github.com/WilliamKSilva/type-1v1/pkg/api/mysql"
)

func main() {
	db := mysql.Connect()

    gameHandler := pkg.MakeGameHandler(db) 

	http.Handle("/games", gameHandler)
    http.HandleFunc("/games/run", gameHandler.RunGameFunc)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
