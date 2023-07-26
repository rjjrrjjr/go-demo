package trpc

import (
	"context"
	"fmt"
	_ "git.code.oa.com/phoenixs/trpc-log-meet/log"
	tgorm "git.code.oa.com/trpc-go/trpc-database/gorm"
	"git.code.oa.com/trpc-go/trpc-go"
	"jinrruan-demo/gorm"
	"testing"
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
	err := db.WithContext(context.Background()).Table("t_user").Limit(10).Find(&userList).Error
	fmt.Println(userList, err)
}