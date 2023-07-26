package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type User struct {
	Id uint64 `json:"id"  gorm:"column:id"`
	Name string	`json:"name" gorm:"column:name"`
}

func InitDB() (db *gorm.DB) {
	dbDSN := "root:123456@tcp(10.91.81.133:3306)/jinrruan_parent?charset=utf8mb4&parseTime=True&loc=Local" +
		"&readTimeout=3s&timeout=3s&writeTimeout=3s"
	db, err := gorm.Open(mysql.Open(dbDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("db init fail", err)
	}
	return db
}
