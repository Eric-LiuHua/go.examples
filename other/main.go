package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//练习，打印变量的地址,int 是值传递，这个a的地址跟调用的a已经不一样了。
func PrintAddr(a int) {
	fmt.Printf("PrintAddr a.addr=%p \n", &a)
}

//练习，修改地址中的值
func ModifyAddrValue(v *int, tmp int) {
	fmt.Printf("ModifyAddrValue new value=%d , addr=%p \n", tmp, v)
	*v = tmp
}

//统计单词数量
func CoutWord(str string) map[string]int {
	res := make(map[string]int)

	if len(str) > 0 {
		slice := strings.Split(str, " ")
		for _, v := range slice {
			_, isexist := res[v]
			if isexist {
				res[v] = res[v] + 1
			} else {
				res[v] = 1
			}
		}
	}

	return res
}

//interface[] 可以存任意类型
func studentStore() {
	rand.Seed(time.Now().UnixNano())
	store := make(map[int]map[string]interface{}, 18)
	for i := 0; i < 100; i++ {
		_, isexist := store[i]
		if !isexist {
			store[i] = make(map[string]interface{}, 10)
		}
		store[i]["id"] = i
		store[i]["name"] = fmt.Sprintf("student%d", i)
		store[i]["age"] = rand.Intn(25)
		store[i]["score"] = rand.Float32() * 100
	}
	fmt.Println("*************** studentStore *****************")
	for key, student := range store {
		fmt.Printf("key:%v,student:%v \n", key, student)
	}

}

func main() {

	a := 983
	fmt.Printf("init a =%d \n", a)
	fmt.Printf("init a.addr=%p \n", &a)

	PrintAddr(a)
	//&a 获取地址
	ModifyAddrValue(&a, 12312)

	fmt.Printf("after modify a =%d \n", a)

	str := "aa bb cc dd ss dd cc cc aa dd bb bb bb bb bb"
	var res map[string]int = CoutWord(str)

	for key, count := range res {
		fmt.Printf("key:%s,count:%v \n", key, count)
	}

	studentStore()
}
