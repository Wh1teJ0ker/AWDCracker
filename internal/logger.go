package internal

import (
	"fmt"
	"log"
	"os"
)

// Logger 是自定义的日志结构体，包含一个文件指针和日志前缀
type Logger struct {
	file   *os.File
	logger *log.Logger
}

// InitLogger 初始化日志记录器，创建或打开日志文件
func InitLogger() *Logger {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个新的日志记录器，日志信息将写入指定的文件
	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		file:   file,
		logger: logger,
	}
}

// Info 记录信息级别的日志
func (l *Logger) Info(message string) {
	l.logger.SetPrefix("INFO: ")
	l.logger.Println(message)
}

// Warning 记录警告级别的日志
func (l *Logger) Warning(message string) {
	l.logger.SetPrefix("WARNING: ")
	l.logger.Println(message)
}

// Error 记录错误级别的日志
func (l *Logger) Error(message string) {
	l.logger.SetPrefix("ERROR: ")
	l.logger.Println(message)
}

// Close 关闭日志文件
func (l *Logger) Close() {
	if err := l.file.Close(); err != nil {
		fmt.Println("Error closing log file:", err)
	}
}
