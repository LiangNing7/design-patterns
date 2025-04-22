package state

import "fmt"

// 状态接口：定义通用的状态规范标准.
type State interface {
	ToGreen(light *TrafficLight)
	ToYellow(light *TrafficLight)
	ToRed(light *TrafficLight)
}

// 定义一个交通信号灯.
type TrafficLight struct {
	state State
}

// 创建一个交通信号灯，并初始化状态为红灯.
func NewTrafficLight() *TrafficLight {
	return &TrafficLight{
		state: NewRed(),
	}
}

// 设置状态
func (t *TrafficLight) SetState(state State) {
	t.state = state
}

// 转换为绿灯
func (t *TrafficLight) ToGreen() {
	t.state.ToGreen(t)
}

// 转换为黄灯
func (t *TrafficLight) ToYellow() {
	t.state.ToYellow(t)
}

// 转换为红灯
func (t *TrafficLight) ToRed() {
	t.state.ToRed(t)
}

// 定义一个红灯
type Red struct{}

// 创建一个红灯实例
func NewRed() *Red {
	return &Red{}
}

// 转换为绿灯
func (r *Red) ToGreen(light *TrafficLight) {
	fmt.Println("错误，红灯不可以切换为绿灯！")
}

// 转换为黄灯
func (r *Red) ToYellow(light *TrafficLight) {
	fmt.Println("黄灯亮起 5 秒！")
	light.SetState(NewYellow())
}

// 转换为红灯
func (r *Red) ToRed(light *TrafficLight) {
	fmt.Println("错误，已经是红灯！")
}

// 定义一个黄灯
type Yellow struct{}

// 创建一个黄灯实例
func NewYellow() *Yellow {
	return &Yellow{}
}

// 转换为绿灯
func (y Yellow) ToGreen(light *TrafficLight) {
	fmt.Println("绿灯亮起 5 秒！")
	light.SetState(NewGreen())
}

// 转换为黄灯
func (y Yellow) ToYellow(light *TrafficLight) {
	fmt.Println("错误，已经是黄灯！")
}

// 转换为红灯
func (y Yellow) ToRed(light *TrafficLight) {
	fmt.Println("红灯亮起 5 秒！")
	light.SetState(NewRed())
}

// 定义一个绿灯
type Green struct{}

// 创建一个绿灯实例
func NewGreen() *Green {
	return &Green{}
}

// 转换为绿灯
func (g *Green) ToGreen(light *TrafficLight) {
	fmt.Println("错误，已经是绿灯！")
}

// 转换为黄灯
func (g *Green) ToYellow(light *TrafficLight) {
	fmt.Println("黄灯亮起 5 秒！")
	light.SetState(NewYellow())
}

// 转换为红灯
func (g *Green) ToRed(light *TrafficLight) {
	fmt.Println("红灯亮起 5 秒！")
	light.SetState(NewRed())
}
