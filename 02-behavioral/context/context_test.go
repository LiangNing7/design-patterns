package context

import (
	"context"
)

func Example() {
	// 创建上下文对象，并设置请求信息
	ctx := context.Background()
	req := RequestInfo{
		Username:  "Alice",
		IPAddress: "192.168.1.1",
	}

	// 在上下文中传递请求信息
	ctx = context.WithValue(ctx, "request", req)

	// 处理请求
	processRequest(ctx, req)
	// Output:
	// Processing request from user Alice at IP 192.168.1.1
}
