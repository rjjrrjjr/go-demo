package gorm

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"testing"
)

func TestInitDB(t *testing.T) {
	db := InitDB()
	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()
	var UserList []User
	err := db.WithContext(context.Background()).Table("t_user").Limit(10).Find(&UserList).Error
	fmt.Println(UserList, err)
}

func TestCreate(t *testing.T) {
	db := InitDB()
	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()
	name := strings.ReplaceAll(uuid.New().String(), "-", "")
	fmt.Println(name)
	user := User{
		Name: name,
	}
	err := db.WithContext(context.Background()).Table("t_user").Create(&user).Error
	fmt.Println(err)
	TestInitDB(t)
}

func TestCreateWithOmitParam(t *testing.T) {
	db := InitDB()
	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()
	name := strings.ReplaceAll(uuid.New().String(), "-", "")
	user := User{
		Id: 6,
		Name: name,
	}
	err := db.WithContext(context.Background()).Table("t_user").Omit("id").Create(&user).Error
	fmt.Println(err)
	TestInitDB(t)
}

func TestCreateBatchWithOmitParam(t *testing.T) {
	db := InitDB()
	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()
	userList := make([]User, 0)
	for i := 6; i < 9; i++ {
		name := strings.ReplaceAll(uuid.New().String(), "-", "")
		userList = append(userList, User{
			Id: uint64(i),
			Name: name,
		})
	}
	err := db.WithContext(context.Background()).Table("t_user").Omit("id").Create(&userList).Error
	fmt.Println(err)
	TestInitDB(t)
}
