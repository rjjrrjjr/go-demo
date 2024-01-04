package gorm

import (
	"log"

	"git.woa.com/fire/firedb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	Id    uint64 `json:"id"  gorm:"column:id"`
	Name  string `json:"name" gorm:"column:name" crypto:"aes"`
	BizId string `json:"biz_id" gorm:"-"`
}

func InitDB() (db *gorm.DB) {
	dbDSN := "root:123456@tcp(10.91.81.167:3306)/jinrruan_parent?charset=utf8mb4&parseTime=True&loc=Local" +
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
	//aesKey := configuration.GetConfigInfo().FdbConf.AesKey
	//fConfig := &firedb.FConfig{
	//	EnableLog: configuration.GetConfigInfo().FdbConf.EnableLog,
	//	Wm:        mode.WriteMode(configuration.GetConfigInfo().FdbConf.WriteMode),
	//	Qm:        mode.QueryMode(configuration.GetConfigInfo().FdbConf.QueryMode),
	//}
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
