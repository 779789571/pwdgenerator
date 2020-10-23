package lib_core

import (
	"flag"
	"strings"
	"os"
	"pwdgenerator/gologger"
)

type Options struct {
	domain_key string
	min_len    int
	max_len    int
	level      int
	file_name  string
}

func ParseOptions() *Options {
	options := &Options{}
	domain_key := flag.String("d", "", "域名关键字")
	min_len := flag.Int("min", 0, "密码最小长度")
	max_len := flag.Int("max", 30, "密码最大长度")
	level := flag.Int("l", 1, "密码复杂级别，1-5，数值越高，复杂度要求越高,字典数量越少")
	//uniq := flag.String("q","","对传入的文件进行去重")
	flag.StringVar(&options.file_name, "o", "", "输出的文件名，默认为txt格式")
	flag.Parse()

	if *domain_key != "" {
		if CheckDomainKey(*domain_key) {
			options.domain_key = *domain_key
		} else {
			gologger.Errorf(" 输入的关键词有问题")
			os.Exit(-1)
		}
	}
	if *min_len >= 0 {
		options.min_len = *min_len
	} else {
		gologger.Errorf(" 设定的密码最小长度有问题")
		os.Exit(-1)
	}
	if *max_len > 0 {
		options.max_len = *max_len
	} else {
		gologger.Errorf(" 设定的密码最大长度有问题")
		os.Exit(-1)
	}
	if *min_len > *max_len {
		gologger.Errorf("密码最大长度需大于最小长度")
		os.Exit(-1)
	}
	if *level > 0 && *level <= 5 {
		options.level = *level
	} else {
		gologger.Errorf(" 密码复杂度 level 仅支持1-5之间选择")
		os.Exit(-1)
	}
	if options.file_name == "" {
		options.file_name = GetFileName(options.domain_key)
	}else {
		if strings.Contains(options.file_name,".txt") == false{
			options.file_name = options.file_name+".txt"
			gologger.Infof("%s",options.file_name)
		}
	}

	//print(options)
	return options
}
