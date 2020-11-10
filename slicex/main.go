package main

import (
	"fmt"
	"time"
)

//全局切片
var slicc []int

//对切片 append
func appendValue(x int) {
	slicc = append(slicc, x)
}

func main() {
	for i := 0; i < 10000; i++ {
		go appendValue(i)
	}

	time.Sleep(time.Second)
	for i, v := range slicc {
		fmt.Printf("i:%d ,v:%d \n", i, v)
	}
}
