package print

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"
	"unicode/utf8"
)

func TestCal(t *testing.T) {

	var result uint64
	for i := 80; i > 60; i-- {
		//result = result * i
	}

	fmt.Println(result)
}

func TestPrintBase64(t *testing.T) {

	meetingName := ""

	title := "d2ViaW5hcg=="

	meetingNameByte, _ := base64.StdEncoding.DecodeString(title)
	if len(meetingNameByte) > 0 {
		meetingName = string(meetingNameByte)
	}
	if meetingName == "" {
		meetingName = title
	}

	fmt.Println(meetingName)
}

type activityPeopleInfo struct {
	OuterId     string
	EventId     string
	SchedulerId uint64
	Uid         string
	Phone       string
}

func TestPrintPointer(t *testing.T) {
	a := &activityPeopleInfo{
		EventId: "1",
	}
	log.Printf("%+v", a)
	b := []*activityPeopleInfo{a}
	result, e := json.Marshal(b)
	log.Printf("%s,%+v", result, e)
}

func TestPrintMapPointer(t *testing.T) {
	audienceWithoutTicketMap := make(map[string]*activityPeopleInfo, 0)

	audienceWithoutTicketMap["1"] = &activityPeopleInfo{
		EventId: "1",
	}
	audienceWithoutTicketMap["2"] = &activityPeopleInfo{
		EventId: "2",
	}
	audienceWithoutTicketMap["3"] = &activityPeopleInfo{
		EventId: "3",
	}

	log.Printf("ressult: %+v", audienceWithoutTicketMap)
}

func TestPrintArrayPointer(t *testing.T) {
	array := make([]*activityPeopleInfo, 0)
	array = append(array, &activityPeopleInfo{
		EventId: "1",
	})
	array = append(array, &activityPeopleInfo{
		EventId: "2",
	})
	array = append(array, &activityPeopleInfo{
		EventId: "3",
	})
	log.Printf("ressult: %+v", array)
	fmt.Println(array)
}

func TestPrintString(t *testing.T) {
	AppletUrlSchemeRequestBody := `{"jump_wxa":{"path": "%s","query": "%s",
		"env_version": "release"},"is_expire":true,"expire_type":1,"expire_interval":30}`

	fmt.Println(fmt.Sprintf(AppletUrlSchemeRequestBody, "12345678", ""))
}

func TestPrintSecond(t *testing.T) {
	fmt.Println(time.Now().Unix() - 15*3600)
	fmt.Println(1690794000 > time.Now().Unix()-15*3600)
}

func TestPrintRune(t *testing.T) {
	meetingTitle := "a阮金"
	fmt.Println(meetingTitle[:2])
	fmt.Println(string([]rune(meetingTitle)[:2]))
	meetingTitle = string([]rune(meetingTitle)[:1]) + "***" + string([]rune(meetingTitle)[utf8.
		RuneCountInString(meetingTitle)-1:])
	fmt.Println(meetingTitle)

}

func TestDivide(t *testing.T) {
	cleanWhiteLogCount := 2000
	batchCount := cleanWhiteLogCount / 1000
	if cleanWhiteLogCount%1000 > 0 {
		batchCount++
	}
	fmt.Println(batchCount)
}
