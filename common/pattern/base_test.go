package pattern

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestEidHost(t *testing.T) {
	fmt.Println(regexp.MustCompile(`https://.*account.tencent.com/`).MatchString("https://account.tencent.com/"))
	fmt.Println(regexp.MustCompile(`https://.*account.tencentcs.com/`).MatchString("https://eid-25.account.tencentcs." +
		"com/"))
	fmt.Println(regexp.MustCompile(`https://.*account.tencentcs.com/`).MatchString("https://ci-746.account.tencentcs." +
		"com/"))
}

func TestPattern(t *testing.T) {

	check(".as_1Da")
	check("asd_3Aa.")
	check("aCs_2d..asd")
	check("哈4哈_B")
	check("asda.asD_5da.as2d.saSd.a2sd")
	check("aC.....s2d..a_sd")
	check("as231dS42AC31d31_2a2312s12231")
	check("as231dS42#AC31d31_2a2312s12231")
	check("as231dS42@AC31d312a2_312s12231")
	check("as231dS42.AC31d312a2312_s12231")

}

func check(param string) {
	if !regexp.MustCompile(`^[a-zA-Z0-9._]*$`).MatchString(param) {
		fmt.Println(param, "not a-zA-Z0-9.", isValidAlphanumericOrDot(param))
		return
	}
	if strings.HasPrefix(param, ".") {
		// 不允许以.开头，不允许以.结尾，不允许.连续
		fmt.Println(param, "HasPrefix .", isValidAlphanumericOrDot(param))
		return
	}

	if strings.HasSuffix(param, ".") {
		fmt.Println(param, "HasSuffix .", isValidAlphanumericOrDot(param))
		return
	}

	if regexp.MustCompile(`^.*\.\..*$`).MatchString(param) {
		fmt.Println(param, "has ..", isValidAlphanumericOrDot(param))
		return
	}

	fmt.Println(param, "is valid", isValidAlphanumericOrDot(param))

}

// 不允许以.开头，不允许以.结尾，不允许.连续
func isValidAlphanumericOrDot(param string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._]*$`).MatchString(param) &&
		!strings.HasPrefix(param, ".") &&
		!strings.HasSuffix(param, ".") &&
		!regexp.MustCompile(`^.*\.\..*$`).MatchString(param)
}
