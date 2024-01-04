package gorm

import (
	"context"
	"fmt"
	"testing"
)

func TestDelete(t *testing.T) {
	db := InitDB()
	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()
	err := db.WithContext(context.Background()).Table("t_user").Where("name = ?",
		"jinrruan").Limit(1).Delete(&User{}).Limit(1).Error
	//err := db.WithContext(context.Background()).Exec("DELETE FROM `t_user` WHERE name = 'jinrruan' limit 1").Error
	fmt.Println(err)
}
