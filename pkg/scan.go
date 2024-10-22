package pkg

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

// pingIP 尝试通过发送 ICMP Echo 请求来 ping 一个 IP 地址，返回是否存活
func pingIP(ip string) bool {
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listening to packet: %v\n", err)
		return false
	}
	defer conn.Close()

	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: []byte("HELLO-R-U-THERE"),
		},
	}

	msgBytes, err := msg.Marshal(nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling ICMP message: %v\n", err)
		return false
	}

	conn.SetDeadline(time.Now().Add(2 * time.Second))
	_, err = conn.WriteTo(msgBytes, &net.IPAddr{IP: net.ParseIP(ip)})
	if err != nil {
		return false
	}

	reply := make([]byte, 1500)
	n, _, err := conn.ReadFrom(reply)
	if err != nil {
		return false
	}

	receivedMsg, err := icmp.ParseMessage(1, reply[:n])
	if err != nil {
		return false
	}

	return receivedMsg.Type == ipv4.ICMPTypeEchoReply
}

// ipRange 根据输入的网段生成 IP 地址列表
func ipRange(subnet string) ([]string, error) {
	var ips []string
	ip, ipNet, err := net.ParseCIDR(subnet)
	if err != nil {
		return nil, err
	}

	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	// 去掉第一个和最后一个地址（网络地址和广播地址）
	return ips[1 : len(ips)-1], nil
}

// inc 增加 IP 地址的最后一位，用于迭代 IP 范围
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// Scan 扫描指定子网的存活 IP 地址
func Scan(subnet string) ([]string, error) {
	ips, err := ipRange(subnet)
	if err != nil {
		return nil, err
	}

	var aliveIPs []string
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, ip := range ips {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			if pingIP(ip) {
				// 加锁，确保同时修改切片时不会出现并发问题
				mutex.Lock()
				aliveIPs = append(aliveIPs, ip)
				mutex.Unlock()
				fmt.Printf("%s is alive!\n", ip)
			}
		}(ip)
	}

	wg.Wait()

	return aliveIPs, nil
}
