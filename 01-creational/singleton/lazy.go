package singleton

import (
	"fmt"
	"sync"
)

// 示例模式：单例模式 -> 懒汉模式.

var (
	// 初始化一个全局的单例变量.
	lazy *Lazy
	once sync.Once
)

// 定义单例模式类型.
type Lazy struct{}

// 定义 Lazy 类型 SayHi 方法.
func (lz *Lazy) SayHi() {
	fmt.Println("Hi!")
}

// 初始化并获取全局单例实例.
// 这里需要加锁，防止多个协程同时获取实例时，造成的并发不安全
// 懒汉方法在获取并初始化实例时，可能需要传入参数，例如 GetLazy(3)，这样不是很优雅，
// 所以在实际开发中，我更喜欢用饿汉方法.
func GetLazy() *Lazy {
	once.Do(func() {
		lazy = &Lazy{}
	})
	return lazy
}
