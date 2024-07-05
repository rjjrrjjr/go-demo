package urls

import (
	"fmt"
	"net/url"
	"testing"
)

func TestParse(t *testing.T) {

	fmt.Println("-----------")
	_, err := url.Parse("asfas")
	fmt.Println(err, err == nil)

	_, err = url.Parse("http:/s/a")
	fmt.Println(err, err == nil)

	fmt.Println("-----------")
}

func TestOneid(t *testing.T) {
	webHandlePageUrl, err := url.QueryUnescape("https%3A%2F%2Fe10fcn8smo0800.account.tencent.com%2F%23%2Fcidp%2Fredirect_identity")
	if err != nil {
		fmt.Println(err)
	}
	parsedWebHandlePageUrl, err := url.Parse(webHandlePageUrl)
	if err != nil {
		fmt.Println(err)
	}
	q := parsedWebHandlePageUrl.Query()
	parsedWebHandlePageUrl.RawQuery = q.Encode()
	fmt.Println(parsedWebHandlePageUrl.String())
}

func TestUrlPayload(t *testing.T) {
	parseUrl, err := url.Parse("www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	values := url.Values{}
	values.Add("a", "hahaah")
	values.Add("b", "heheeh")

	parseUrl.RawQuery = values.Encode()
	fmt.Println(parseUrl.String())
}
