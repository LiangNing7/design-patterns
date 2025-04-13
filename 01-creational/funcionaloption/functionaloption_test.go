package funcionaloption

import (
	"fmt"
	"time"
)

func Example() {
	// 使用自定义超时时间创建一个新的 HTTP 客户端.
	client := NewHTTPClient(WithTimeout(5 * time.Second))

	// 使用自定义 HTTP 客户端发起 GET 请求.
	resp, err := client.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.Status)
	// Output:
	// Status Code: 200 OK
}
