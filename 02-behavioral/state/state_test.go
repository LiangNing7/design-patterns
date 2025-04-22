package state

func Example() {
	traffic := NewTrafficLight()
	traffic.ToYellow()
	traffic.ToGreen()
	traffic.ToRed()
	// Output:
	// 黄灯亮起 5 秒！
	// 绿灯亮起 5 秒！
	// 红灯亮起 5 秒！
}
