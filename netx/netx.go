package netx

import (
	"fmt"
	"net"
	"strings"
)

//获取host
func GetHost(host string) (nhost string) {
	nhost = host
	if strings.Index(host, "http://") == 0 || strings.Index(host, "https://") == 0 {
		nhost = strings.Split(host, "//")[1]
	}
	return
}

//通过利用本地解析器查找host，返回主机ipv4和ipv6地址的一个数组
func TestLookupIP(host string) {
	iprecords, _ := net.LookupIP(GetHost(host))
	for i, ip := range iprecords {
		fmt.Printf("TestLookupIP index:%d , ip:%s \n", i, ip)
	}
}

//返回给定名字规范的DNS主机名称，如果调用者不关心name是否规范，可以直接调用LookupHost或者LookupIP,这两个函数都会在查询时考虑到name的规范性
func TestLookupCNAME(host string) {
	canme, _ := net.LookupCNAME(GetHost(host))
	fmt.Println("TestLookupCNAME:", canme)
}

//
func TestLookupAddr(ip string) {
	ptr, e := net.LookupAddr(ip)
	if e != nil {
		fmt.Println("LookupAddr error:", e)
	}
	for i, ptrval := range ptr {
		fmt.Printf("TestLookupAddr index:%d , ip:%s \n", i, ptrval)
	}
}

//查找DNS NS记录
func TestLookupNS(host string) {
	iprecords, _ := net.LookupNS(GetHost(host))
	for i, ip := range iprecords {
		fmt.Printf("TestLookupNS key:%v , v:%v \n", i, ip)
	}
}

//查找DNS NS记录
func TestLookupMX(host string) {
	iprecords, _ := net.LookupMX(GetHost(host))
	for i, ip := range iprecords {
		fmt.Printf("TestLookupMX key:%v , v:%v \n", i, ip)
	}
}

//查找DNS TXT记录
func TestLookupTXT(host string) {
	iprecords, _ := net.LookupTXT(GetHost(host))
	for i, ip := range iprecords {
		fmt.Printf("TestLookupTXT key:%v , v:%v \n", i, ip)
	}
}
func NetXMain() {
	host := "www.baidu.com"
	host1 := "baidu.com"
	ip := "8.8.8.8"
	ip1 := "127.0.0.1"
	TestLookupIP(host)

	TestLookupCNAME(host)

	TestLookupNS(host1)
	TestLookupMX(host1)
	TestLookupTXT(host1)

	//查找DNS PTR记录
	TestLookupAddr(ip)

	fmt.Println("JoinHostPort:", net.JoinHostPort(ip1, "8080"))

}
