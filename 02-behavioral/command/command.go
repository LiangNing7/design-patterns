package command

import "fmt"

// Command 接口定义了命令的执行方法.
type Command interface {
	Execute()
}

// Light 表示灯的接收者.
type Light struct{}

// TurnOn 打开灯的操作.
func (l *Light) TurnOn() {
	fmt.Println("Light is on")
}

// TurnOff 关闭灯的操作.
func (l *Light) TurnOff() {
	fmt.Println("Light is off")
}

// TurnOnCommand 表示开灯命令.
type TurnOnCommand struct {
	light *Light
}

// Execute 执行开灯命令.
func (c *TurnOnCommand) Execute() {
	c.light.TurnOn()
}

// TurnOffCommand 表示关灯命令.
type TurnOffCommand struct {
	light *Light
}

// Execute 执行关灯命令.
func (c *TurnOffCommand) Execute() {
	c.light.TurnOff()
}

// Invoker 表示调用者.
type Invoker struct {
	command Command
}

// ExecuteCommand 调用执行命令.
func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}
