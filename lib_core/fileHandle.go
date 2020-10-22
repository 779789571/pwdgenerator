package lib_core

import (
	"bufio"
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
		_, err := r.Read(b)
		if err != nil {
			//gologger.Printf("Error reading file:", err)
			break
		}
		//res = string(b)
		chunk = append(chunk,b...)

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