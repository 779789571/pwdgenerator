package lib_core

import (
	"flag"
	"pwdgenerator/gologger"
)

type Options struct {
	domain_key string
	min_len    int64
	max_len    int64
	level      int64
	file_name  string
}

func ParseOptions() *Options {
	options := &Options{}
	domain_key := flag.String("d", "", "域名关键字")
	min_len := flag.Int64("min", 3, "密码最小长度")
	max_len := flag.Int64("max", 0, "密码最大长度")
	level := flag.Int64("l", 1, "密码复杂级别，1-5，数值越高，复杂度要求越高,字典数量越少")
	//uniq := flag.String("q","","对传入的文件进行去重")
	flag.StringVar(&options.file_name, "o", "", "输出的文件名，默认为txt格式")
	flag.Parse()


	if *domain_key != "" {
		if CheckDomainKey(*domain_key) {
			options.domain_key = *domain_key
		} else {
			gologger.Errorf("domain_key has some error")
		}
	}
	if *min_len > 0 {
		options.min_len = *min_len
	}
	if *max_len > 0 {
		options.max_len = *max_len
	}
	if *level > 0 && *level < 5 {
		options.level = *level
	}
	if options.file_name == "" {
		options.file_name = GetFileName(options.domain_key)

	}

	//print(options)
	return options
}


