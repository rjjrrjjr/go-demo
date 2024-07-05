package time

import (
	"fmt"
	"testing"
	"time"
)

func TestInt642Time(t *testing.T) {
	timeZoneLocal, _ := time.LoadLocation("Local")
	fmt.Println(time.Unix(1705399354, 0).In(timeZoneLocal))
	fmt.Println(time.Unix(1705399354, 0))
}

func TestTimeConvert(t *testing.T) {
	fmt.Println(time.Now())

	fmt.Println(time.Now().Unix())      // 时间戳，秒
	fmt.Println(time.Now().UnixMilli()) // 时间戳，毫秒
	fmt.Println(time.Now().UnixMicro()) // 时间戳，微妙
	fmt.Println(time.Now().UnixNano())  // 时间戳，纳秒

	fmt.Println(time.Now().String())
	fmt.Println(time.Now().GoString())

	fmt.Println(time.Now().Date())
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Month())
	fmt.Println(int(time.Now().Month()))
	fmt.Println(time.Now().Month().String())
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Clock())
	fmt.Println(time.Now().Hour())
	fmt.Println(time.Now().Minute())
	fmt.Println(time.Now().Second())

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	timeZoneLocal, _ := time.LoadLocation("Local") // "Asia/Shanghai"
	timeZoneUTC, _ := time.LoadLocation("UTC")     // "Asia/Shanghai"

	fmt.Println(time.Unix(1687316362, 0).Format("2006-01-02 15:04:05"))
	fmt.Println(time.Unix(1687316362, 0).In(timeZoneUTC).Format("2006-01-02 15:04:05"))

	time7 := "2023-06-21 11:00:50"
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", time7, timeZoneLocal)
	fmt.Println(tt.String())

	tt, _ = time.Parse("2006-01-02 15:04:05", time7) // 默认会被转换成UTC time
	fmt.Println(tt.In(timeZoneLocal).String())

	year := 2023
	month := time.October
	day := 1

	date := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	fmt.Println(day, date)

}

func TestBatchConvert(t *testing.T) {
	for _, value := range []int64{1691893800, 1691896500, 1691893800, 1691896500, 1691893800, 1691896500,
		1691906400, 1691916300} {
		fmt.Println(time.Unix(value, 0).Format("2006-01-02 15:04:05"))
	}
}

func TestPrintForamt(t *testing.T) {
	fmt.Println(time.Now().Format("20060102"))
}

func TestParse(t *testing.T) {
	date, err := time.ParseInLocation("2006-01-02", "2023-10-1", time.Local)
	fmt.Println(date, err)
}
