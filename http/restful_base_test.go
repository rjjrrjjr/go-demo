package http

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestPost(t *testing.T) {
	doPost(0)
}

func TestBatchPost(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func(index int) {
			doPost(index)
		}(i)
	}
	time.Sleep(time.Second * 10)
}

func doPostD(traceId int)  {
	// 创建一个用于发送请求的http.Client
	client := &http.Client{}

	// 创建POST请求的URL、请求参数和请求体
	url := "https://meeting.tencent.com/wemeet-tapi/enroll/v1/setAnswers?c_ua=Mozilla/5.0%20(Macintosh;%20Intel%20Mac%20OS%20X%2010_15_7)%20AppleWebKit/605.1.15%20(KHTML,%20like%20Gecko)%20Version/14.1%20Safari/605.1.15&c_wemeet=aHR0cHM6Ly9tZWV0aW5nLnRlbmNlbnQuY29tL2R3L012UFVoRXg2aGUxQQ==&c_instance_id_org=&c_sdk_id=&c_instance_id=5&c_os_model=web&c_os=web&c_nonce=fBXaE5if&c_os_version=Mozilla%252F5.0%2520(Macintosh%253B%2520Intel%2520Mac%2520OS%2520X%252010_15_7)%2520AppleWebKit%252F605.1.15%2520(KHTML%252C%2520like%2520Gecko)%2520Version%252F14.1%2520Safari%252F605.1.15&c_timestamp=1688645869700&c_app_id=200000001&c_app_uid=144115262532588291&c_platform=0&c_token=&c_app_version=&platform=Web&app_id=200000001"
	requestBody := []byte(`{"meeting_id":"9650021035566862293","user_name":"amlu","answer_list":[],"isMp":false,"utm_source":"","utm_medium":"","utm_campaign":"","district":"0"}`)

	// 创建一个新的请求对象
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "_ga_6WSZ0YS5ZQ=GS1.1.1688613169.3.0.1688613169.0.0.0; _ga=GA1.2.1925374862." +
		"1688122087; ts_uid=3674862090; pgv_pvid=9105898336; app_uid=144115262532588291; corp_id=200000001; token_expire_time=1688734037; we_meet_token=eJw0j99umzAchd-Ft0ybbTBxkXZBKhgjjDWUsmY3iD8OdUmMwQRCpr17BSW*-L5z5N-5B*Lg*SsXqs9EwVJeAgsQ8GWhmZTpZSHIMBAi2MREx4RS-IDWiCrrdI4tKQw-312yq*QdS7NjzzpgARMaFMLVDaxTvBFLC2NIIZmbq8wvigumFLDAyPIV9vzM5lNMShF*wPpm5aem4iLtJzlb*dYIlko1rrJkA7-PWlHHqs*PqXkfwStggcS7xc97W9lOk4ShOFR6WdWN4Q27F-8UsS1rA*eHkeDdFBcded*hy-6FOPGtLQZW2GWr4Csxtzf4u9RGw54ikeun4PVbdCRvUasPiDyem8nz-ITTgzbVj-KPm-*9brb2Md*f1VN4yH6OfnitCR90zelaR5FAOm7m*m7n4nf36ZfGN*aINFx9B-8-AgAA--9*I4T-; landing_path=https://meeting.tencent.com/; landing_referraldomain=; landing_referralurl=; landing_url=https://meeting.tencent.com/; web_uid=2e48ea0e-ffb5-42b3-9d87-39aad6394543")

	fmt.Println("=================================begin ", traceId)
	// 发送请求并获取响应
	resp, err := client.Do(req)
	fmt.Println("=================================end ", traceId)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容失败:", err)
		return
	}

	// 输出响应结果
	fmt.Println("响应状态码:", resp.StatusCode)
	fmt.Println("响应内容:", string(body))
}

func doPost(traceId int) {
	// 创建一个用于发送请求的http.Client
	client := &http.Client{}

	// 创建POST请求的URL、请求参数和请求体
	url := "https://meeting.tencent.com/wemeet-tapi/enroll/v1/setAnswers?c_ua=Mozilla/5.0%20(Macintosh;%20Intel%20Mac%20OS%20X%2010_15_7)%20AppleWebKit/605.1.15%20(KHTML,%20like%20Gecko)%20Version/14.1%20Safari/605.1.15&c_wemeet=aHR0cHM6Ly9tZWV0aW5nLnRlbmNlbnQuY29tL2R3L0k3RTJWVVlSRVJmVw==&c_instance_id_org=&c_sdk_id=&c_instance_id=5&c_os_model=web&c_os=web&c_nonce=GhBrD25m&c_os_version=Mozilla%252F5.0%2520(Macintosh%253B%2520Intel%2520Mac%2520OS%2520X%252010_15_7)%2520AppleWebKit%252F605.1.15%2520(KHTML%252C%2520like%2520Gecko)%2520Version%252F14.1%2520Safari%252F605.1.15&c_timestamp=1688728433319&c_app_id=200000001&c_app_uid=144115262532588291&c_platform=0&c_token=&c_app_version=&platform=Web&app_id=200000001"
	requestBody := []byte(`{"meeting_id":"17766249734471345333","user_name":"` + base64.StdEncoding.EncodeToString(
		[]byte(strconv.FormatInt(int64(traceId), 10))) + `",
"answer_list":[],
"isMp":false,"utm_source":"","utm_medium":"","utm_campaign":"","district":"0"}`)

	// 创建一个新的请求对象
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "_ga_6WSZ0YS5ZQ=GS1.1.1688613169.3.0.1688613169.0.0.0; _ga=GA1.2.1925374862.1688122087; ts_uid=3674862090; pgv_pvid=9105898336; app_uid=144115262532588291; corp_id=200000001; token_expire_time=1688734037; we_meet_token=eJw0j99umzAchd-Ft0ybbTBxkXZBKhgjjDWUsmY3iD8OdUmMwQRCpr17BSW*-L5z5N-5B*Lg*SsXqs9EwVJeAgsQ8GWhmZTpZSHIMBAi2MREx4RS-IDWiCrrdI4tKQw-312yq*QdS7NjzzpgARMaFMLVDaxTvBFLC2NIIZmbq8wvigumFLDAyPIV9vzM5lNMShF*wPpm5aem4iLtJzlb*dYIlko1rrJkA7-PWlHHqs*PqXkfwStggcS7xc97W9lOk4ShOFR6WdWN4Q27F-8UsS1rA*eHkeDdFBcded*hy-6FOPGtLQZW2GWr4Csxtzf4u9RGw54ikeun4PVbdCRvUasPiDyem8nz-ITTgzbVj-KPm-*9brb2Md*f1VN4yH6OfnitCR90zelaR5FAOm7m*m7n4nf36ZfGN*aINFx9B-8-AgAA--9*I4T-; landing_path=https://meeting.tencent.com/; landing_referraldomain=; landing_referralurl=; landing_url=https://meeting.tencent.com/; web_uid=2e48ea0e-ffb5-42b3-9d87-39aad6394543")

	fmt.Println("=================================begin ", traceId)
	// 发送请求并获取响应
	resp, err := client.Do(req)
	fmt.Println("=================================end ", traceId)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容失败:", err)
		return
	}

	// 输出响应结果
	fmt.Println(traceId, "响应状态码:", resp.StatusCode)
	fmt.Println(traceId, "响应内容:", string(body))
}

func doPostInENV120(traceId int) {
	// 创建一个用于发送请求的http.Client
	client := &http.Client{}

	// 创建POST请求的URL、请求参数和请求体
	url := "https://test-120.cicd.tencentmeeting.com/wemeet-tapi/enroll/v1/setAnswers?c_ua=Mozilla/5.0%20(Macintosh;%20Intel%20Mac%20OS%20X%2010_15_7)%20AppleWebKit/605.1.15%20(KHTML,%20like%20Gecko)%20Version/14.1%20Safari/605.1.15&c_wemeet=aHR0cHM6Ly90ZXN0LTEyMC5jaWNkLnRlbmNlbnRtZWV0aW5nLmNvbS9kdy9nRjdUQ0NXN3RGUWc/c3luY190b19tZWV0aW5nPXRydWUmZnJvbV9zb3VyY2U9ZnJvbV9oNV9zYXZlX21lZXRpbmdfY2xpY2s=&c_instance_id_org=&c_sdk_id=&c_instance_id=5&c_os_model=web&c_os=web&c_nonce=3nkYh2BJ&c_os_version=Mozilla%252F5.0%2520(Macintosh%253B%2520Intel%2520Mac%2520OS%2520X%252010_15_7)%2520AppleWebKit%252F605.1.15%2520(KHTML%252C%2520like%2520Gecko)%2520Version%252F14.1%2520Safari%252F605.1.15&c_timestamp=1688711172759&c_app_id=200000001&c_app_uid=144115237561191187&c_platform=0&c_token=&c_app_version=&platform=Web&app_id=200000001"
	requestBody := []byte(`{"meeting_id":"16314652493248479800","user_name":"` + base64.StdEncoding.EncodeToString(
		[]byte(strconv.FormatInt(int64(traceId), 10))) + `",
"answer_list":[],
"isMp":false,"utm_source":"","utm_medium":"","utm_campaign":"","district":"0"}`)

	// 创建一个新的请求对象
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "_ga=GA1.2.1717795211.1688711145; _ga_6WSZ0YS5ZQ=GS1.1.1688711145.1.1.1688711170.0.0.0; _gid=GA1.2.485213432.1688711145; ACTIVITY_TICKET=eJxMUF1vm0AQ-C-3XDUsx3FgqQ9pSVMZLMfGUQIvJz7OcDUcVzgwuOp-r4JByr7tzO7O7PxFpyD8mmRZ00vN9KQ42iADfZlhkXOpxVnwFm2QadwLGFgWADExJTaAC%2ADQZT5RSuQs0Qy3%2AaczXX5hM4U2CAyg1HUxXjg%2AKtFylpz1LAKOsW4NvO1EI2dl0zQcg3yoL6QW9YdRsB2HAgBd8aophFzfUGUjOVPddfUhCrRBfq%2ArmKiwiv2o3v2sp8Zr3s7Db7KtH16v77dAXVKq4OnZCU6l1UcHm9Q6eTqkfeWG0vwRiEnupb3Pg20h8a83T46di7NbjLOS0-F6tOsotoKjotzXOPFeHl4P4P%2A5iJfSsLbvKbHD3VQ-9sSLqQpvzVB8P9nuY3jcu8QTZYKjZz6GWbqLo2%2AL8ZYX9ySWPueDyDibA0X--gcAAP--bCuJJw__; app_uid=144115237561191187; corp_id=200000001; token_expire_time=1689315970; we_meet_token=eJw0j11vmzAUQP*LXzNtNsXGQ9pDmox0EEgyOtTtBREwYAPmw4Qwpv33Ckr8eM658r3-wOvR-8yl6iMZs5AnwAQYfFpo1DThbSFI1xHC2pOBCUJfEaLGmqikCOdsqTT48dAq2djwjoVR2rMOmIBAnUK4uoF1itdymdI0SCGeJ1d5vSkumVLABHd2XWHPKzavQig1EELGIy7rjMuw-9vMtslrycJG3VeZsIE-zlpRx7KPjyl5HMEzYILkZdyJUXmoarPfuEkO-i-SUvjTq-*cTyKRjn1IxSVl*kbm-fX0g6K2jsXLlmRx-UUEQeHkF*uGcDkZ-BThvMqDt-KMM6I62-ruiN1*SF0rfRbDuB*I741bx3KPb7Z2KF-1ttu4zwO7bJ886MZTYk*7-UYUFa*y44ac74XURcD8NpgCDcJv4P97AAAA--*-A4YR; _gat_UA-205111495-1=1; landing_path=https://test-120.cicd.tencentmeeting.com/login.html; landing_referraldomain=https://test-120.cicd.tencentmeeting.com; landing_referralurl=https://test-120.cicd.tencentmeeting.com/dw/gF7TCCW7tFQg; landing_url=https://test-120.cicd.tencentmeeting.com/login.html?m_login_source_id=2&redirect_link=dw%2FgF7TCCW7tFQg%3Flogin_type%3D0%26sync_to_meeting%3Dtrue%26from_source%3Dfrom_h5_save_meeting_click%26extra_params%3Dsignup_click&from_source=from_h5_save_meeting_click&pagetype=2&opetype=1; web_uid=dab872a8-b66e-40fc-97d2-9f186c17fd46")

	fmt.Println("=================================begin ", traceId)
	// 发送请求并获取响应
	resp, err := client.Do(req)
	fmt.Println("=================================end ", traceId)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应内容失败:", err)
		return
	}

	// 输出响应结果
	fmt.Println(traceId, "响应状态码:", resp.StatusCode)
	fmt.Println(traceId, "响应内容:", string(body))
}