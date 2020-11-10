package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	lenght  int
	charset string
	crypto  bool
)

const (
	CHARS   string = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	NUMBERS string = "0123456789"
	SPECS   string = "@#$%&*?"
)

//参数帮助
func InitArgs() {
	flag.IntVar(&lenght, "len", 16, "-len 密码长度！")
	flag.StringVar(&charset, "type", "", `-type 指定密码生成的字符集。
num:数字[0-9]
char:英文字母[a-z A-Z]
mix:数字字母混合
advance:mix 基础数增加特殊字符
`)

	//解析，
	flag.Parse()
}

//生成随机字符
func GeneratePassword() (passwordStr string, cryptoStr string) {
	var password []byte = make([]byte, lenght, lenght)
	var source string
	switch charset {
	case "num":
		source = NUMBERS
	case "char":
		source = CHARS
	case "mix":
		source = fmt.Sprintf("%s%s", NUMBERS, CHARS)
	case "advance":
		source = SPECS
	default:
		source = fmt.Sprintf("%s%s%s", NUMBERS, CHARS, SPECS)
	}
	fmt.Println("source:", source)
	for i := 0; i < lenght; i++ {
		index := rand.Intn(len(source))
		password[i] = source[index]
	}
	passwordStr = string(password)
	//md5加密
	cryptoStr = fmt.Sprintf("%x", md5.Sum(password))
	return
}

func main() {
	//初始化种子
	rand.Seed(time.Now().Unix())
	InitArgs()
	fmt.Printf("lenght:%d,charset:%s \n", lenght, charset)

	passwordStr, cryptoStr := GeneratePassword()
	fmt.Printf("passwordStr: %s \ncryptoStr: %s \n", passwordStr, cryptoStr)
}
