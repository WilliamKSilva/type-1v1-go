package mysql

import (
	"github.com/WilliamKSilva/type-1v1/pkg/api"
	"gorm.io/gorm"
)

type gameRepositoryDB struct {
    connection *gorm.DB
}

func NewGameRepository (db *gorm.DB) *gameRepositoryDB {
    return &gameRepositoryDB{db} 
} 

func (db *gameRepositoryDB) Create(game *api.Game) error {
    err := db.connection.Create(game).Error

    if err != nil {
        return err
    } 

    return nil
}
