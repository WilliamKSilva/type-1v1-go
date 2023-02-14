package mysql

import (
	"errors"

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

func (db *gameRepositoryDB) Update(id uint, updateGameData api.UpdateGameData) (*api.Game, error) {
    var updatedGame *api.Game
    err := db.connection.Model(updatedGame).Where("id = ?", id).Updates(api.Game{Status: updateGameData.Status, PlayerTwo: updateGameData.PlayerTwo}).Error

    if err != nil {
        return nil, errors.New(err.Error())
    }

    return updatedGame, nil 
}
