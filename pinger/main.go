package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	count    int
	timeout  int64
	size     int
	autostop bool
)

const (
	ECHO_REQUEST_HEAD_LEN = 8
	ECHO_REPLY_HEAD_LEN   = 20
)

func ArgsInit() {
	flag.IntVar(&count, "c", 4, "要发送的回显请求次数。")
	flag.BoolVar(&autostop, "a", true, "是否自动停止。")
	flag.IntVar(&size, "s", 32, "要发送缓冲区大小。")
	flag.Int64Var(&timeout, "o", 1000, "等待每次回复的超时时间（毫秒）。")

	flag.Parse()

}

func Ping(host string, c chan int, args map[string]interface{}) {
	var count int = args["c"].(int)
	var timeout int64 = args["o"].(int64)
	var size int = args["s"].(int)
	var autostop bool = args["a"].(bool)
	//查找规范的dns主机名字  eg.www.baidu.com->www.a.shifen.com

	if strings.Index(host, "http://") == 0 || strings.Index(host, "https://") == 0 {
		host = strings.Split(host, "//")[1]
	}
	cname, _ := net.LookupCNAME(host)
	fmt.Println("DialTimeout start with :", host)
	//这里只为拿到ip，后续才继续根据ip进行请求
	conn, err := net.DialTimeout("ip4:icmp", host, time.Duration(timeout*1000*1000))
	checkError(err)
	//每个域名可能对应多个ip，但实际连接时，请求只会转发到某一个上，故需要获取实际连接的远程ip，才能知道实际ping的机器是哪台
	ip := conn.RemoteAddr()
	fmt.Printf("正在 Ping %s [%s]  具有 32 字节的数据:\n", cname, ip.String())
	starttime := time.Now()
	var length int = size + ECHO_REQUEST_HEAD_LEN
	var seq int16 = 1

	var msg []byte = make([]byte, length)
	msg[0] = 8
	msg[1] = 0
	msg[2] = 0
	msg[3] = 0

	var receive []byte = make([]byte, length+ECHO_REPLY_HEAD_LEN)

	sendSumDuration := 0
	shotDuration := 0
	longDuration := 0
	sendTimes := 0
	reciveTimes := 0
	lostTimes := 0

	for count > 0 || !autostop {
		sendTimes++

		msg[4], msg[5] = genidentifier(host)
		msg[6], msg[7] = gensequenceint(seq)
		check := checkSum(msg[0:length])
		msg[2], msg[3] = gensequenceuint(check)

		conn, err = net.DialTimeout("ip:icmp", host, time.Duration(timeout*1000*1000))
		checkError(err)
		starttime = time.Now()
		conn.SetDeadline(starttime.Add(time.Duration(timeout * 1000 * 1000)))

		_, err = conn.Write(msg[0:length])

		//接收
		n, err := conn.Read(receive)
		_ = n
		var endduration int = int(int64(time.Since(starttime))) / (1000 * 1000)
		sendSumDuration += endduration
		//除了判断err!=nil，还有判断请求和应答的ID标识符，sequence序列码是否一致，以及ICMP是否超时（receive[ECHO_REPLY_HEAD_LEN] == 11，即ICMP报头的类型为11时表示ICMP超时）
		if err != nil || receive[ECHO_REPLY_HEAD_LEN+4] != msg[4] || receive[ECHO_REPLY_HEAD_LEN+5] != msg[5] || receive[ECHO_REPLY_HEAD_LEN+6] != msg[6] || receive[ECHO_REPLY_HEAD_LEN+7] != msg[7] || endduration >= int(timeout) || receive[ECHO_REPLY_HEAD_LEN] == 11 {
			lostTimes++
			fmt.Println("对 " + cname + "[" + ip.String() + "]" + " 的请求超时。")
		} else {
			if shotDuration > endduration {
				shotDuration = endduration
			}
			if longDuration < endduration {
				longDuration = endduration
			}
			reciveTimes++
			ttl := int(receive[8])
			fmt.Println("来自 " + cname + "[" + ip.String() + "]" + " 的回复: 字节=32 时间=" + strconv.Itoa(endduration) + "ms TTL=" + strconv.Itoa(ttl))
		}
		seq++
		count--
	}
	stat(ip.String(), sendTimes, lostTimes, reciveTimes, shotDuration, longDuration, sendSumDuration)
	c <- 1
}

func stat(ip string, sendN int, lostN int, recvN int, shortT int, longT int, sumT int) {
	fmt.Println()
	fmt.Println(ip, " 的 Ping 统计信息:")
	fmt.Printf("    数据包: 已发送 = %d，已接收 = %d，丢失 = %d (%d%% 丢失)，\n", sendN, recvN, lostN, int(lostN*100/sendN))
	fmt.Println("往返行程的估计时间(以毫秒为单位):")
	if recvN != 0 {
		fmt.Printf("    最短 = %dms，最长 = %dms，平均 = %dms\n", shortT, longT, sumT/sendN)
	}
}

func checkSum(msg []byte) uint16 {
	sum := 0
	length := len(msg)

	for i := 0; i < length-1; i += 2 {
		sum += int(msg[i])*256 + int(msg[i+1])
	}
	if (length & 1) == 1 {
		sum += int(msg[length-1]) * 256
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	return uint16(^sum)
}

func gensequenceint(v int16) (byte, byte) {
	return byte(v >> 8), byte(v & 255)
}
func gensequenceuint(v uint16) (byte, byte) {
	return byte(v >> 8), byte(v & 255)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error :%s ", err.Error())
		os.Exit(-1)
	}
}

func genidentifier(host string) (byte, byte) {
	return host[0], host[1]
}

func main() {

	ArgsInit()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: ", os.Args[0], "host")
		flag.PrintDefaults()
		os.Exit(1)
	}
	//fmt.Println(os.Args[1], os.Args[2])
	fmt.Println(os.Args)
	ch := make(chan int)
	argsmap := map[string]interface{}{}
	argsmap["o"] = timeout
	argsmap["c"] = count
	argsmap["s"] = size
	argsmap["a"] = autostop
	for _, host := range args {
		go Ping(host, ch, argsmap)
	}
	for i := 0; i < len(args); i++ {
		<-ch
	}
	os.Exit(0)

}
