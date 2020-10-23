package lib_core

import (
	"os"
	"pwdgenerator/gologger"
)



func Start(option *Options){
	gologger.Infof(" 当前关键词：%s\n",option.domain_key)
	gologger.Infof(" 当前level：%d\n",option.level)
	//存储的文件路径
	path := GetPath()
	//文件名加路径
	file_path := path+option.file_name
	//print(file_path+"\n")
	file, err := os.Create(file_path)
	if err != nil {
		gologger.Errorf("err=%v\n",err)
	}
	defer file.Close()
	gologger.Infof(" 根据规则生成对应密码\n")
	if AddRulePass(file,option.domain_key){
		gologger.Infof(" 根据规则生成对应密码成功\n")
	}
	gologger.Infof(" 加入一点常见弱口令\n")
	if AddCommonPass(file){
		gologger.Infof(" 增加常见弱口令成功\n")
	}
	gologger.Infof(" 加入一点键盘顺序弱口令\n")
	if AddKeyboardPass(file){
		gologger.Infof(" 增加键盘顺序弱口令成功\n")
	}
	gologger.Infof(" 对结果进行去重、复杂度检测、长度筛选\n")
	if Uniq(file_path,option.file_name,option){
		gologger.Infof(" 去重、复杂度检测、长度筛选成功\n")
	}

}
