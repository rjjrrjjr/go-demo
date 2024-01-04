package crypt

import (
	"crypto/md5"
	"fmt"
	"git.woa.com/fire/firedb/mode"
	"git.woa.com/fire/firedb/strategy"
	"testing"
)

func TestMd5(t *testing.T) {
	str := "hello world"
	data := []byte(str)
	md5Sum := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", md5Sum)
	fmt.Println(md5Str)
}

func TestMd52(t *testing.T) {
	fmt.Println(getMeetingSignature("12369175966676674434"))
}

func getMeetingSignature(meetingId string) string {
	data := []byte(meetingId + meetingId[:6] + "@#$")
	md5Sum := md5.Sum(data)
	return fmt.Sprintf("%x", md5Sum)
}

func TestAESS(t *testing.T) {
	aesKey := "a2dnTvwFe3gJbskt"

	cs := strategy.NewAesCryptoStrategy(aesKey)

	fmt.Println(cs.Decrypt(mode.PrefixCrypto, "{cpt}MTMzMTk5OTk3NjUFBQUFBZkkHbObkiUsxp9SzS3nDPg="))

}

func TestAES(t *testing.T) {
	// aesKey := "a2dnTvwFe3gJbskt"
	aesKey := "a2dnTvwFe3gJbskt"

	cs := strategy.NewAesCryptoStrategy(aesKey)

	//encrypt := cs.Encrypt(mode.PrefixCrypto, "18710as2861871001286187100128618@qw012861871001q861871001286qqqx")
	//fmt.Println(encrypt)
	//
	//encrypt2 := cs.Encrypt(mode.PrefixCrypto, "18710012863")
	//fmt.Println(encrypt2)

	phoneArr := []string{
		"{cpt}MTczNzQyOTg3OTMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU4NjExMjE5MjQFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgzODk3NzEzNDkFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMzMDcwNzA2ODIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",

		"{cpt}MTc4NTM0MzE3MjAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgwMjU0NTI3MjMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM5MjM3ODQxNDEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg1ODExNzE0MzEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",

		"{cpt}MTMzMTg3MTQ5MTMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTc2MTAwODcxODIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU3Mzc4NzI3NjMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgwMDE5OTQ3NTMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",

		"{cpt}MzQ3NDMzMTk1MQYGBgYGBpkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMwOTU3MTE3MTEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM1NjQxODIxOTcFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU2ODM5NjAwMDUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",

		"{cpt}MTkzNzY4MDE5MjMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMzOTY1Mjc4MjcFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTUxNzUyNzY1NjUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTc4MzkzNjA4ODgFBQUFBZkkHbObkiUsxp9SzS3nDPg=",

		"{cpt}MTc1NjE0Mjc4MjUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMzODkwNzY1NzgFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMxNjk4MTQ5OTMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTc4NTgxNTE4NTEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",

		"{cpt}MTgwOTY2MTE3NDIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTk1MzQzMDE5ODQFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM3OTA5NTc2OTYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg3MTgzMjQ5MTkFBQUFBZkkHbObkiUsxp9SzS3nDPg=",

		"{cpt}MTMxNzc2MTEyMjgFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTUxMTI5OTE5NTkFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU3MTQxMTg2NjYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}NDMwODYwOTIyBwcHBwcHB5kkHbObkiUsxp9SzS3nDPg=",

		"{cpt}MTU2MjcyMjI3NjgFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTUyMjc5NDUyNzEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU2ODU1ODE2MDMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU3NjAzNzUyODYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU2Mjc3NjMxNDMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM4OTA4NTI0MzIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg0NzY2NTg1MjYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg3NjYxODU4NjIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg1MTM3MjA1MjIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTc1MTMzNzUyMjEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU5MjA1NTE3MjcFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMzODI1MDczMzYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU4MDE5ODc4OTkFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTUzMTM3ODI4NTYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTk2MjEzMTQ2MDYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM3MDE5NTQwNzgFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMyNDMzNjA2MzAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM3NzMyMjY2NjYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg2Mjk0ODI5MzAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU1OTgxMTY4MTEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}NDE0ODM5NTgwBwcHBwcHB5kkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM3Mjg5NDMzMTEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTczNjI1ODgzMTEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM5OTU2ODYyMjUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU2MDYyODE4ODQFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM5MjEwNTE1MjUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgyNjA3MjgwMTQFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMyOTY1NDY2ODYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg3ODAxNzkxMzIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM3ODE5Njc0OTYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU4NjYzNDIyMTAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgxNzg1ODU0NDMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg4MjM3ODEzODMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM2NjE3NTE1MDcFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM5NzkzNjk4MDIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTk5MzI3Njg1NzYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU1Mjc4Mjc2MzMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTk1Mjg4NTk0MDkFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
	}

	phoneMap := map[string]struct{}{"15627222768": {}, "17853431720": {}, "13095711711": {}, "13790957696": {},
		"15714118666": {}, "13169814993": {}, "18096611742": {}}
	for index, s := range phoneArr {
		decrypt, _ := cs.Decrypt(mode.PrefixCrypto, s)
		if _, exist := phoneMap[decrypt]; exist {
			continue
		}
		fmt.Println(index+1, decrypt)
	}

	fmt.Println("------")

}

func TestA(t *testing.T) {

	phoneArrayA := []string{
		"{cpt}MTczNzQyOTg3OTMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU4NjExMjE5MjQFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgzODk3NzEzNDkFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMzMDcwNzA2ODIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgwMjU0NTI3MjMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM5MjM3ODQxNDEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMzMTg3MTQ5MTMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTc2MTAwODcxODIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU3Mzc4NzI3NjMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgwMDE5OTQ3NTMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM1NjQxODIxOTcFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU2ODM5NjAwMDUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTkzNzY4MDE5MjMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMzOTY1Mjc4MjcFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTUxNzUyNzY1NjUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTc1NjE0Mjc4MjUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMzODkwNzY1NzgFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTc4NTgxNTE4NTEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTk1MzQzMDE5ODQFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg3MTgzMjQ5MTkFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMxNzc2MTEyMjgFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTUxMTI5OTE5NDkFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTUyMjc5NDUyNzEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU2ODU1ODE2MDMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU3NjAzNzUyODYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU2Mjc3NjMxNDMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM4OTA4NTI0MzIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg0NzY2NTg1MjYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg3NjYxODU4NjIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg1MTM3MjA1MjIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTc1MTMzNzUyMjEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU5MjA1NTE3MjcFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMzODI1MDczMzYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTk2MjEzMTQ2MDYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM3MDE5NTQwNzgFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMyNDMzNjA2MzAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM3NzMyMjY2NjYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg2Mjk0ODI5MzAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU1OTgxMTY4MTEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM3Mjg5NDMzMTEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTczNjI1ODgzMTEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM5OTU2ODYyMjUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU2MDYyODE4ODQFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM5MjEwNTE1MjUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgyNjA3MjgwMTQFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMyOTY1NDY2ODYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU4NjYzNDIyMTAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgxNzg1ODU0NDMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg4MjM3ODEzODMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM2NjE3NTE1MDcFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM5NzkzNjk4MDIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTk5MzI3Njg1NzYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTk5ODMxMDg5MjYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgzODEwNTQ5MTkFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM1MjE4NTUyODAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTUxNDI1MDc0NTkFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMzMTk5OTk3NjUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMzNzY4NzM3NzgFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU5NzA2ODEwOTAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgzNjgwMjI1MTAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTUzODk4MDM2MTYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}NDMwODYwOTIyBwcHBwcHB5kkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTUyMDczMjA4MDUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}NDE0ODM5NTgwBwcHBwcHB5kkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM3NjcwODk1NDAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgwMTg3NjY2NjUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
	}

	phoneArrB := []string{
		"{cpt}MTczNzQyOTg3OTMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU4NjExMjE5MjQFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgzODk3NzEzNDkFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMzMDcwNzA2ODIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",

		"{cpt}MTc4NTM0MzE3MjAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgwMjU0NTI3MjMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM5MjM3ODQxNDEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg1ODExNzE0MzEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",

		"{cpt}MTMzMTg3MTQ5MTMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTc2MTAwODcxODIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU3Mzc4NzI3NjMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgwMDE5OTQ3NTMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",

		"{cpt}MzQ3NDMzMTk1MQYGBgYGBpkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMwOTU3MTE3MTEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM1NjQxODIxOTcFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU2ODM5NjAwMDUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",

		"{cpt}MTkzNzY4MDE5MjMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMzOTY1Mjc4MjcFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTUxNzUyNzY1NjUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTc4MzkzNjA4ODgFBQUFBZkkHbObkiUsxp9SzS3nDPg=",

		"{cpt}MTc1NjE0Mjc4MjUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMzODkwNzY1NzgFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMxNjk4MTQ5OTMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTc4NTgxNTE4NTEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",

		"{cpt}MTgwOTY2MTE3NDIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTk1MzQzMDE5ODQFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM3OTA5NTc2OTYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg3MTgzMjQ5MTkFBQUFBZkkHbObkiUsxp9SzS3nDPg=",

		"{cpt}MTMxNzc2MTEyMjgFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTUxMTI5OTE5NTkFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU3MTQxMTg2NjYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}NDMwODYwOTIyBwcHBwcHB5kkHbObkiUsxp9SzS3nDPg=",

		"{cpt}MTU2MjcyMjI3NjgFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTUyMjc5NDUyNzEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU2ODU1ODE2MDMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU3NjAzNzUyODYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU2Mjc3NjMxNDMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM4OTA4NTI0MzIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg0NzY2NTg1MjYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg3NjYxODU4NjIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg1MTM3MjA1MjIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTc1MTMzNzUyMjEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU5MjA1NTE3MjcFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMzODI1MDczMzYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU4MDE5ODc4OTkFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTUzMTM3ODI4NTYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTk2MjEzMTQ2MDYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM3MDE5NTQwNzgFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMyNDMzNjA2MzAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM3NzMyMjY2NjYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg2Mjk0ODI5MzAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU1OTgxMTY4MTEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}NDE0ODM5NTgwBwcHBwcHB5kkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM3Mjg5NDMzMTEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTczNjI1ODgzMTEFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM5OTU2ODYyMjUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU2MDYyODE4ODQFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM5MjEwNTE1MjUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgyNjA3MjgwMTQFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMyOTY1NDY2ODYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg3ODAxNzkxMzIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM3ODE5Njc0OTYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU4NjYzNDIyMTAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgxNzg1ODU0NDMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTg4MjM3ODEzODMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM2NjE3NTE1MDcFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM5NzkzNjk4MDIFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTk5MzI3Njg1NzYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU1Mjc4Mjc2MzMFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTk1Mjg4NTk0MDkFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
	}

	phoneArrayBMap := make(map[string]struct{}, 0)

	for _, s := range phoneArrB {
		phoneArrayBMap[s] = struct{}{}
	}

	for _, s := range phoneArrayA {
		if _, exist := phoneArrayBMap[s]; !exist {
			fmt.Println(s)
		}
	}
}

func TestB(t *testing.T) {
	phoneArray := []string{
		"{cpt}MTUxMTI5OTE5NDkFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTk5ODMxMDg5MjYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgzODEwNTQ5MTkFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM1MjE4NTUyODAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTUxNDI1MDc0NTkFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMzMTk5OTk3NjUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTMzNzY4NzM3NzgFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTU5NzA2ODEwOTAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgzNjgwMjI1MTAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTUzODk4MDM2MTYFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTUyMDczMjA4MDUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTM3NjcwODk1NDAFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
		"{cpt}MTgwMTg3NjY2NjUFBQUFBZkkHbObkiUsxp9SzS3nDPg=",
	}

	aesKey := "a2dnTvwFe3gJbskt"

	cs := strategy.NewAesCryptoStrategy(aesKey)

	for _, s := range phoneArray {
		decrypt, _ := cs.Decrypt(mode.PrefixCrypto, s)

		fmt.Println(decrypt)
	}
}
