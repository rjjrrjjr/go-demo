package string

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode/utf8"
)

// 去空格和换行
func TestTrim(t *testing.T) {
	str := "{\n    \"extension\": \"{\\\"爱好\\\":\\\"旅游\\\",\\\"年龄\\\":\\\"24\\\"}\",\n    \"unionid\": \"z21HjQliSzpw0YWxxxxx\",\n    \"boss\": \"true\",\n    \"role_list\": {\n        \"group_name\": \"职务\",\n        \"name\": \"总监\",\n        \"id\": \"100\"\n    },\n    \"exclusive_account\": false,\n    \"manager_userid\": \"manager240\",\n    \"admin\": \"true\",\n    \"remark\": \"备注备注\",\n    \"title\": \"技术总监\",\n    \"hired_date\": \"1597573616828\",\n    \"userid\": \"zhangsan\",\n    \"work_place\": \"未来park\",\n    \"dept_order_list\": {\n        \"dept_id\": \"2\",\n        \"order\": \"1\"\n    },\n    \"real_authed\": \"true\",\n    \"dept_id_list\": \"[2,3,4]\",\n    \"job_number\": \"4\",\n    \"email\": \"test@xxx.com\",\n    \"leader_in_dept\": {\n        \"leader\": \"true\",\n        \"dept_id\": \"2\"\n    },\n    \"mobile\": \"18513027676\",\n    \"active\": \"true\",\n    \"org_email\": \"test@xxx.com\",\n    \"telephone\": \"010-86123456-2345\",\n    \"avatar\": \"xxx\",\n    \"hide_mobile\": \"false\",\n    \"senior\": \"true\",\n    \"name\": \"张三\",\n    \"union_emp_ext\": {\n        \"union_emp_map_list\": {\n            \"userid\": \"5000\",\n            \"corp_id\": \"dingxxx\"\n        },\n        \"userid\": \"500\",\n        \"corp_id\": \"dingxxx\"\n    },\n    \"state_code\": \"86\"\n}"
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	//str = strings.Replace(str, "\"", "\\\"", -1)
	fmt.Println(str)

}

func TestDesensitization(t *testing.T) {
	fmt.Println(phoneDesensitization("1234567"))
	fmt.Println(phoneDesensitization("123456"))
	fmt.Println(phoneDesensitization("12345"))
	fmt.Println(phoneDesensitization("1234"))
	fmt.Println(phoneDesensitization("123"))
	fmt.Println(phoneDesensitization("12"))
	fmt.Println(phoneDesensitization("1"))
}

func TestStartWith(t *testing.T) {
	fmt.Println(reBuildArea("12345"))
	fmt.Println(reBuildArea("+12345"))
	fmt.Println(reBuildArea("+"))
}

func chunkArray(sourceArray []int, chunkSize int) [][]int {
	endIndex := chunkSize
	result := make([][]int, 0)
	for endIndex <= len(sourceArray) {
		result = append(result, sourceArray[endIndex-chunkSize:endIndex])
		endIndex = endIndex + chunkSize
	}
	if endIndex > len(sourceArray) {
		result = append(result, sourceArray[endIndex-chunkSize:])
	}
	return result
}

func TestSubString(t *testing.T) {
	numList := make([]int, 0)
	for i := 0; i < 1002; i++ {
		numList = append(numList, i)
	}

	fmt.Println(numList[:10])
	fmt.Println(numList[:10])

	// 500 一组打印
	result := chunkArray(numList, 200)
	for _, item := range result {
		fmt.Println(item)
	}
}

func TestI2A(t *testing.T) {
	fmt.Println("---------")
	a, b := strconv.ParseUint("32", 10, 3)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("---------")
}

func TestStr(t *testing.T) {
	a := "我是中国人"
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(utf8.RuneCountInString(a))
	fmt.Println(a[:1])
	fmt.Println(a[:2])
	fmt.Println(a[:3])
	fmt.Println(string([]rune(a)[:4]))
}

func reBuildArea(area string) string {
	if len(area) > 1 && strings.HasPrefix(area, "+") {
		return area[1:]
	}
	return area
}

func phoneDesensitization(phone string) string {
	if len(phone) > 6 {
		return phone[:3] + "***" + phone[len(phone)-3:]
	} else if len(phone) > 4 {
		return phone[:2] + "***" + phone[len(phone)-2:]
	} else if len(phone) > 2 {
		return phone[:1] + "***" + phone[len(phone)-1:]
	} else {
		return "***"
	}
}

func TestChunkTime(t *testing.T) {
	paramList := make([][]int, 0)

	for i := 0; i < 100; i++ {
		curParam := make([]int, 0)
		for j := i; j < i+100; j++ {
			curParam = append(curParam, j)
		}
		paramList = append(paramList, curParam)
	}
	fmt.Println(time.Now())
	for index, curParam := range paramList {
		if index%9 == 0 {
			fmt.Println(chunkInt(curParam, 9))
		}
		chunkInt(curParam, 9)
	}
	fmt.Println(time.Now())
	for index, curParam := range paramList {
		if index%9 == 0 {
			fmt.Println(chunkInt2(curParam, 9))
		}
		chunkInt2(curParam, 9)
	}
	fmt.Println(time.Now())
}

func TestChunkStr(t *testing.T) {
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	fmt.Println(s.Alloc, bToMb(s.Alloc))
	param := make([]string, 0)
	for i := 0; i < 32000; i++ {
		param = append(param, strconv.FormatInt(int64(i), 10))
	}
	//fmt.Println(param)
	runtime.ReadMemStats(&s)
	fmt.Println(s.Alloc, bToMb(s.Alloc))

	//fmt.Println(chunkApprovalList(param, 100))
	chunkApprovalList(param, 500)
	runtime.ReadMemStats(&s)
	fmt.Println(s.Alloc, bToMb(s.Alloc))
	fmt.Println("----------")
}

func chunkApprovalList(approvalList []string, chunkSize int) [][]string {
	result := make([][]string, 0)

	tempList := make([]string, 0)

	for index, approval := range approvalList {
		if index%chunkSize == 0 && len(tempList) > 0 {
			result = append(result, tempList)
			tempList = make([]string, 0)
		}
		tempList = append(tempList, approval)
	}
	if len(tempList) > 0 {
		result = append(result, tempList)
	}
	return result
}

func chunkCurStruct(approvalList []*CurStruct, chunkSize int) [][]*CurStruct {
	result := make([][]*CurStruct, 0)

	tempList := make([]*CurStruct, 0)

	for index, approval := range approvalList {
		if index%chunkSize == 0 && len(tempList) > 0 {
			result = append(result, tempList)
			tempList = make([]*CurStruct, 0)
		}
		tempList = append(tempList, approval)
	}
	if len(tempList) > 0 {
		result = append(result, tempList)
	}
	return result
}

func chunkInt(approvalList []int, chunkSize int) [][]int {
	result := make([][]int, 0)

	tempList := make([]int, 0)

	for index, approval := range approvalList {
		if index%chunkSize == 0 && len(tempList) > 0 {
			result = append(result, tempList)
			tempList = make([]int, 0)
		}
		tempList = append(tempList, approval)
	}
	if len(tempList) > 0 {
		result = append(result, tempList)
	}
	return result
}

func chunkInt2(approvalList []int, chunkSize int) [][]int {
	endIndex := chunkSize
	result := make([][]int, 0)
	if len(approvalList) == 0 {
		return result
	}
	for endIndex <= len(approvalList) {
		result = append(result, approvalList[endIndex-chunkSize:endIndex])
		endIndex = endIndex + chunkSize
	}
	if endIndex > len(approvalList) {
		result = append(result, approvalList[endIndex-chunkSize:])
	}
	return result
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func TestChunkStruct(t *testing.T) {
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	fmt.Println(s.Alloc, bToMb(s.Alloc))
	param := make([]*CurStruct, 0)
	for i := 0; i < 32000; i++ {
		param = append(param, &CurStruct{
			Id:        uint64(i),
			UserName:  strconv.FormatInt(int64(i), 10),
			MeetingId: "asdsa",
		})
	}
	runtime.ReadMemStats(&s)
	fmt.Println(s.Alloc, bToMb(s.Alloc))

	chunkCurStruct(param, 500)
	runtime.ReadMemStats(&s)
	fmt.Println(s.Alloc, bToMb(s.Alloc))
	fmt.Println("----------")
}

type CurStruct struct {
	Id                uint64    `json:"id" gorm:"column:id"`                               // 主键ID
	MeetingId         string    `json:"meeting_id" gorm:"column:meeting_id"`               // 会议ID
	Appid             string    `json:"appid" gorm:"column:appid"`                         // 用户appid
	Uid               string    `json:"uid" gorm:"column:uid"`                             // 用户uid
	UserName          string    `json:"user_name" gorm:"column:user_name"`                 // 用户姓名
	Status            uint32    `json:"status" gorm:"column:status"`                       // 批准状态：1 待审核，2 已拒绝，3 已批准
	RealIp            string    `json:"real_ip" gorm:"column:real_ip"`                     // 用户真实报名Ip地址
	IsDelete          uint32    `json:"is_delete" gorm:"column:is_delete"`                 // 是否删除：0 否，1 是
	Area              string    `json:"area" gorm:"column:area"`                           // 手机号地区码
	Phone             string    `json:"phone" gorm:"column:phone" crypto:"aes"`            // 手机号
	Source            uint32    `json:"source" gorm:"column:source"`                       // 报名方式：1-自主报名 2-批量导入
	EnrollCode        string    `json:"enroll_code" gorm:"column:enroll_code"`             // 报名码
	OrderId           string    `json:"order_id" gorm:"column:order_id"`                   // 订单id
	OrderPrice        uint32    `json:"order_price" gorm:"column:order_price"`             // 订单金额
	OrderStatus       uint32    `json:"order_status" gorm:"column:order_status"`           // 订单状态：0未创建，1下单，4已支付，7订单超时关闭，8退款中，9退款完成，10退款失败
	UserNamePlaintext string    `json:"user_name_plaintext" gorm:"user_name_plaintext"`    // 用户昵称：明文存储
	OrderCreateTime   uint32    `json:"order_create_time" gorm:"column:order_create_time"` // 订单创建时间，单位秒
	OrderUpdateTime   uint32    `json:"order_update_time" gorm:"column:order_update_time"` // 订单更新时间，单位秒
	OrderPayTime      uint32    `json:"order_pay_time" gorm:"column:order_pay_time"`       // 订单支付时间，单位秒
	ChannelName       string    `json:"channel_name" gorm:"column:channel_name"`           // 渠道名称
	ChannelUrl        string    `json:"channel_url" gorm:"column:channel_url"`             // 渠道名称映射的URL子串
	CreateTime        time.Time `json:"create_time" gorm:"default:null"`                   // 创建时间
	District          uint32    `json:"district" gorm:"column:district"`                   // 区域，国内-0，海外-1
	AuditorUid        string    `json:"auditor_uid" gorm:"column:auditor_uid"`             // 审核者Uid
	AuditorName       string    `json:"auditor_name" gorm:"column:auditor_name"`           // 审核者名字
	LoginType         string    `json:"login_type" gorm:"column:login_type"`               // 登陆类型
	PayFlowId         string    `json:"pay_flow_id" gorm:"column:pay_flow_id"`             // 活动的支付流水id
}
