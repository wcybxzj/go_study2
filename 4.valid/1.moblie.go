package main

import (
	"fmt"
	"regexp"
)

const (
	moblieRegular = "^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\\\d{8}$"
)

func validateMobile(mobileNum string) bool {
	reg := regexp.MustCompile(moblieRegular)
	return reg.MatchString(mobileNum)
}

func main() {
	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)
	s := []string{"18505921256", "13489594009", "12759029321", "18610755072"}
	for _, v := range s {
		fmt.Println(rgx.MatchString(v))
	}
}
