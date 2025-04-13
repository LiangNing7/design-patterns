package singleton

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const lazyCount = 500

// 调用 GetLazy 函数获取单例实例.
func TestGetLazy(t *testing.T) {
	assert := assert.New(t)

	// 使用时初始化.
	lazy2 := GetLazy()
	lazy3 := GetLazy()

	// 测试 GetLazy 返回的是同一个对象.
	assert.Same(lazy2, lazy3)
}

// 获取 500 次，Lazy 是否总是同一个 Lazy.
func TestParalleGetLazy(t *testing.T) {
	assert := assert.New(t)

	wg := sync.WaitGroup{}

	wg.Add(lazyCount)
	lazies := [lazyCount]*Lazy{}
	for i := 0; i < lazyCount; i++ {
		go func(index int) {
			lazies[index] = GetLazy()
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 1; i < lazyCount; i++ {
		assert.Same(lazies[i], lazies[i-1])
	}
}
