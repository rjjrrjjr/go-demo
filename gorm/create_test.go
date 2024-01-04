package gorm

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestInitDB(t *testing.T) {
	db := InitDB()
	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()
	var UserList []User
	err := db.WithContext(context.Background()).Table("t_user").Order("id desc").Limit(100).Find(&UserList).Error
	fmt.Println(UserList, err)
}

func TestSelect(t *testing.T) {
	// db := InitDB()
	// defer func() {
	// 	sqlDB, _ := db.DB()
	// 	_ = sqlDB.Close()
	// }()
	// var user *User
	// err := db.WithContext(context.Background()).Table("t_user").Order("id desc").Where("id = ?",
	// 	9999).Find(&user).Error
	// fmt.Println(user, err)
	var user User
	fmt.Println(&user == nil)
	fmt.Println(&user != nil && (&user).Id != 0)
}

func TestSave(t *testing.T) {
	db := InitDB()
	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()
	user := User{
		//Id: 17,
		Name:  "jinrruanm",
		BizId: strings.ReplaceAll(uuid.New().String(), "-", ""),
	}
	err := db.WithContext(context.Background()).Model(&User{}).Table("t_user").CreateInBatches([]*User{&user},
		10).Error
	fmt.Println(err)
	fmt.Println(user)

}

func TestFDBSave(t *testing.T) {
	fdb, err := GetFireDB()
	if err != nil {
		return
	}
	db := fdb.DB
	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()
	user := User{
		//Id: 17,
		Name:  "1",
		BizId: strings.ReplaceAll(uuid.New().String(), "-", ""),
	}
	userList := []*User{&user, &User{
		Name:  "2",
		BizId: strings.ReplaceAll(uuid.New().String(), "-", ""),
	}}
	err = db.WithContext(context.Background()).Table("t_user").CreateInBatches(userList, 10).Error
	fmt.Println(err)
	bytes, _ := json.Marshal(userList)
	fmt.Printf("%s\n", bytes)
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
		Id:   6,
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
		//name := strings.ReplaceAll(uuid.New().String(), "-", "")
		name := "jinrruan"
		userList = append(userList, User{
			Id:   uint64(i),
			Name: name,
		})
	}
	err := db.WithContext(context.Background()).Table("t_user").Omit("id").Create(&userList).Error
	fmt.Println(err)
	TestInitDB(t)
}
