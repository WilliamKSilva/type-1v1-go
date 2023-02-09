package mysql

import (
	"github.com/WilliamKSilva/type-1v1/pkg/api"
	"gorm.io/gorm"
)

type gameRepository struct {
    connection *gorm.DB
}

func (db *gameRepository) Create(game *api.Game) error {
    err := db.connection.Create(game).Error

    if err != nil {
        return err
    } 

    return nil
}
