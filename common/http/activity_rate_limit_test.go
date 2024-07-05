package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

// https://pre.txmeeting.tencent.com/wemeet-tapi/v2/activities/dashboard/guest/add-guest?c_instance_id=5&c_os=web&c_os_model=web&c_os_version=Mozilla%252F5.0%2520(Macintosh%253B%2520Intel%2520Mac%2520OS%2520X%252010_15_7)%2520AppleWebKit%252F537.36%2520(KHTML%252C%2520like%2520Gecko)%2520Chrome%252F122.0.0.0%2520Safari%252F537.36&c_nonce=1710830426765&c_timestamp=1710830426765&c_app_id=200000001&c_app_uid=144115262788474108&c_platform=0&c_token=&c_app_version=&platform=Web&app_id=200000001&c_lang=zh-CN

func TestRateLimit(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func(index int) {
			addOrgMember(index)
		}(i)
	}

	time.Sleep(time.Second * 10)
}

func addGuest(traceId int) {

	// 创建一个用于发送请求的http.Client
	client := &http.Client{}

	// 创建POST请求的URL、请求参数和请求体
	url := "https://pre.txmeeting.tencent.com/wemeet-tapi/v2/activities/dashboard/guest/add-guest?c_instance_id=5&c_os=web&c_os_model=web&c_os_version=Mozilla%252F5.0%2520(Macintosh%253B%2520Intel%2520Mac%2520OS%2520X%252010_15_7)%2520AppleWebKit%252F537.36%2520(KHTML%252C%2520like%2520Gecko)%2520Chrome%252F122.0.0.0%2520Safari%252F537.36&c_nonce=1710830426765&c_timestamp=1710830426765&c_app_id=200000001&c_app_uid=144115262788474108&c_platform=0&c_token=&c_app_version=&platform=Web&app_id=200000001&c_lang=zh-CN"
	requestBody := []byte(`{"name":"jin","email":"","phone":"15669489934","company":"CAO","position":"",
"area_code":"+86","describe":"","invite_type":0,"invite_code":"1769979522967068672","event_id":"E176997694423985356874108","avatar":"","avatar_cos_key":"","host_id":"8651634848540374108"}`)

	// 创建一个新的请求对象
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "pgv_pvid=5818805490; ab_uuid=20e67858f47142468f5c14b81cd160a5; _ga_J5ZYJSX9HN=GS1.1.1650597486.1.0.1650597503.0; web_uid=389dd316-9ce4-4985-9ab1-3b347744420a; landing_url=https://pre.txmeeting.tencent.com/qrcode-login.html?redirect_type=2&redirect_link=events-lobby%2FE165865955832550195210041%3Flogin_type%3D0; landing_path=https://pre.txmeeting.tencent.com/qrcode-login.html; landing_referralurl=https://pre.txmeeting.tencent.com/events-lobby/E165865955832550195210041; landing_referraldomain=https://pre.txmeeting.tencent.com; _qimei_uuid42=181020e2c1310082ad7d9f3cf146dbb693b7d0afda; _qimei_q36=; _qimei_h38=7e79c3e7712dd761818677250300000f917904; _gcl_au=1.1.2122262273.1706153534; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%22100024540397%22%2C%22first_id%22%3A%2218c6b6132dab4c-0efb27669c56ce8-1f525637-3686400-18c6b6132dbdd6%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfY29va2llX2lkIjoiMThjNmI2MTMyZGFiNGMtMGVmYjI3NjY5YzU2Y2U4LTFmNTI1NjM3LTM2ODY0MDAtMThjNmI2MTMyZGJkZDYiLCIkaWRlbnRpdHlfbG9naW5faWQiOiIxMDAwMjQ1NDAzOTcifQ%3D%3D%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%22100024540397%22%7D%2C%22%24device_id%22%3A%2218c6b6132dab4c-0efb27669c56ce8-1f525637-3686400-18c6b6132dbdd6%22%7D; pgv_info=ssid=s119735884; _gid=GA1.2.1692842897.1710830199; _qimei_fingerprint=1f52e8a2f7c812be4941e1f63b7c116e; we_meet_token=eJxUkF1vmzAUhv*Lb7Os-gq2I*2iE1Fo1gitYcvHDXLghHgQcDA0pNP**xRK1ZY7nud9j338F0WPq6*mdI0uE4hNiqZogr70VFsbtz0hnBMyoR4VUnLBCZZDxKV5fIv1KYpfPzJI6KypIdaHBmo0RR7mEuPBPUPtTFX2LcoIJuLWHOS*daYE59AUXWA-wMac4HYVQbBkmCpv4EWVmTJurvZm7bEqIbbuMsgUns3bWgOqIXs9WL5NcCZDU3Sad-YXlrsw*WGjs2F*twy-B5tkXTF6p6*yOOXLn611i4XZzPTLwS3lmkUBz66riWj80UsqOrusqMmL6DF80rOOzDy1n8-Jlu7c7-v1cbe3NKBUhosR-6MAm6fNA6iiLNpDe17fi1FeFXUebkbbmmcqupvQB94c05X2L16gtr7waR76AZt9*-D6H1c7mxMY5r0DW1dpmzSfQq7-0YoyCoSMGVA85qBgrKVWY5ImnMFBcOrh9wKr*wH--gcAAP--prGkJA__; corp_id=200000001; app_uid=144115262788474108; token_expire_time=1711435096; ACTIVITY_TICKET=eJxMUEtz2jAY-C%2A6ttPRwyDhmR4oqXnEIQVDgV40Aku2Yiz5IRuSTv97J47JoNvuamf3279gE0bfxOlkG%2AO4ey0k8AH42rE6lsZppWUFfIDhx0MceR5CAzzElDGPegiy-r8oCh1z4TipYuAD2NN1nPFOAj5AEFE6GhHSa-Ja6EpyoVwXghi8uVpZ1dqaLhkTSPHgPb0Xnc7feyKKICMQj4Y9f7aJNrcritQayYv6cuuhE%2ACDjFYkmh7L7X5Hqi-FDL1sDoPjmIqfT9eH-Md8dLyO3aNYZdNDsFa0WeM5zM%2AHVagedNics3ONBh7c1c-7Uiy13NrJWj0n47ZdDudQRZOMLZ7Sizr%2AthvmLfDkFw3NIS3-aLMohxEJ63xx2gaevswe0XSXrKBQsrH7NGokSoKXIH9Ll6-BDG9Z-b0vXsnkY4kex7LVJ8m7Qe82vsfOZtLwT2PyBv79DwAA--%2AZfpT4; _ga=GA1.2.204146153.1705055152; _ga_6WSZ0YS5ZQ=GS1.1.1710830297.9.1.1710830316.0.0.0")

	// 发送请求并获取响应
	resp, err := client.Do(req)
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

func createScheduler(traceId int) {
	// 创建一个用于发送请求的http.Client
	client := &http.Client{}

	// 创建POST请求的URL、请求参数和请求体
	url := "https://pre.txmeeting.tencent.com/wemeet-tapi/v2/activities/dashboard/scheduler/create?c_instance_id=5&c_os=web&c_os_model=web&c_os_version=Mozilla%252F5.0%2520(Macintosh%253B%2520Intel%2520Mac%2520OS%2520X%252010_15_7)%2520AppleWebKit%252F537.36%2520(KHTML%252C%2520like%2520Gecko)%2520Chrome%252F122.0.0.0%2520Safari%252F537.36&c_nonce=1710831311968&c_timestamp=1710831311968&c_app_id=200000001&c_app_uid=144115262788474108&c_platform=0&c_token=&c_app_version=&platform=Web&app_id=200000001&c_lang=zh-CN"

	requestBody := []byte(`{"scheduler_req_data":{"name":"jin","begin_time":1711436400,"end_time":1711438200,"guest_ids":[],"host_ids":["452664"],"description":"","pics_source_list":[],"plans":[],"venue_template_id":"820851","venue_standard":"Stage200","scheduler_id":0},"event_id":"E176997694423985356874108"}`)

	// 创建一个新的请求对象
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "pgv_pvid=5818805490; ab_uuid=20e67858f47142468f5c14b81cd160a5; _ga_J5ZYJSX9HN=GS1.1.1650597486.1.0.1650597503.0; web_uid=389dd316-9ce4-4985-9ab1-3b347744420a; landing_url=https://pre.txmeeting.tencent.com/qrcode-login.html?redirect_type=2&redirect_link=events-lobby%2FE165865955832550195210041%3Flogin_type%3D0; landing_path=https://pre.txmeeting.tencent.com/qrcode-login.html; landing_referralurl=https://pre.txmeeting.tencent.com/events-lobby/E165865955832550195210041; landing_referraldomain=https://pre.txmeeting.tencent.com; _qimei_uuid42=181020e2c1310082ad7d9f3cf146dbb693b7d0afda; _qimei_q36=; _qimei_h38=7e79c3e7712dd761818677250300000f917904; _gcl_au=1.1.2122262273.1706153534; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%22100024540397%22%2C%22first_id%22%3A%2218c6b6132dab4c-0efb27669c56ce8-1f525637-3686400-18c6b6132dbdd6%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfY29va2llX2lkIjoiMThjNmI2MTMyZGFiNGMtMGVmYjI3NjY5YzU2Y2U4LTFmNTI1NjM3LTM2ODY0MDAtMThjNmI2MTMyZGJkZDYiLCIkaWRlbnRpdHlfbG9naW5faWQiOiIxMDAwMjQ1NDAzOTcifQ%3D%3D%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%22100024540397%22%7D%2C%22%24device_id%22%3A%2218c6b6132dab4c-0efb27669c56ce8-1f525637-3686400-18c6b6132dbdd6%22%7D; pgv_info=ssid=s119735884; _gid=GA1.2.1692842897.1710830199; _qimei_fingerprint=1f52e8a2f7c812be4941e1f63b7c116e; we_meet_token=eJxUkF1vmzAUhv*Lb7Os-gq2I*2iE1Fo1gitYcvHDXLghHgQcDA0pNP**xRK1ZY7nud9j338F0WPq6*mdI0uE4hNiqZogr70VFsbtz0hnBMyoR4VUnLBCZZDxKV5fIv1KYpfPzJI6KypIdaHBmo0RR7mEuPBPUPtTFX2LcoIJuLWHOS*daYE59AUXWA-wMac4HYVQbBkmCpv4EWVmTJurvZm7bEqIbbuMsgUns3bWgOqIXs9WL5NcCZDU3Sad-YXlrsw*WGjs2F*twy-B5tkXTF6p6*yOOXLn611i4XZzPTLwS3lmkUBz66riWj80UsqOrusqMmL6DF80rOOzDy1n8-Jlu7c7-v1cbe3NKBUhosR-6MAm6fNA6iiLNpDe17fi1FeFXUebkbbmmcqupvQB94c05X2L16gtr7waR76AZt9*-D6H1c7mxMY5r0DW1dpmzSfQq7-0YoyCoSMGVA85qBgrKVWY5ImnMFBcOrh9wKr*wH--gcAAP--prGkJA__; corp_id=200000001; app_uid=144115262788474108; token_expire_time=1711435096; ACTIVITY_TICKET=eJxMUEtz2jAY-C%2A6ttPRwyDhmR4oqXnEIQVDgV40Aku2Yiz5IRuSTv97J47JoNvuamf3279gE0bfxOlkG%2AO4ey0k8AH42rE6lsZppWUFfIDhx0MceR5CAzzElDGPegiy-r8oCh1z4TipYuAD2NN1nPFOAj5AEFE6GhHSa-Ja6EpyoVwXghi8uVpZ1dqaLhkTSPHgPb0Xnc7feyKKICMQj4Y9f7aJNrcritQayYv6cuuhE%2ACDjFYkmh7L7X5Hqi-FDL1sDoPjmIqfT9eH-Md8dLyO3aNYZdNDsFa0WeM5zM%2AHVagedNics3ONBh7c1c-7Uiy13NrJWj0n47ZdDudQRZOMLZ7Sizr%2AthvmLfDkFw3NIS3-aLMohxEJ63xx2gaevswe0XSXrKBQsrH7NGokSoKXIH9Ll6-BDG9Z-b0vXsnkY4kex7LVJ8m7Qe82vsfOZtLwT2PyBv79DwAA--%2AZfpT4; _ga=GA1.2.204146153.1705055152; _ga_6WSZ0YS5ZQ=GS1.1.1710830297.9.1.1710830316.0.0.0")

	// 发送请求并获取响应
	resp, err := client.Do(req)
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

func addEventHost(traceId int) {
	// 创建一个用于发送请求的http.Client
	client := &http.Client{}

	// 创建POST请求的URL、请求参数和请求体
	url := "https://pre.txmeeting.tencent.com/wemeet-tapi/v2/activities/dashboard/host/add-host?c_instance_id=5&c_os=web&c_os_model=web&c_os_version=Mozilla%252F5.0%2520(Macintosh%253B%2520Intel%2520Mac%2520OS%2520X%252010_15_7)%2520AppleWebKit%252F537.36%2520(KHTML%252C%2520like%2520Gecko)%2520Chrome%252F122.0.0.0%2520Safari%252F537.36&c_nonce=1710831581728&c_timestamp=1710831581728&c_app_id=200000001&c_app_uid=144115262788474108&c_platform=0&c_token=&c_app_version=&platform=Web&app_id=200000001&c_lang=zh-CN"

	requestBody := []byte(`{"event_id":"E176997694423985356874108","name":"jin","email":"","area":"86",
"phone":"15669489941","scheduler_id":"1769980927307898880"}`)

	// 创建一个新的请求对象
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "pgv_pvid=5818805490; ab_uuid=20e67858f47142468f5c14b81cd160a5; _ga_J5ZYJSX9HN=GS1.1.1650597486.1.0.1650597503.0; web_uid=389dd316-9ce4-4985-9ab1-3b347744420a; landing_url=https://pre.txmeeting.tencent.com/qrcode-login.html?redirect_type=2&redirect_link=events-lobby%2FE165865955832550195210041%3Flogin_type%3D0; landing_path=https://pre.txmeeting.tencent.com/qrcode-login.html; landing_referralurl=https://pre.txmeeting.tencent.com/events-lobby/E165865955832550195210041; landing_referraldomain=https://pre.txmeeting.tencent.com; _qimei_uuid42=181020e2c1310082ad7d9f3cf146dbb693b7d0afda; _qimei_q36=; _qimei_h38=7e79c3e7712dd761818677250300000f917904; _gcl_au=1.1.2122262273.1706153534; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%22100024540397%22%2C%22first_id%22%3A%2218c6b6132dab4c-0efb27669c56ce8-1f525637-3686400-18c6b6132dbdd6%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfY29va2llX2lkIjoiMThjNmI2MTMyZGFiNGMtMGVmYjI3NjY5YzU2Y2U4LTFmNTI1NjM3LTM2ODY0MDAtMThjNmI2MTMyZGJkZDYiLCIkaWRlbnRpdHlfbG9naW5faWQiOiIxMDAwMjQ1NDAzOTcifQ%3D%3D%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%22100024540397%22%7D%2C%22%24device_id%22%3A%2218c6b6132dab4c-0efb27669c56ce8-1f525637-3686400-18c6b6132dbdd6%22%7D; pgv_info=ssid=s119735884; _gid=GA1.2.1692842897.1710830199; _qimei_fingerprint=1f52e8a2f7c812be4941e1f63b7c116e; we_meet_token=eJxUkF1vmzAUhv*Lb7Os-gq2I*2iE1Fo1gitYcvHDXLghHgQcDA0pNP**xRK1ZY7nud9j338F0WPq6*mdI0uE4hNiqZogr70VFsbtz0hnBMyoR4VUnLBCZZDxKV5fIv1KYpfPzJI6KypIdaHBmo0RR7mEuPBPUPtTFX2LcoIJuLWHOS*daYE59AUXWA-wMac4HYVQbBkmCpv4EWVmTJurvZm7bEqIbbuMsgUns3bWgOqIXs9WL5NcCZDU3Sad-YXlrsw*WGjs2F*twy-B5tkXTF6p6*yOOXLn611i4XZzPTLwS3lmkUBz66riWj80UsqOrusqMmL6DF80rOOzDy1n8-Jlu7c7-v1cbe3NKBUhosR-6MAm6fNA6iiLNpDe17fi1FeFXUebkbbmmcqupvQB94c05X2L16gtr7waR76AZt9*-D6H1c7mxMY5r0DW1dpmzSfQq7-0YoyCoSMGVA85qBgrKVWY5ImnMFBcOrh9wKr*wH--gcAAP--prGkJA__; corp_id=200000001; app_uid=144115262788474108; token_expire_time=1711435096; ACTIVITY_TICKET=eJxMUEtz2jAY-C%2A6ttPRwyDhmR4oqXnEIQVDgV40Aku2Yiz5IRuSTv97J47JoNvuamf3279gE0bfxOlkG%2AO4ey0k8AH42rE6lsZppWUFfIDhx0MceR5CAzzElDGPegiy-r8oCh1z4TipYuAD2NN1nPFOAj5AEFE6GhHSa-Ja6EpyoVwXghi8uVpZ1dqaLhkTSPHgPb0Xnc7feyKKICMQj4Y9f7aJNrcritQayYv6cuuhE%2ACDjFYkmh7L7X5Hqi-FDL1sDoPjmIqfT9eH-Md8dLyO3aNYZdNDsFa0WeM5zM%2AHVagedNics3ONBh7c1c-7Uiy13NrJWj0n47ZdDudQRZOMLZ7Sizr%2AthvmLfDkFw3NIS3-aLMohxEJ63xx2gaevswe0XSXrKBQsrH7NGokSoKXIH9Ll6-BDG9Z-b0vXsnkY4kex7LVJ8m7Qe82vsfOZtLwT2PyBv79DwAA--%2AZfpT4; _ga=GA1.2.204146153.1705055152; _ga_6WSZ0YS5ZQ=GS1.1.1710830297.9.1.1710830316.0.0.0")

	// 发送请求并获取响应
	resp, err := client.Do(req)
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

func setEventHost(traceId int) {
	// 创建一个用于发送请求的http.Client
	client := &http.Client{}

	// 创建POST请求的URL、请求参数和请求体
	url := "https://pre.txmeeting.tencent.com/wemeet-tapi/v2/activities/dashboard/scheduler/set-event-host?c_instance_id=5&c_os=web&c_os_model=web&c_os_version=Mozilla%252F5.0%2520(Macintosh%253B%2520Intel%2520Mac%2520OS%2520X%252010_15_7)%2520AppleWebKit%252F537.36%2520(KHTML%252C%2520like%2520Gecko)%2520Chrome%252F122.0.0.0%2520Safari%252F537.36&c_nonce=1710831742441&c_timestamp=1710831742441&c_app_id=200000001&c_app_uid=144115262788474108&c_platform=0&c_token=&c_app_version=&platform=Web&app_id=200000001&c_lang=zh-CN"

	requestBody := []byte(`{"event_id":"E176997694423985356874108","scheduler_id":"1769981562782703616","event_host_id_list":[452664,452665]}`)

	// 创建一个新的请求对象
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "pgv_pvid=5818805490; ab_uuid=20e67858f47142468f5c14b81cd160a5; _ga_J5ZYJSX9HN=GS1.1.1650597486.1.0.1650597503.0; web_uid=389dd316-9ce4-4985-9ab1-3b347744420a; landing_url=https://pre.txmeeting.tencent.com/qrcode-login.html?redirect_type=2&redirect_link=events-lobby%2FE165865955832550195210041%3Flogin_type%3D0; landing_path=https://pre.txmeeting.tencent.com/qrcode-login.html; landing_referralurl=https://pre.txmeeting.tencent.com/events-lobby/E165865955832550195210041; landing_referraldomain=https://pre.txmeeting.tencent.com; _qimei_uuid42=181020e2c1310082ad7d9f3cf146dbb693b7d0afda; _qimei_q36=; _qimei_h38=7e79c3e7712dd761818677250300000f917904; _gcl_au=1.1.2122262273.1706153534; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%22100024540397%22%2C%22first_id%22%3A%2218c6b6132dab4c-0efb27669c56ce8-1f525637-3686400-18c6b6132dbdd6%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfY29va2llX2lkIjoiMThjNmI2MTMyZGFiNGMtMGVmYjI3NjY5YzU2Y2U4LTFmNTI1NjM3LTM2ODY0MDAtMThjNmI2MTMyZGJkZDYiLCIkaWRlbnRpdHlfbG9naW5faWQiOiIxMDAwMjQ1NDAzOTcifQ%3D%3D%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%22100024540397%22%7D%2C%22%24device_id%22%3A%2218c6b6132dab4c-0efb27669c56ce8-1f525637-3686400-18c6b6132dbdd6%22%7D; pgv_info=ssid=s119735884; _gid=GA1.2.1692842897.1710830199; _qimei_fingerprint=1f52e8a2f7c812be4941e1f63b7c116e; we_meet_token=eJxUkF1vmzAUhv*Lb7Os-gq2I*2iE1Fo1gitYcvHDXLghHgQcDA0pNP**xRK1ZY7nud9j338F0WPq6*mdI0uE4hNiqZogr70VFsbtz0hnBMyoR4VUnLBCZZDxKV5fIv1KYpfPzJI6KypIdaHBmo0RR7mEuPBPUPtTFX2LcoIJuLWHOS*daYE59AUXWA-wMac4HYVQbBkmCpv4EWVmTJurvZm7bEqIbbuMsgUns3bWgOqIXs9WL5NcCZDU3Sad-YXlrsw*WGjs2F*twy-B5tkXTF6p6*yOOXLn611i4XZzPTLwS3lmkUBz66riWj80UsqOrusqMmL6DF80rOOzDy1n8-Jlu7c7-v1cbe3NKBUhosR-6MAm6fNA6iiLNpDe17fi1FeFXUebkbbmmcqupvQB94c05X2L16gtr7waR76AZt9*-D6H1c7mxMY5r0DW1dpmzSfQq7-0YoyCoSMGVA85qBgrKVWY5ImnMFBcOrh9wKr*wH--gcAAP--prGkJA__; corp_id=200000001; app_uid=144115262788474108; token_expire_time=1711435096; ACTIVITY_TICKET=eJxMUEtz2jAY-C%2A6ttPRwyDhmR4oqXnEIQVDgV40Aku2Yiz5IRuSTv97J47JoNvuamf3279gE0bfxOlkG%2AO4ey0k8AH42rE6lsZppWUFfIDhx0MceR5CAzzElDGPegiy-r8oCh1z4TipYuAD2NN1nPFOAj5AEFE6GhHSa-Ja6EpyoVwXghi8uVpZ1dqaLhkTSPHgPb0Xnc7feyKKICMQj4Y9f7aJNrcritQayYv6cuuhE%2ACDjFYkmh7L7X5Hqi-FDL1sDoPjmIqfT9eH-Md8dLyO3aNYZdNDsFa0WeM5zM%2AHVagedNics3ONBh7c1c-7Uiy13NrJWj0n47ZdDudQRZOMLZ7Sizr%2AthvmLfDkFw3NIS3-aLMohxEJ63xx2gaevswe0XSXrKBQsrH7NGokSoKXIH9Ll6-BDG9Z-b0vXsnkY4kex7LVJ8m7Qe82vsfOZtLwT2PyBv79DwAA--%2AZfpT4; _ga=GA1.2.204146153.1705055152; _ga_6WSZ0YS5ZQ=GS1.1.1710830297.9.1.1710830316.0.0.0")

	// 发送请求并获取响应
	resp, err := client.Do(req)
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

func batchCopyScheduler(traceId int) {
	// 创建一个用于发送请求的http.Client
	client := &http.Client{}

	// 创建POST请求的URL、请求参数和请求体
	url := "https://pre.txmeeting.tencent.com/wemeet-tapi/v2/activities/dashboard/scheduler/batch-copy?c_instance_id=5&c_os=web&c_os_model=web&c_os_version=Mozilla%252F5.0%2520(Macintosh%253B%2520Intel%2520Mac%2520OS%2520X%252010_15_7)%2520AppleWebKit%252F537.36%2520(KHTML%252C%2520like%2520Gecko)%2520Chrome%252F122.0.0.0%2520Safari%252F537.36&c_nonce=1710831895036&c_timestamp=1710831895036&c_app_id=200000001&c_app_uid=144115262788474108&c_platform=0&c_token=&c_app_version=&platform=Web&app_id=200000001&c_lang=zh-CN"

	requestBody := []byte(`{"event_id":"E176997694423985356874108","scheduler_id_list":["1769980927307898880","1769981562782703616"]}`)

	// 创建一个新的请求对象
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "pgv_pvid=5818805490; ab_uuid=20e67858f47142468f5c14b81cd160a5; _ga_J5ZYJSX9HN=GS1.1.1650597486.1.0.1650597503.0; web_uid=389dd316-9ce4-4985-9ab1-3b347744420a; landing_url=https://pre.txmeeting.tencent.com/qrcode-login.html?redirect_type=2&redirect_link=events-lobby%2FE165865955832550195210041%3Flogin_type%3D0; landing_path=https://pre.txmeeting.tencent.com/qrcode-login.html; landing_referralurl=https://pre.txmeeting.tencent.com/events-lobby/E165865955832550195210041; landing_referraldomain=https://pre.txmeeting.tencent.com; _qimei_uuid42=181020e2c1310082ad7d9f3cf146dbb693b7d0afda; _qimei_q36=; _qimei_h38=7e79c3e7712dd761818677250300000f917904; _gcl_au=1.1.2122262273.1706153534; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%22100024540397%22%2C%22first_id%22%3A%2218c6b6132dab4c-0efb27669c56ce8-1f525637-3686400-18c6b6132dbdd6%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfY29va2llX2lkIjoiMThjNmI2MTMyZGFiNGMtMGVmYjI3NjY5YzU2Y2U4LTFmNTI1NjM3LTM2ODY0MDAtMThjNmI2MTMyZGJkZDYiLCIkaWRlbnRpdHlfbG9naW5faWQiOiIxMDAwMjQ1NDAzOTcifQ%3D%3D%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%22100024540397%22%7D%2C%22%24device_id%22%3A%2218c6b6132dab4c-0efb27669c56ce8-1f525637-3686400-18c6b6132dbdd6%22%7D; pgv_info=ssid=s119735884; _gid=GA1.2.1692842897.1710830199; _qimei_fingerprint=1f52e8a2f7c812be4941e1f63b7c116e; we_meet_token=eJxUkF1vmzAUhv*Lb7Os-gq2I*2iE1Fo1gitYcvHDXLghHgQcDA0pNP**xRK1ZY7nud9j338F0WPq6*mdI0uE4hNiqZogr70VFsbtz0hnBMyoR4VUnLBCZZDxKV5fIv1KYpfPzJI6KypIdaHBmo0RR7mEuPBPUPtTFX2LcoIJuLWHOS*daYE59AUXWA-wMac4HYVQbBkmCpv4EWVmTJurvZm7bEqIbbuMsgUns3bWgOqIXs9WL5NcCZDU3Sad-YXlrsw*WGjs2F*twy-B5tkXTF6p6*yOOXLn611i4XZzPTLwS3lmkUBz66riWj80UsqOrusqMmL6DF80rOOzDy1n8-Jlu7c7-v1cbe3NKBUhosR-6MAm6fNA6iiLNpDe17fi1FeFXUebkbbmmcqupvQB94c05X2L16gtr7waR76AZt9*-D6H1c7mxMY5r0DW1dpmzSfQq7-0YoyCoSMGVA85qBgrKVWY5ImnMFBcOrh9wKr*wH--gcAAP--prGkJA__; corp_id=200000001; app_uid=144115262788474108; token_expire_time=1711435096; ACTIVITY_TICKET=eJxMUEtz2jAY-C%2A6ttPRwyDhmR4oqXnEIQVDgV40Aku2Yiz5IRuSTv97J47JoNvuamf3279gE0bfxOlkG%2AO4ey0k8AH42rE6lsZppWUFfIDhx0MceR5CAzzElDGPegiy-r8oCh1z4TipYuAD2NN1nPFOAj5AEFE6GhHSa-Ja6EpyoVwXghi8uVpZ1dqaLhkTSPHgPb0Xnc7feyKKICMQj4Y9f7aJNrcritQayYv6cuuhE%2ACDjFYkmh7L7X5Hqi-FDL1sDoPjmIqfT9eH-Md8dLyO3aNYZdNDsFa0WeM5zM%2AHVagedNics3ONBh7c1c-7Uiy13NrJWj0n47ZdDudQRZOMLZ7Sizr%2AthvmLfDkFw3NIS3-aLMohxEJ63xx2gaevswe0XSXrKBQsrH7NGokSoKXIH9Ll6-BDG9Z-b0vXsnkY4kex7LVJ8m7Qe82vsfOZtLwT2PyBv79DwAA--%2AZfpT4; _ga=GA1.2.204146153.1705055152; _ga_6WSZ0YS5ZQ=GS1.1.1710830297.9.1.1710830316.0.0.0")

	// 发送请求并获取响应
	resp, err := client.Do(req)
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

func importExcelData(traceId int) {
	// 创建一个用于发送请求的http.Client
	client := &http.Client{}

	// 创建POST请求的URL、请求参数和请求体
	url := "https://pre.txmeeting.tencent.com/wemeet-tapi/v2/activities/dashboard/common/import-excel-data?c_instance_id=5&c_os=web&c_os_model=web&c_os_version=Mozilla%252F5.0%2520(Macintosh%253B%2520Intel%2520Mac%2520OS%2520X%252010_15_7)%2520AppleWebKit%252F537.36%2520(KHTML%252C%2520like%2520Gecko)%2520Chrome%252F122.0.0.0%2520Safari%252F537.36&c_nonce=1710832245068&c_timestamp=1710832245068&c_app_id=200000001&c_app_uid=144115262788474108&c_platform=0&c_token=&c_app_version=&platform=Web&app_id=200000001&c_lang=zh-CN"

	requestBody := []byte(`{"cos_id":"activity_temporary/application/vnd.openxmlformats-officedocument.spreadsheetml.sheet/505501933960817591/505501933960883127.xlsx","event_id":"E176997694423985356874108","module":1}`)

	// 创建一个新的请求对象
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "pgv_pvid=5818805490; ab_uuid=20e67858f47142468f5c14b81cd160a5; _ga_J5ZYJSX9HN=GS1.1.1650597486.1.0.1650597503.0; web_uid=389dd316-9ce4-4985-9ab1-3b347744420a; landing_url=https://pre.txmeeting.tencent.com/qrcode-login.html?redirect_type=2&redirect_link=events-lobby%2FE165865955832550195210041%3Flogin_type%3D0; landing_path=https://pre.txmeeting.tencent.com/qrcode-login.html; landing_referralurl=https://pre.txmeeting.tencent.com/events-lobby/E165865955832550195210041; landing_referraldomain=https://pre.txmeeting.tencent.com; _qimei_uuid42=181020e2c1310082ad7d9f3cf146dbb693b7d0afda; _qimei_q36=; _qimei_h38=7e79c3e7712dd761818677250300000f917904; _gcl_au=1.1.2122262273.1706153534; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%22100024540397%22%2C%22first_id%22%3A%2218c6b6132dab4c-0efb27669c56ce8-1f525637-3686400-18c6b6132dbdd6%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfY29va2llX2lkIjoiMThjNmI2MTMyZGFiNGMtMGVmYjI3NjY5YzU2Y2U4LTFmNTI1NjM3LTM2ODY0MDAtMThjNmI2MTMyZGJkZDYiLCIkaWRlbnRpdHlfbG9naW5faWQiOiIxMDAwMjQ1NDAzOTcifQ%3D%3D%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%22100024540397%22%7D%2C%22%24device_id%22%3A%2218c6b6132dab4c-0efb27669c56ce8-1f525637-3686400-18c6b6132dbdd6%22%7D; pgv_info=ssid=s119735884; _gid=GA1.2.1692842897.1710830199; _qimei_fingerprint=1f52e8a2f7c812be4941e1f63b7c116e; we_meet_token=eJxUkF1vmzAUhv*Lb7Os-gq2I*2iE1Fo1gitYcvHDXLghHgQcDA0pNP**xRK1ZY7nud9j338F0WPq6*mdI0uE4hNiqZogr70VFsbtz0hnBMyoR4VUnLBCZZDxKV5fIv1KYpfPzJI6KypIdaHBmo0RR7mEuPBPUPtTFX2LcoIJuLWHOS*daYE59AUXWA-wMac4HYVQbBkmCpv4EWVmTJurvZm7bEqIbbuMsgUns3bWgOqIXs9WL5NcCZDU3Sad-YXlrsw*WGjs2F*twy-B5tkXTF6p6*yOOXLn611i4XZzPTLwS3lmkUBz66riWj80UsqOrusqMmL6DF80rOOzDy1n8-Jlu7c7-v1cbe3NKBUhosR-6MAm6fNA6iiLNpDe17fi1FeFXUebkbbmmcqupvQB94c05X2L16gtr7waR76AZt9*-D6H1c7mxMY5r0DW1dpmzSfQq7-0YoyCoSMGVA85qBgrKVWY5ImnMFBcOrh9wKr*wH--gcAAP--prGkJA__; corp_id=200000001; app_uid=144115262788474108; token_expire_time=1711435096; ACTIVITY_TICKET=eJxMUEtz2jAY-C%2A6ttPRwyDhmR4oqXnEIQVDgV40Aku2Yiz5IRuSTv97J47JoNvuamf3279gE0bfxOlkG%2AO4ey0k8AH42rE6lsZppWUFfIDhx0MceR5CAzzElDGPegiy-r8oCh1z4TipYuAD2NN1nPFOAj5AEFE6GhHSa-Ja6EpyoVwXghi8uVpZ1dqaLhkTSPHgPb0Xnc7feyKKICMQj4Y9f7aJNrcritQayYv6cuuhE%2ACDjFYkmh7L7X5Hqi-FDL1sDoPjmIqfT9eH-Md8dLyO3aNYZdNDsFa0WeM5zM%2AHVagedNics3ONBh7c1c-7Uiy13NrJWj0n47ZdDudQRZOMLZ7Sizr%2AthvmLfDkFw3NIS3-aLMohxEJ63xx2gaevswe0XSXrKBQsrH7NGokSoKXIH9Ll6-BDG9Z-b0vXsnkY4kex7LVJ8m7Qe82vsfOZtLwT2PyBv79DwAA--%2AZfpT4; _ga=GA1.2.204146153.1705055152; _ga_6WSZ0YS5ZQ=GS1.1.1710830297.9.1.1710830316.0.0.0")

	// 发送请求并获取响应
	resp, err := client.Do(req)
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

func addOrgMember(traceId int) {
	// 创建一个用于发送请求的http.Client
	client := &http.Client{}

	// 创建POST请求的URL、请求参数和请求体
	url := "https://pre.txmeeting.tencent.com/wemeet-tapi/v2/activities/dashboard/org/add-org-member?c_instance_id=5&c_os=web&c_os_model=web&c_os_version=Mozilla%252F5.0%2520(Macintosh%253B%2520Intel%2520Mac%2520OS%2520X%252010_15_7)%2520AppleWebKit%252F537.36%2520(KHTML%252C%2520like%2520Gecko)%2520Chrome%252F122.0.0.0%2520Safari%252F537.36&c_nonce=1710832797479&c_timestamp=1710832797479&c_app_id=200000001&c_app_uid=144115262788474108&c_platform=0&c_token=&c_app_version=&platform=Web&app_id=200000001&c_lang=zh-CN"

	requestBody := []byte(`{"host_id":"8651634848540374108","role_id":2,"name":"haha","phone":"15669489962","email":"","area_code":"86"}`)

	// 创建一个新的请求对象
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "pgv_pvid=5818805490; ab_uuid=20e67858f47142468f5c14b81cd160a5; _ga_J5ZYJSX9HN=GS1.1.1650597486.1.0.1650597503.0; web_uid=389dd316-9ce4-4985-9ab1-3b347744420a; landing_url=https://pre.txmeeting.tencent.com/qrcode-login.html?redirect_type=2&redirect_link=events-lobby%2FE165865955832550195210041%3Flogin_type%3D0; landing_path=https://pre.txmeeting.tencent.com/qrcode-login.html; landing_referralurl=https://pre.txmeeting.tencent.com/events-lobby/E165865955832550195210041; landing_referraldomain=https://pre.txmeeting.tencent.com; _qimei_uuid42=181020e2c1310082ad7d9f3cf146dbb693b7d0afda; _qimei_q36=; _qimei_h38=7e79c3e7712dd761818677250300000f917904; _gcl_au=1.1.2122262273.1706153534; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%22100024540397%22%2C%22first_id%22%3A%2218c6b6132dab4c-0efb27669c56ce8-1f525637-3686400-18c6b6132dbdd6%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfY29va2llX2lkIjoiMThjNmI2MTMyZGFiNGMtMGVmYjI3NjY5YzU2Y2U4LTFmNTI1NjM3LTM2ODY0MDAtMThjNmI2MTMyZGJkZDYiLCIkaWRlbnRpdHlfbG9naW5faWQiOiIxMDAwMjQ1NDAzOTcifQ%3D%3D%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%22100024540397%22%7D%2C%22%24device_id%22%3A%2218c6b6132dab4c-0efb27669c56ce8-1f525637-3686400-18c6b6132dbdd6%22%7D; pgv_info=ssid=s119735884; _gid=GA1.2.1692842897.1710830199; _qimei_fingerprint=1f52e8a2f7c812be4941e1f63b7c116e; we_meet_token=eJxUkF1vmzAUhv*Lb7Os-gq2I*2iE1Fo1gitYcvHDXLghHgQcDA0pNP**xRK1ZY7nud9j338F0WPq6*mdI0uE4hNiqZogr70VFsbtz0hnBMyoR4VUnLBCZZDxKV5fIv1KYpfPzJI6KypIdaHBmo0RR7mEuPBPUPtTFX2LcoIJuLWHOS*daYE59AUXWA-wMac4HYVQbBkmCpv4EWVmTJurvZm7bEqIbbuMsgUns3bWgOqIXs9WL5NcCZDU3Sad-YXlrsw*WGjs2F*twy-B5tkXTF6p6*yOOXLn611i4XZzPTLwS3lmkUBz66riWj80UsqOrusqMmL6DF80rOOzDy1n8-Jlu7c7-v1cbe3NKBUhosR-6MAm6fNA6iiLNpDe17fi1FeFXUebkbbmmcqupvQB94c05X2L16gtr7waR76AZt9*-D6H1c7mxMY5r0DW1dpmzSfQq7-0YoyCoSMGVA85qBgrKVWY5ImnMFBcOrh9wKr*wH--gcAAP--prGkJA__; corp_id=200000001; app_uid=144115262788474108; token_expire_time=1711435096; ACTIVITY_TICKET=eJxMUEtz2jAY-C%2A6ttPRwyDhmR4oqXnEIQVDgV40Aku2Yiz5IRuSTv97J47JoNvuamf3279gE0bfxOlkG%2AO4ey0k8AH42rE6lsZppWUFfIDhx0MceR5CAzzElDGPegiy-r8oCh1z4TipYuAD2NN1nPFOAj5AEFE6GhHSa-Ja6EpyoVwXghi8uVpZ1dqaLhkTSPHgPb0Xnc7feyKKICMQj4Y9f7aJNrcritQayYv6cuuhE%2ACDjFYkmh7L7X5Hqi-FDL1sDoPjmIqfT9eH-Md8dLyO3aNYZdNDsFa0WeM5zM%2AHVagedNics3ONBh7c1c-7Uiy13NrJWj0n47ZdDudQRZOMLZ7Sizr%2AthvmLfDkFw3NIS3-aLMohxEJ63xx2gaevswe0XSXrKBQsrH7NGokSoKXIH9Ll6-BDG9Z-b0vXsnkY4kex7LVJ8m7Qe82vsfOZtLwT2PyBv79DwAA--%2AZfpT4; _ga=GA1.2.204146153.1705055152; _ga_6WSZ0YS5ZQ=GS1.1.1710830297.9.1.1710830316.0.0.0")

	// 发送请求并获取响应
	resp, err := client.Do(req)
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
