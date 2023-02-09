package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connect() *gorm.DB {
    dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local" 
    mysql := mysql.Open(dsn)
    gorm, err := gorm.Open(mysql, &gorm.Config{})

    if err != nil {
        panic(err)
    }

    return gorm
}
