package main

import (
	"fmt"
	"sync"
	"time"
)

//go 中map 不是线程安全的。
func readMap(d map[int]string) {
	fmt.Println(d[1])
}
func writeMap(d map[int]string) {
	for i := 0; i < 19; i++ {
		d[i%2] = string(i)
	}
}
func TestMapmain() {
	var d = make(map[int]string, 10)
	go readMap(d)
	writeMap(d)
	time.Sleep(time.Second)
}

//-------------- sync.Mutex 互斥锁 ---------------------
//结构体可以自动继承匿名内部结构体的所有方法：
type safeDict struct {
	Data map[string]int
	*sync.Mutex
}

//初始化
func NewDict() *safeDict {
	return &safeDict{make(map[string]int), &sync.Mutex{}}
}
func (d *safeDict) Get(key string) int {
	d.Lock()
	defer d.Unlock()
	return d.Data[key]
}
func (d *safeDict) Put(key string, value int) {
	d.Lock()
	defer d.Unlock()
	d.Data[key] = value
}

func TestMutex() {
	m := NewDict()
	go func(d *safeDict) {
		for i := 0; i < 19; i++ {
			fmt.Println("put,", string(i), i)
			d.Put(string(i), i)
		}
	}(m)

	go func(d *safeDict) {
		for i := 0; i < 19; i++ {
			fmt.Println("get,", d.Get(string(i)))
		}
	}(m)

	time.Sleep(time.Second)
}

//-------------- sync.RWMutex 读写锁 ---------------------
//写锁会锁读写，读锁只锁写。
type SafeMap struct {
	Data map[string]int
	*sync.RWMutex
}

//初始化
func NewMap() *SafeMap {
	return &SafeMap{make(map[string]int), &sync.RWMutex{}}
}

//读锁
func (d *SafeMap) Get(key string) int {
	d.RLock()
	defer d.RUnlock()
	return d.Data[key]
}

//写锁
func (d *SafeMap) Put(key string, value int) {
	d.Lock()
	defer d.Unlock()
	d.Data[key] = value
}

func TestSafeMap() {
	m := NewMap()

	go func(d *SafeMap) {
		fmt.Println("get,", d.Get("name"))
	}(m)
	m.Put("name", 999)
}
func main() {
	//TestMapmain()
	TestMutex()
	//TestSafeMap()
}
