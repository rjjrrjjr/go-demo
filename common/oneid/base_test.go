package oneid

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestOneid(t *testing.T) {
	var idpInstanceSettings IdpInstanceSettings
	err := json.Unmarshal([]byte("eyJtYXRjaE1hcHBpbmdSdWxlcyI6W3sibWF0Y2hNYXBwaW5ncyI6W3sic291cmNlIjoiTmFtZUlEIiwidGFyZ2V0IjoiYWxpYXNJZCJ9XX1dLCJjcmVhdGVNYXBwaW5ncyI6W3sidGFyZ2V0IjoibmFtZSJ9LHsidGFyZ2V0IjoidXNlcm5hbWUifSx7InRhcmdldCI6Im1vYmlsZSJ9LHsidGFyZ2V0IjoiZW1haWwifV0sImNhbGxiYWNrVXJsIjoiaHR0cHM6Ly9vYXV0aDIuZWlkLTYuYWNjb3VudC50ZW5jZW50Y3MuY29tL3YxL3Nzby9jYWxsYmFjay9zYW1sMl92MS8xMTM3MzA2OTc4ODA1MDg4MjU2LzExNjIzNTg1ODY2MzkxMjI0MzIiLCJyZWRpcmVjdFVybCI6Imh0dHBzOi8vb2F1dGgyLmVpZC02LmFjY291bnQudGVuY2VudGNzLmNvbS92MS9zc28vbG9naW4vc2FtbDJfdjEvMTEzNzMwNjk3ODgwNTA4ODI1Ni8xMTYyMzU4NTg2NjM5MTIyNDMyIiwiZXh0cmEiOiJ7XCJpbXBvcnRUeXBlXCI6XCJNQU5VQUxcIixcImlzc3VlclwiOlwiaHR0cDovL2EuYlwiLFwic3NvVXJsXCI6XCJodHRwOi8vYS5iXCIsXCJzaWdDZXJ0XCI6XCIw77+9XFx1MDAwM++/vVxcdTAwMDZcXHQq77+9SO+/ve+/vVxcclxcdTAwMDFcXHUwMDA3XFx1MDAwMu+/ve+/vVxcdTAwMDPvv70w77+9XFx1MDAwM++/vVxcdTAwMDJcXHUwMDAxXFx1MDAwMTFcXHUwMDAwMFxcdTAwMGJcXHUwMDA2XFx0Ku+/vUjvv73vv71cXHJcXHUwMDAxXFx1MDAwN1xcdTAwMDHvv73vv71cXHUwMDAzfjDvv71cXHUwMDAzejDvv71cXHUwMDAyYu+/vVxcdTAwMDNcXHUwMDAyXFx1MDAwMVxcdTAwMDJcXHUwMDAyXFx1MDAwNlxcdTAwMDHvv71cXHJWMGAwXFxyXFx1MDAwNlxcdCrvv71I77+977+9XFxyXFx1MDAwMVxcdTAwMDFcXHUwMDBiXFx1MDAwNVxcdTAwMDAwfjFcXHUwMDBiMFxcdFxcdTAwMDZcXHUwMDAzVVxcdTAwMDRcXHUwMDA2XFx1MDAxM1xcdTAwMDJVUzFcXHUwMDE2MFxcdTAwMTRcXHUwMDA2XFx1MDAwM1VcXHUwMDA0XFxuXFxmXFxyUGluZyBJZGVudGl0eTFcXHUwMDE2MFxcdTAwMTRcXHUwMDA2XFx1MDAwM1VcXHUwMDA0XFx1MDAwYlxcZlxcclBpbmcgSWRlbnRpdHkxPzA9XFx1MDAwNlxcdTAwMDNVXFx1MDAwNFxcdTAwMDNcXGY2UGluZ09uZSBTU08gQ2VydGlmaWNhdGUgZm9yIEFkbWluaXN0cmF0b3JzIGVudmlyb25tZW50MFxcdTAwMWVcXHUwMDE3XFxyMjQwMzA1MDYzODU2WlxcdTAwMTdcXHIyNTAzMDUwNjM4NTZaMH4xXFx1MDAwYjBcXHRcXHUwMDA2XFx1MDAwM1VcXHUwMDA0XFx1MDAwNlxcdTAwMTNcXHUwMDAyVVMxXFx1MDAxNjBcXHUwMDE0XFx1MDAwNlxcdTAwMDNVXFx1MDAwNFxcblxcZlxcclBpbmcgSWRlbnRpdHkxXFx1MDAxNjBcXHUwMDE0XFx1MDAwNlxcdTAwMDNVXFx1MDAwNFxcdTAwMGJcXGZcXHJQaW5nIElkZW50aXR5MT8wPVxcdTAwMDZcXHUwMDAzVVxcdTAwMDRcXHUwMDAzXFxmNlBpbmdPbmUgU1NPIENlcnRpZmljYXRlIGZvciBBZG1pbmlzdHJhdG9ycyBlbnZpcm9ubWVudDDvv71cXHUwMDAxXFxcIjBcXHJcXHUwMDA2XFx0Ku+/vUjvv73vv71cXHJcXHUwMDAxXFx1MDAwMVxcdTAwMDFcXHUwMDA1XFx1MDAwMFxcdTAwMDPvv71cXHUwMDAxXFx1MDAwZlxcdTAwMDAw77+9XFx1MDAwMVxcblxcdTAwMDLvv71cXHUwMDAxXFx1MDAwMVxcdTAwMDDvv73vv70wXFx1MDAxZO+/vTrvv71J77+9Vkzvv73vv73vv73vv71oWe+/vS1I77+9XFx1MDAxOe+/ve+/vXvvv73VuVxcdTAwMTLvv73vv71cXGJ8UFxcdTAwMDVydXvIhVxcdTAwMDDvv73vv71gQu+/vVLvv71cXHUwMDEzJndY77+9dWM877+977+9X1xcXCLvv73vv71b77+9alxcdTAwMWJq77+977+9SO+/vWt1bzjvv73vv73vv73vv703YO+/ve+/ve+/vVxcdTAwMWNcXHUwMDAz7JqX1Jnvv73vv73vv714XFx1MDAxMlxcdTAwMTfvv70+77+9ae+/ve+/ve+/vTvvv70zzp1cXHUwMDEz77+977+977+977+9f++/vVxcdTAwMTdcXHUwMDBmyIHvv73vv73vv70qL1xcZu+/vSTvv70vXFxy77+9Uibvv71cXHUwMDA077+9XFx1MDAxOe+/vVxcdTAwMWLvv70877+9XFx1MDAwNzV2XFxy77+977+9ejhcXHLvv73vv73vv73vv70zay7vv70677+977+9Ie+/ve+/vX1cXHUwMDE277+977+954aw77+977+9Qu+/vTbvv701W8ip77+977+91YtJ77+9XFx1MDAxNu+/vVxcdTAwMTV0bTorXu+/vV3vv71cXHRcXHUwMDE377+92LDvv71cXHUwMDBmdELvv73vv73vv71cXHUwMDBi77+9cFxcdO+/vVxcbixcXHUwMDAwWm3Elu+/ve+/ve+/vVxcdTAwMTnvv73vv71x77+9Me+/vT85MO+/ve+/ve+/vVxcdTAwMTZcXHUwMDEzdjPvv73vv73vv71dXFx1MDAwM03vv71cXHUwMDAyXFx1MDAwM1xcdTAwMDFcXHUwMDAwXFx1MDAwMTBcXHJcXHUwMDA2XFx0Ku+/vUjvv73vv71cXHJcXHUwMDAxXFx1MDAwMVxcdTAwMGJcXHUwMDA1XFx1MDAwMFxcdTAwMDPvv71cXHUwMDAxXFx1MDAwMVxcdTAwMDBcXG5cXGI377+977+9KWXvv71qXFxuzo9cXHUwMDBm77+9XFx1MDAxYdiKXFx1MDAxNu+/ve+/ve+/ve+/vU5mQe+/vTIm77+9Nu+/ve+/ve+/vV5cXHUwMDFm77+977+977+9PO+/vTbvv73vv73vv711XFx1MDAxOGlcXHUwMDExc++/vS9i77+977+977+977+9VDvvv71b77+977+977+977+977+977+977+977+9YlxcdTAwMThcXHUwMDAyMu+/ve+/vWXvv71a77+977+977+977+94Lq5J++/vSzvv73vv70577+977+9XFx1MDAxOO+/ve+/vVEz77+977+9QVfvv70p77+977+9XFx1MDAxYTBcXHUwMDFj77+977+977+977+9Y++/vXd877+9XFx1MDAxNO+/ve+/ve+/ve+/vWwv77+977+9XFxu77+9TO+/ve+/vXBR77+9XFx1MDAwZu+/vStd77+9TO+/ve+/vSNU77+9RDzvv703XFx1MDAxOSDvv73vv712YllcXHUwMDAx77+977+977+977+9XFx1MDAwN++/vVI1eSvvv73vv73vv70l77+9VO+/ve+/ve+/vWFS77+977+977+977+977+9U0vvv73vv71cXHI8XFx1MDAxYu+/vVxcdTAwMTdn77+977+9XFx1MDAxMu+/vUtGYVxcdTAwMDPvv71M77+9Pu+/ve+/ve+/vVxcXCLKh++/ve+/ve+/vW46N1xcYk5t77+9Qjjvv73vv73vv73vv73vv71cXHUwMDAyW++/ve+/ve+/ve+/ve+/ve+/vVp277+9XFx1MDAwMWJcXHUwMDFk77+977+9PO+/vVxcdTAwMDLvv71cXHUwMDFiXFxcIjFcXHUwMDAwXCIsXCJyZXFCaW5kaW5nXCI6XCJ1cm46b2FzaXM6bmFtZXM6dGM6U0FNTDoyLjA6YmluZGluZ3M6SFRUUC1QT1NUXCIsXCJyZXFTaWdcIjp0cnVlLFwicmVxU2lnQWxnXCI6XCJodHRwOi8vd3d3LnczLm9yZy8yMDAxLzA0L3htbGRzaWctbW9yZSNyc2Etc2hhMjU2XCIsXCJyZXNwU2lnQWxnXCI6XCJodHRwOi8vd3d3LnczLm9yZy8yMDAxLzA0L3htbGRzaWctbW9yZSNyc2Etc2hhMjU2XCIsXCJyZXNwU2lnVmVyXCI6XCJBTllcIixcIm5hbWVJREZvcm1hdFwiOlwidXJuOm9hc2lzOm5hbWVzOnRjOlNBTUw6Mi4wOm5hbWVpZC1mb3JtYXQ6cGVyc2lzdGVudFwifSIsInJlZGlyZWN0Q291bnRkb3duIjpudWxsLCJqaXRXaGVuTm9NYXRjaCI6ZmFsc2V9"),
		&idpInstanceSettings)
	fmt.Println("=================")
	fmt.Println(idpInstanceSettings, err)
	fmt.Println("=================")
}

// IdpInstanceSettings idp instance settings
type IdpInstanceSettings struct {
	// IDP匹配的映射关系
	MatchMappingRules []*AttrMappingRules `json:"matchMappingRules,omitempty"`
	// 无法找到匹配时是否进行即时创建(JIT)
	JitWhenNoMatch bool `json:"jitWhenNoMatch,omitempty"`
	// IDP即时创建(JIT)的映射关系
	CreateMappings []*AttrMapping `json:"createMappings,omitempty"`
	// 自动跳转倒计时
	RedirectCountdown int32 `json:"redirectCountdown,omitempty"`
	// 额外的配置信息，格式为json序列化后的字符串
	// OIDC: client_id, scopes, issuer, well_known_url, max_clock_skew
	// feishu: app_id
	Extra string `json:"extra,omitempty"`
	// 凭证，格式为json序列化后的字符串
	// OIDC: client_secret
	// feishu: app_secret
	Credentials string `json:"credentials,omitempty"`
	// 回调链接，我方地址
	CallbackUrl string `json:"callbackUrl,omitempty"`
	// 跳转地址，发起认证用，我方地址
	RedirectUrl string `json:"redirectUrl,omitempty"`
	// 新建用户归属部门
	DefaultDepartment string `json:"defaultDepartment"`
}

// AttrMappingRules user attr match mapping rules
type AttrMappingRules struct {
	MatchMappings []*AttrMapping `json:"matchMappings,omitempty"`
}

// AttrMapping 属性映射
type AttrMapping struct {
	Type MappingType `json:"type,omitempty"`
	// 字段匹配、映射的源字段
	Source string `json:"source,omitempty"`
	// 字段匹配、映射的目标字段
	Target   string `json:"target,omitempty"`
	Required bool   `json:"required,omitempty"`
	Extra    string `json:"extra,omitempty"`
}

// MappingType 属性映射类型
type MappingType int8
