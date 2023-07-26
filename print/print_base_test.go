package print

import (
	"fmt"
	"log"
	"testing"
	"time"
)

type activityPeopleInfo struct {
	OuterId string
	EventId string
	SchedulerId uint64
	Uid		string
	Phone	string
}

func TestPrintMapPointer(t *testing.T) {
	audienceWithoutTicketMap:= make(map[string]*activityPeopleInfo, 0)

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
	fmt.Println(time.Now().Unix() - 15 * 3600)
	fmt.Println(1690794000 > time.Now().Unix() - 15 * 3600)
}