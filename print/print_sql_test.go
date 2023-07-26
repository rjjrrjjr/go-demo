package print

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"testing"
)

func TestPrintSql(t *testing.T) {
	signalSql := "SELECT * FROM user_ticket_record_%d where event_id = 'E168410505126674022404154' union all "

	for i := 0; i < 32; i++ {
		fmt.Println(fmt.Sprintf(signalSql, i))
	}
}

func TestPrintSqlIntoFile(t *testing.T) {
	file, err := os.OpenFile("a.sql", os.O_CREATE | os.O_RDWR | os.O_TRUNC, fs.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	signalSql := "CREATE TABLE `enroll_white_log_%d` (\n      `id` bigint(" +
		"20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',\n      `meeting_id` varchar(32) NOT NULL COMMENT '会议id',\n      `creator_appid` varchar(32) NOT NULL COMMENT '操作用户appid',\n      `creator_uid` varchar(32) NOT NULL COMMENT '操作用户uid',\n      `is_success` tinyint(3) unsigned DEFAULT '0' COMMENT '是否插入成功：0 否，1 是',\n      `modify_type` tinyint(3) unsigned DEFAULT '1' COMMENT '插入白名单类型:1 插入 2 更新',\n      `phone` varchar(32) NOT NULL DEFAULT '' COMMENT '对应白名单的电话号码',\n      `to_uid` varchar(32) NOT NULL DEFAULT '' COMMENT '添加指定的uid，非必填',\n      `err_status` tinyint(3) unsigned DEFAULT '0' COMMENT '错误类型:0 未发生错误，1 系统错误 2 手机号码不合法错误 3 手机填写为空 4 手机号码重复',\n      `err_info` varchar(2048) NOT NULL DEFAULT '' COMMENT '错误信息',\n      `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,\n      PRIMARY KEY (`id`),\n      KEY `idx_meeting_appid_uid` (`meeting_id`,`creator_uid`,`creator_appid`) USING BTREE,\n      KEY `idx_meeting_is_success` (`meeting_id`,`is_success`) USING BTREE,\n      KEY `idx_phone` (`phone`) USING BTREE\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='报名白名单导入日志表';"

	for i := 0; i < 8; i++ {

		_, _ = file.WriteString(fmt.Sprintln(fmt.Sprintf(signalSql, i)))
	}
}
