package main

import (
	"log"
	"net/http"

	"github.com/WilliamKSilva/type-1v1/pkg/api/mysql"
	"github.com/WilliamKSilva/type-1v1/pkg"
)

func main() {
	db := mysql.Connect()

    gameHandler := pkg.MakeGameHandler(db) 

	http.Handle("/games", gameHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
