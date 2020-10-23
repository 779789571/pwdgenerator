package lib_core

import (
	"strings"
)

func Filter_pass(password string, option *Options) bool {
	var kind int
	password = strings.Replace(password, " ", "", -1)
	password = strings.Replace(password, "\n", "", -1)
	//gologger.Infof(password)
	if len(password) < option.min_len {
		return false
	}
	if len(password) > option.max_len {
		return false
	}
	kind = 1
	if CheckSpecialChar(password) {
		kind += 1
	}
	if CheckNumber(password) {
		kind += 1
	}
	if CheckUpperLetter(password) {
		kind += 1
	}
	if CheckLowerLetter(password) {
		kind += 1
	}
	if kind <= option.level {
		//gologger.Infof("didn't pass kind : %d", kind)
		//gologger.Infof("didn't pass: %s", password)
		return false

	}
	//gologger.Infof("pass: %s", password)
	return true
}

//参考https://www.coder.work/article/200560
func StringToInts(s string) (intSlice []int) {
	intSlice = make([]int, len(s))
	for i, _ := range s {
		intSlice[i] = int(s[i])
	}
	return
}

//参考bit4woo表哥的实现方式，采用ascii判断
func CheckSpecialChar(password string) bool {
	pass_to_ascii := StringToInts(password)
	for _, v := range pass_to_ascii {
		if 32 <= v && v <= 47 || 58 <= v && v <= 64 || 91 <= v && v <= 96 || 123 <= v && v <= 126 {
			return true
		}
	}
	return false
}
func CheckNumber(password string) bool {
	pass_to_ascii := StringToInts(password)
	for _, v := range pass_to_ascii {
		if 48 <= v && v <= 57 {
			return true
		}
	}
	return false
}
func CheckUpperLetter(password string) bool {
	pass_to_ascii := StringToInts(password)
	for _, v := range pass_to_ascii {
		if 65 <= v && v <= 90 {
			return true
		}
	}
	return false
}
func CheckLowerLetter(password string) bool {
	pass_to_ascii := StringToInts(password)
	for _, v := range pass_to_ascii {
		if 97 <= v && v <= 122 {
			return true
		}
	}
	return false
}
