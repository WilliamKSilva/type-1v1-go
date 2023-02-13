package mysql

import (
	"fmt"

	"github.com/WilliamKSilva/type-1v1/pkg/api"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
    dsn := "root:type1v1@tcp(127.0.0.1:3306)/type1v1?charset=utf8mb4&parseTime=True&loc=Local" 
    mysql := mysql.Open(dsn)
    gorm, err := gorm.Open(mysql, &gorm.Config{})

    gorm.AutoMigrate(&api.Game{})

    if err != nil {
        panic(err)
    }

    fmt.Printf("MySQL connected")

    return gorm
}
