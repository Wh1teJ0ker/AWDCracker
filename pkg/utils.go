package pkg

import (
	"bufio"
	"os"
	"strings"
)

// loadIPs 从给定的文本文件中导入 IP 地址，并返回一个 IP 地址列表
func LoadIPs(filename string) ([]string, error) {
	var ips []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// 逐行读取文件
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// 假设每行都是一个 IP 地址，去掉前后的空白字符
		ip := strings.TrimSpace(line)
		if ip != "" {
			ips = append(ips, ip) // 将有效的 IP 地址添加到切片中
		}
	}
	// 检查读取过程中是否出现错误
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ips, nil
}

func ExportIPs(filename string, ips []string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, ip := range ips {
		_, err := writer.WriteString(ip + "\n") // 每个 IP 地址写入一行
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
