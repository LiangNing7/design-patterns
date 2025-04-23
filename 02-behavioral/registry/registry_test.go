package registry

import (
	"fmt"
)

func Example() {
	// 创建注册表实例
	reg := Registry{
		registry: make(map[string]any),
	}

	// 注册对象到注册表中
	reg.Register("logger", "LoggerInstance")
	reg.Register("database", "DatabaseInstance")

	// 从注册表中获取对象并使用
	logger := reg.Get("logger").(string)
	database := reg.Get("database").(string)

	fmt.Println("Logger:", logger)
	fmt.Println("Database:", database)
	// Output:
	// Logger: LoggerInstance
	// Database: DatabaseInstance
}
