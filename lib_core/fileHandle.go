package lib_core

import (
	"bufio"
	"io"
	"os"
	"pwdgenerator/gologger"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}
//加载字典
func FileLoad(file_path string , bufSize int) string{
	var res string
	var chunk []byte
	file,err := os.OpenFile(file_path,os.O_RDONLY,0666)
	check(err)
	defer file.Close()
	r := bufio.NewReader(file)
	b := make([]byte, bufSize)
	for {
		n, err := r.Read(b)
		chunk = append(chunk,b[:n]...)
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			} else {
				gologger.Errorf("%s",err)
				os.Exit(-1)
			}
		}
		//res = string(b)


	}
	res = string(chunk)
	if res != ""{
		return res
	}
	return ""

}

func RemoveFile() {
	err := os.Remove("test.txt")
	if err != nil {
		gologger.Errorf("%s",err)
		//fmt.Println(err)
	}
}