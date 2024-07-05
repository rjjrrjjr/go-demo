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
	//err := db.WithContext(context.Background()).Table("t_user").Where("name = ?",
	// "jinrruan").Limit(100).Delete(&User{}).Error
	tx := db.WithContext(context.Background()).Exec("DELETE FROM ? WHERE id < ? limit ?", "t_user", 12, 100)
	fmt.Println(tx.RowsAffected, tx.Error)
}
