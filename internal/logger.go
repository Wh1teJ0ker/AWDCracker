package internal

import (
	"log"
	"os"
)

// 初始化日志
func InitLogger() *os.File {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	return file
}
