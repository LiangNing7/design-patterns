package singleton

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// 程序运行时初始化.
	InitEager(3)
	os.Exit(m.Run())
}

// 调用 GetEager 函数获取单例实例.
func TestGetEager(t *testing.T) {
	assert := assert.New(t)

	eager := &Eager{count: 3}
	ins := GetEager()

	assert.Equal(ins, eager)
}
