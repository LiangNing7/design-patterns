package context

import (
	"context"
	"fmt"
	"time"
)

type RequestInfo struct {
	Username  string
	IPAddress string
}

func processRequest(ctx context.Context, req RequestInfo) {
	// 从上下文中获取请求信息.
	username := ctx.Value("request").(RequestInfo).Username
	// 模拟处理请求的逻辑.
	fmt.Printf("Processing request from user %s at IP %s\n", username, req.IPAddress)
	// 等待一段时间模拟处理请求.
	time.Sleep(2 * time.Second)
}
