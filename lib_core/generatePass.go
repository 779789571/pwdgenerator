package lib_core

import (
	"os"
	"pwdgenerator/gologger"
	"strings"
)

//处理常规密码生成
func AddCommonPass(file *os.File) bool {

	res := FileLoad(common_passFile, 1000000)
	//f ,err := os.OpenFile(commonpassFile,os.O_APPEND,0666)
	//check(err)
	//defer f.Close()
	if res != "" {
		file.Write([]byte(res))
		return true
	}
	//fmt.Println(res)
	return false
}

//处理定制化密码生成
func AddRulePass(file *os.File, key string) bool {

	var rules = []string{"key+special_letter+year", "key+special_letter+keyboard_walk", "key+keyboard_walk", "key+special_letter+common_pass", "common_pass+special_letter+key","key+common_pass", "key+special_letter+china_name", "china_name+special_letter+key", "key+special_letter+common_number", "key+common_number+special_letter"}

	for count, value := range rules {
		gologger.Infof(" 增加第%d种可能：%s\n", count+1, value)
		//用+号分割为多个切片
		var rule = strings.Split(value, "+")

		result_list := RuleGotPass(key, rule)

		if len(result_list) != 0 {
			for _, pass := range result_list {
				pass = pass + "\n"
				file.WriteString(pass)
			}
		}
	}
	return true
}

//通过解析规则，生成密码
func RuleGotPass(key string, rule []string) []string {
	var Rule_list = []string{}      // 返回的总切片
	var tmp_pass_list = []string{}  //存放逐一rule读取的切片
	var store_tmp_list = []string{} //存放一步步聚合的切片

	//var tmpKey = ""
	for _, k := range rule {
		//gologger.Infof(" rule:%s  \n", k)
		if CheckFormat(k) {
			//判断切片为空时
			if len(store_tmp_list) == 0 {
				tmp_pass_list = GetListByFormat(k, key)
				store_tmp_list = MergeSlice(store_tmp_list, tmp_pass_list)
				//清空读取rule的列表
				tmp_pass_list = []string{}
				//gologger.Infof("test:%s\n",Rule_list[0])
			} else {
				var tmp = []string{} //存放一步拼接完成的pass
				tmp_pass_list = GetListByFormat(k, key)
				for _, v := range store_tmp_list {
					for _, vv := range tmp_pass_list {
						pass := v + vv
						//gologger.Infof("test:"+pass)
						tmp = append(tmp, strings.Replace(pass, "\n", "", -1))
						//gologger.Infof("tmp:%s",tmp)
						//tmp = append(tmp,pass)
					}
				}
				store_tmp_list = tmp
				//for _,b := range store_tmp_list{
				//	gologger.Infof("test:%s",b)
				//}
				tmp = []string{}


			}

		} else {
			gologger.Errorf("rules had wrong fotmat: %s", k)
		}
	}
	Rule_list = store_tmp_list
	return Rule_list
}

//确认规则拼写是否正确
func CheckFormat(k string) bool {
	var Format = []string{"key", "special_letter", "year", "keyboard_walk", "common_pass", "china_name", "common_number"}
	for _, f := range Format {
		if k == f {
			return true
		}

	}
	return false
}

//确认类型，返回切片
func GetListByFormat(k string, key string) []string {
	var special_letter = []string{"!", "@", "#", "$", "%", "*"}
	var year = []string{"2015", "2016", "2017", "2018", "2019", "2020"}
	china_name := FileLoad(china_nameFile, 10000)
	keyboard_walk := FileLoad(keyboard_walkFile, 10000)
	common_number := FileLoad(common_numberFile, 10000)
	common_pass := FileLoad(common_passFile, 10000)
	var list = []string{}
	switch k {
	case "key":
		list = append(list, key)

	case "special_letter":
		list = MergeSlice(list, special_letter)
		return list
	case "year":
		list = MergeSlice(list, year)
		return list
	case "china_name":
		list = append(list, china_name)

	case "keyboard_walk":
		list = append(list, keyboard_walk)

	case "common_number":
		list = append(list, common_number)

	case "common_pass":
		list = append(list, common_pass)

	}

	list = strings.Split(list[0], "\n")

	return list
}
func AddKeyboardPass(file *os.File) bool{
	res := FileLoad(keyboard_pass, 1000000)
	//f ,err := os.OpenFile(commonpassFile,os.O_APPEND,0666)
	//check(err)
	//defer f.Close()
	if res != "" {
		file.Write([]byte(res))
		return true
	}
	//fmt.Println(res)
	return false
}