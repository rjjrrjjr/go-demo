package gorm

import (
	"fmt"
	"git.code.oa.com/trpc-go/trpc-go"
	"github.com/google/uuid"
	"strings"
	"testing"
)

// 索引测试

func TestUniqueError(t *testing.T) {

	db := InitDB()

	user := User{
		Id: 1,
		Name: strings.ReplaceAll(uuid.New().String(), "-", ""),
	}

	err := db.WithContext(trpc.BackgroundContext()).Table("t_user").Create(&user).Error

	fmt.Println("===============")
	fmt.Println(err)
	a := fmt.Sprintf("%+v", err)
	fmt.Println("A:", a)
	fmt.Println(strings.HasPrefix(a, "Error 1062"))
	fmt.Println("===============")
	fmt.Println(err.Error())
	fmt.Println(strings.HasPrefix(err.Error(), "Error 1062"))
	fmt.Println("===============")

}
