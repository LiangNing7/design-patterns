package factorymethod

import (
	"fmt"
)

// Logger 是日志记录器接口.
type Logger interface {
	Log(message string)
}

// FileLogger 是文件记录器实现.
type FileLogger struct{}

func (f *FileLogger) Log(message string) {
	fmt.Println("Log to file: " + message)
}

// ConsoleLogger 是控制台日志记录器实现.
type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println("Log to console: " + message)
}

// LoggerFactory 是工厂方法接口，定义了创建日志记录器的方法.
type LoggerFactory interface {
	CreateLogger() Logger
}

// FileLoggerFactory 是文件日志记录器的工厂实现.
type FileLoggerFactory struct{}

func (f *FileLoggerFactory) CreateLogger() Logger {
	return &FileLogger{}
}

// ConsoleLoggerFactory 是控制台日志记录器工厂实现.
type ConsoleLoggerFactory struct{}

func (c *ConsoleLoggerFactory) CreateLogger() Logger {
	return &ConsoleLogger{}
}
