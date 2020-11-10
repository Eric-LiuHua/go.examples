package filex

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//错误处理
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

//判断文件是否存在
func FileExist(filename string) bool {
	var Exist bool = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		Exist = false
	}
	return Exist
}

//缓存方式写入
func WriteFile4(filepath string, filename string, data []byte) {
	file := fmt.Sprintf("%s%s", filepath, filename)
	f, err := os.Create(file)
	CheckError(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	len, err := w.Write(data)
	fmt.Printf("file:%v Write %d  byte\n", file, len)

	w.Flush()

}

//File(Write,WriteString)
func WriteFile3(filepath string, filename string, data []byte) {
	file := fmt.Sprintf("%s%s", filepath, filename)

	f, err := os.Create(file)
	CheckError(err)
	defer f.Close()

	len, err := f.Write(data)
	fmt.Printf("file:%v Write %d  byte\n", file, len)
	f.Sync()
}

//ioutil.WriteFile
func WriteFile2(filepath string, filename string, data []byte) {
	file := fmt.Sprintf("%s%s", filepath, filename)
	err := ioutil.WriteFile(file, data, 0666)
	fmt.Printf("file:%v Write %d  byte\n", file, len(data))
	CheckError(err)
}

//io.WriteString
func WriteFile1(filepath string, filename string, data []byte) {
	var f *os.File
	var err error
	file := fmt.Sprintf("%s%s", filepath, filename)
	if FileExist(file) {
		//OpenFile,Open()的区别，Open()只能用于读取文件。
		f, err = os.OpenFile(file, os.O_APPEND, 0666)
	} else {
		f, err = os.Create(file)
	}
	CheckError(err)

	len, err := io.WriteString(f, string(data[:]))
	fmt.Printf("file:%v Write %d  byte\n", file, len)
}
