package gorm

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
	"time"
)

// IdpInstance 账户idp配置
type IdpInstance struct {
	Id          uint64    `json:"id" gorm:"column:id;type:bigint unsigned"`                   // id
	Uid         uint64    `json:"uid" gorm:"column:uid;type:bigint unsigned"`                 // 唯一id（对外）
	AliasId     string    `json:"alias_id" gorm:"column:alias_id;type:varchar(64)"`           // 业务自定义id（对外）
	Name        string    `json:"name" gorm:"sadasa;column:na_me;type:varchar(64)"`           // 名称
	Description *string   `json:"description" gorm:"column:description"`                      // 描述
	IdpId       uint64    `json:"idp_id" gorm:"column:idp_id;type:bigint unsigned"`           // idp主键
	AccountUid  uint64    `json:"account_uid" gorm:"column:account_uid;type:bigint unsigned"` // account主键
	Priority    int8      `json:"priority" gorm:"column:priority;type:tinyint"`               // 优先级
	Settings    string    `json:"settings" gorm:"column:settings;type:varchar(2048)"`         // 配置数据
	Status      int8      `json:"status" gorm:"column:status;type:tinyint"`                   // 状态
	IsDeleted   int64     `json:"is_deleted" gorm:"column:is_deleted;type:bigint"`            // 是否删除
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;type:datetime"`          // 创建时间
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;type:datetime"`          // 更新时间
	LogoUrl     *string   `json:"logo_url" gorm:"column:logo_url;type:varchar(255)"`          // Logo
	UniqueName  *string   `json:"unique_name" gorm:"column:unique_name;type:varchar(12)"`     // 通用标识（可读）
}

func TestTag(tt *testing.T) {
	b := ""
	instance := &IdpInstance{
		Name:        "jinrr",
		Settings:    "a",
		Description: &b,
	}

	v := reflect.ValueOf(*instance)
	t := v.Type()

	pattern := ".*column:(\\w*).*"
	columnRegexp := regexp.MustCompile(pattern)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !isZero(field) {
			tag := t.Field(i).Tag.Get("gorm")
			result := columnRegexp.FindStringSubmatch(tag)

			fmt.Println(result[1], t.Field(i).Name, field.Interface())
		}
	}
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Struct:
		if v.Type().String() == "time.Time" {
			zeroTime := time.Time{}
			return v.Interface() == zeroTime
		}
	default:
		return false
	}
	return false
}
