package trpc

import (
	"context"
	"fmt"
	"testing"

	_ "git.code.oa.com/phoenixs/trpc-log-meet/log"
	tgorm "git.code.oa.com/trpc-go/trpc-database/gorm"
	"git.code.oa.com/trpc-go/trpc-go"
	"jinrruan-demo/common/gorm"
)

func TestGOrmTrpc(t *testing.T) {
	trpc.ServerConfigPath = "./trpc_go.yaml"
	_ = trpc.NewServer()

	db, _ := tgorm.NewClientProxy("trpc.mysql.wemeet.jinrruan")
	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()

	}()
	var userList []gorm.User
	err := db.WithContext(context.Background()).Table("t_user").Limit(1).Find(&userList).Error
	fmt.Println(userList, err)
	err = db.WithContext(context.Background()).Table("t_user").Limit(1).Find(&userList).Error
	fmt.Println(userList, err)
	err = db.WithContext(context.Background()).Table("t_user").Limit(1).Find(&userList).Error
	fmt.Println(userList, err)
	err = db.WithContext(context.Background()).Table("t_user").Limit(1).Find(&userList).Error
	fmt.Println(userList, err)
}

func TestGormTrpcDelete(t *testing.T) {
	trpc.ServerConfigPath = "./trpc_go.yaml"
	_ = trpc.NewServer()
	db, _ := tgorm.NewClientProxy("trpc.mysql.wemeet.jinrruan")
	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()
	tx := db.WithContext(context.Background()).Table("t_user").Where("id in (?)", []int64{34, 35, 36, 341, 342}).Limit(1).Delete(&gorm.User{})
	fmt.Println("-================")
	fmt.Println(tx.RowsAffected, tx.Error)
	fmt.Println("-================")
}
