package jsons

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/tidwall/gjson"
)

func TestInden(t *testing.T) {
	str := "{\"extension\":\"{\\\"爱好\\\":\\\"旅游\\\",\\\"年龄\\\":\\\"24\\\"}\",\"unionid\":\"z21HjQliSzpw0YWxxxxx\",\"boss\":\"true\",\"role_list\":{\"group_name\":\"职务\",\"name\":\"总监\",\"id\":\"100\"},\"exclusive_account\":false,\"manager_userid\":\"manager240\",\"admin\":\"true\",\"remark\":\"备注备注\",\"title\":\"技术总监\",\"hired_date\":\"1597573616828\",\"userid\":\"zhangsan\",\"work_place\":\"未来park\",\"dept_order_list\":{\"dept_id\":\"2\",\"order\":\"1\"},\"real_authed\":\"true\",\"dept_id_list\":\"[2,3,4]\",\"job_number\":\"4\",\"email\":\"test@xxx.com\",\"leader_in_dept\":{\"leader\":\"true\",\"dept_id\":\"2\"},\"mobile\":\"18513027676\",\"active\":\"true\",\"org_email\":\"test@xxx.com\",\"telephone\":\"010-86123456-2345\",\"avatar\":\"xxx\",\"hide_mobile\":\"false\",\"senior\":\"true\",\"name\":\"张三\",\"union_emp_ext\":{\"union_emp_map_list\":{\"userid\":\"5000\",\"corp_id\":\"dingxxx\"},\"userid\":\"500\",\"corp_id\":\"dingxxx\"},\"state_code\":\"86\"}"

	var data interface{}
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		fmt.Println("解析JSON失败:", err)
		return
	}

	formattedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("格式化JSON失败:", err)
		return
	}

	fmt.Println(string(formattedJSON))
}

func TestD(t *testing.T) {
	extra := "{\"expiresInSeconds\":\"300\",\"clockSkewSeconds\":300,\"secret_key_method\":\"upload\",\"max_clock_skew\":null}"
	var settings IdpJwt
	err := json.Unmarshal([]byte(extra), &settings)
	fmt.Println(err)
}

func TestC(t *testing.T) {
	var oauthSetting IdpOAuth2
	err := json.Unmarshal([]byte("{}"), &oauthSetting)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !oauthSetting.IDaaSCompliant {
		oauthSetting.IDaaSCompliant = false
	}
	extraStr, _ := json.Marshal(oauthSetting)
	fmt.Println(string(extraStr))
}

func TestA(t *testing.T) {
	setting := "{\"extra\":\"{\\\"agent_id\\\":\\\"AgentId124132\\\",\\\"app_id\\\":\\\"AppKey1232131231\\\"}\",\"defaultDepartment\":\"asdasda\"}"
	deptResult := gjson.GetBytes([]byte(setting), "defaultDepartment").String()
	fmt.Println("====", deptResult, "====")
}

func TestJson(t *testing.T) {
	var user User
	err := json.Unmarshal([]byte("{\"name\":\"jinrruan\", \"age\":18}"), &user)
	fmt.Println(user, err)

	var user2 User
	err = json.Unmarshal([]byte("{}"), &user2)
	fmt.Println(user2, err)

	var user3 User
	err = json.Unmarshal(nil, &user3)

	fmt.Println(user3, err)
	fmt.Println(err.Error())

}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestB(t *testing.T) {
	setting := "{\"extra\":\"{\\\"agent_id\\\":\\\"dingding_agentid\\\",\\\"app_id\\\":\\\"dingding_appkey\\\"}\",\n        \"credentials\":\"{\\\"app_secret\\\":\\\"dingding_appsecret\\\"}\",\n        \"callbackUrl\":\"\"}"

	var appId, appSecret, agentId string
	extra := gjson.GetBytes([]byte(setting), "extra").String()
	credentials := gjson.GetBytes([]byte(setting), "credentials").String()
	if extra != "" {
		appId = gjson.Get(extra, "app_id").String()
		agentId = gjson.Get(extra, "agent_id").String()
	}
	if credentials != "" {
		appSecret = gjson.Get(credentials, "app_secret").String()
	}

	fmt.Println(appId, appSecret, agentId)
}

// IdpOAuth2 OAuth2 IdP config
type IdpOAuth2 struct {
	UserIDKey               string `json:"user_id_key"`     // 用户IdP的唯一标识
	IDaaSCompliant          bool   `json:"idaas_compliant"` // 兼容老的idaas逻辑, 对于新建的一律为false, 从idaas导入的为true(不校验最大的过期时间)
	AuthorizationEndpoint   string `json:"authorization_endpoint,omitempty"`
	TokenEndpoint           string `json:"token_endpoint,omitempty"`
	TokenEndpointAuthMethod string `json:"token_endpoint_auth_method,omitempty"`
	UserinfoEndpoint        string `json:"userinfo_endpoint,omitempty"`
	// access_token传入方式, 1:header-basic auth,2:query-query string,3:form-form body
	UserinfoEndpointAuthMethod string   `json:"userinfo_endpoint_auth_method,omitempty"`
	ClientID                   string   `json:"client_id"`
	Scopes                     []string `json:"scopes"`
}

type ActiveDirectorySettings struct {
	ServerUrl          string `json:"serverUrl,omitempty"`
	Host               string `json:"host,omitempty"`
	Port               string `json:"port,omitempty"`
	Protocol           string `json:"protocol,omitempty"`
	AdminDn            string `json:"adminDN"`
	BaseDn             string `json:"baseDN"`
	UnitSource         string `json:"unitSource"`
	UniqueIDAttribute  string `json:"uniqueIDAttribute,omitempty"`
	UserBindTemplate   string `json:"userBindTemplate"`
	UserFilterTemplate string `json:"userObjectFilter"`
	UsernameFormat     string `json:"usernameFormat"`
	AgentEnabled       bool   `json:"agentEnabled,omitempty"`
	RelayMode          int32  `json:"relayMode,omitempty"`
}

type IDaaSSigPubKey string

type IdpJwt struct {
	UserKey      string `json:"userKey"`      // 用户IdP的唯一标识
	IDTokenParam string `json:"idTokenParam"` // 认证请求中, id_token对应的参数名, 缺省为id_token

	// 签名算法和公钥
	SignAlgorithm string         `json:"sigAlg"`    // 签名算法, 默认为RS256, 将来可扩展支持ES256(注: 不支持对称算法)
	SignKeyPub    IDaaSSigPubKey `json:"sigKeyPub"` // 对应的签名公钥, 和SignAlgorithm配套使用

	// 校验id_token时的选项
	IdaasCompliant     bool   `json:"idaasCompliant"`   // 兼容老的idaas逻辑, 对于新建的一律为false, 从idaas导入的为true(不校验最大的过期时间)
	Issuer             string `json:"iss"`              // id_token中的iss claim, 当Issuer不为空时, 校验id_token时需要校验issuer匹配
	ExpiresInSeconds   int64  `json:"expiresInSeconds"` // id_token允许最大的过期时间, 单位秒, 缺省5分钟
	ClockSkewInSeconds int64  `json:"clockSkewSeconds"` // 系统时钟最大的偏差, 单位秒, 缺省为1分钟
}
