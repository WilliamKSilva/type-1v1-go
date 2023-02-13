package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
    dsn := "type1v1user:type1v1@tcp(127.0.0.1:3306)/type1v1?charset=utf8mb4&parseTime=True&loc=Local" 
    mysql := mysql.Open(dsn)
    gorm, err := gorm.Open(mysql, &gorm.Config{})

    if err != nil {
        panic(err)
    }

    fmt.Printf("MySQL connected")

    return gorm
}
