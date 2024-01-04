package io

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"
)

func TestParseCleanFile(t *testing.T) {
	// 打开文件
	file, err := os.Open("2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var meetingNameReg = regexp.MustCompile("MN：(.+?)(；MT)")

	reStrList := []string{"(webinar_onthewayli).*", "(webinar_enroll)\\d+", "(ethanxxli).*", "(自动化.*\\d{4}-\\d{2}-\\d{2})", "(自动化.*?)\\d+", "(webinar)\\d+", "(onthewayli).*", "(测试会议)", "(特别关注webinar会议)", "(webinar测试)\\d+", "(query meeting item list test)", "(webinar会议会中详情查询)", "(会中详情测试)", "admin\\d+(预定的网络研讨会)", "(auto_Manager2预定的网络研讨会)", "(meeting waiting rooms test)", "(auto_Manager预定的网络研讨会)"}

	reList := make([]*regexp.Regexp, 0)

	for _, s := range reStrList {
		reList = append(reList, regexp.MustCompile(s))
	}

	// 创建一个带缓冲的读取器
	scanner := bufio.NewScanner(file)

	// 逐行读取文件内容
	for scanner.Scan() {
		line := scanner.Text()

		if meetingNameReg.MatchString(line) {
			meetNameMatchArr := meetingNameReg.FindStringSubmatch(line)
			match := false
			for _, re := range reList {
				if match = re.MatchString(meetNameMatchArr[1]); match {
					break
				}
			}
			if match {
				continue
			}
		}
		fmt.Println(line)
	}

	// 检查是否有错误发生
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
