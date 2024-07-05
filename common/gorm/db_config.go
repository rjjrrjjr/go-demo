package gorm

import (
	"log"
	"time"

	"git.woa.com/fire/firedb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	Id         uint64     `json:"id"  gorm:"column:id"`
	Age        *int8      `json:"age" gorm:"column:age"`
	Name       string     `json:"name" gorm:"column:name" crypto:"aes"`
	BizId      string     `json:"biz_id" gorm:"-"`
	CreateTime *time.Time `json:"create_time" gorm:"column:create_time"`
}

func InitDB() (db *gorm.DB) {
	dbDSN := "root:123456@tcp(10.91.81.64:3306)/jinrruan_parent?parseTime=true&charset=utf8mb4&loc=Local" +
		"&readTimeout=3s&timeout=3s&writeTimeout=3s"
	db, err := gorm.Open(mysql.Open(dbDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("db init fail", err)
	}
	return db
}

func GetFireDB() (fdb *firedb.FireDB, err error) {
	aesKey := "d1s3HvwAdxySvmi0"
	fConfig := &firedb.FConfig{
		EnableLog: true,
		Wm:        1,
		Qm:        2,
	}

	fdb, err = firedb.NewFireDBWithGormDB(InitDB())
	fdb.EnableAES(aesKey)
	err = fdb.ApplyConfig(fConfig)
	return
}
