package lib_core

import (
	"bufio"
	"io"
	"os"
	"pwdgenerator/gologger"
	"strconv"
	"strings"
	"time"
)

func GetFileName(filename string) string {
	timeUnix := time.Now().Unix()
	str := filename + "_" + strconv.FormatInt(timeUnix, 10) + ".txt"
	gologger.Infof(" filename未设定，生成名为：%s\n", str)
	return str
}
func CheckDomainKey(domain_key string) bool {
	//后期写检查
	return true
}
func GetPath() string {
	//获取生成文件路径,后期改成绝对路径
	path := "./results/"
	return path
}

//合并数组
func MergeSlice(s1 []string, s2 []string) []string {
	for _, value := range s2 {
		s1 = append(s1, value)
	}
	return s1
}

//字典去重，复杂度、长度筛选、
func Uniq(file_path string, filename string, option *Options) bool {
	file, err := os.OpenFile(file_path, os.O_RDONLY, 0666)
	check(err)
	var chunk []byte
	r := bufio.NewReader(file)
	b := make([]byte, 1024)
	for {
		n, err := r.Read(b)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		chunk = append(chunk, b[:n]...)
	}
	slice_list := strings.Split(string(chunk), "\n")
	//for _,p := range slice_list{
	//	gologger.Infof("pass:%s ",p)
	//}
	result := removeDuplicateElement(slice_list) //去重后保存的切片
	path := GetPath()
	filePathNew := path + "筛选完毕" + option.file_name
	file_new, err_new := os.Create(filePathNew)
	if err_new != nil {
		gologger.Errorf("%s", err_new)
		return false
	}
	defer file_new.Close()
	for _, password := range result {
		if Filter_pass(password, option) {
			password = password + "\n"
			file_new.Write([]byte(password))
		}

	}
	return true
}

//参考http://www.36nu.com/post/329 和 https://www.yuque.com/fz420/golang/ky17s2
func removeDuplicateElement(addrs []string) []string {
	result := make([]string, 0, len(addrs))
	temp := map[string]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
