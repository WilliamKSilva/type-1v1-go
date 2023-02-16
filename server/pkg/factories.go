package pkg 

import (
	"github.com/WilliamKSilva/type-1v1/pkg/api"
	"github.com/WilliamKSilva/type-1v1/pkg/api/mysql"
	"github.com/WilliamKSilva/type-1v1/pkg/web"
	"github.com/WilliamKSilva/type-1v1/pkg/ex"
	"gorm.io/gorm"
)

func MakeGameHandler (db *gorm.DB) *web.GameHandler {
    gameRepository := mysql.NewGameRepository(db)
    textService := ex.NewTextService() 
    gameService := api.NewGameService(gameRepository, textService)
    gameHandler := web.NewGameHandler(gameService)

    return gameHandler
}
