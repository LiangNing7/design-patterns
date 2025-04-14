package mediator

func ExampleMediator() {
	mediator := NewConcreteMediator()

	// 创建三个用户参与者.
	user1 := NewUser("Alice")
	user2 := NewUser("Bob")
	user3 := NewUser("Charlie")

	// 将用户参与者添加到中介者模式.
	mediator.AddColleague(user1)
	mediator.AddColleague(user2)
	mediator.AddColleague(user3)

	// 为每个用户参与者设置中介者.
	user1.SetMidiator(mediator)
	user2.SetMidiator(mediator)
	user3.SetMidiator(mediator)

	// 发送消息给所有用户参与者.
	user1.SendMessage("Hello, everyone!")
	user2.SendMessage("Hi, there!")
	user3.SendMessage("Hey, guys!")

	// Output:
	// [Bob] Received message: Hello, everyone!
	// [Charlie] Received message: Hello, everyone!
	// [Alice] Received message: Hi, there!
	// [Charlie] Received message: Hi, there!
	// [Alice] Received message: Hey, guys!
	// [Bob] Received message: Hey, guys!
}
