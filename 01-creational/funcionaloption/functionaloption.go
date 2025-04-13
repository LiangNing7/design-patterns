package funcionaloption

import (
	"net/http"
	"time"
)

// HTTPClient 结构体表示 HTTP 客户端的配置选项.
type HTTPClient struct {
	Timeout time.Duration // Timeout 表示 HTTP 客户端等待响应的最大时间.
}

// DefaultHTTPClient 函数返回 HTTPClient 的默认配置.
func DefaultHTTPClient() HTTPClient {
	return HTTPClient{
		Timeout: 10 * time.Second,
	}
}

// Option 类型用于应用选项到 HTTPClient.
type Option func(*HTTPClient)

// WithTimeout 函数设置 HTTPClient 的超时时间.
func WithTimeout(timeout time.Duration) Option {
	return func(hc *HTTPClient) {
		hc.Timeout = timeout // 设置 HTTP 客户端的超时时间.
	}
}

// NewHTTPClient 函数使用给定选项创建一个新的 HTTP 客户端.
func NewHTTPClient(opts ...Option) *http.Client {
	httpClient := DefaultHTTPClient()
	for _, opt := range opts {
		opt(&httpClient)
	}
	return &http.Client{
		// 根据提供的选项设置 HTTP 客户端的超时时间.
		Timeout: httpClient.Timeout,
	}
}
