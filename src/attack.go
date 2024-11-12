package src

import "fmt"

// 接受Ip地址+get参数的payload
func JointIP() {
	fmt.Println("ip:port?")

}

// 接受指定ip
func ReceiveIP() {
	fmt.Println("ip")
}

// 解析http参数，待考虑
func ParseHttp() {
	fmt.Println("解析请求头")
}

// 将所有参数进行组合，得到完整的请求
func SendHttp() {
	fmt.Println("SendHttp")
}

// 注入选定的webshell
func InjectShell() {
	fmt.Println("Injecting...")
}

// 利用上述的shell实现基础的rce功能
func ShellRce() {
	fmt.Println("Rce")
}

// 获取所有flag
func GetFlags() {
	fmt.Println("获取所有flag")
}
