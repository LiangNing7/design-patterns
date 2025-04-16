package command

func Example() {
	// 创建灯对象.
	light := &Light{}

	// 创建开灯和关灯命令对象.
	turnOnCommand := &TurnOnCommand{light: light}
	turnOffCommand := &TurnOffCommand{light: light}

	// 创建调用者对象并执行命令.
	invoker := Invoker{command: turnOnCommand}
	invoker.ExecuteCommand()

	invoker.command = turnOffCommand
	invoker.ExecuteCommand()
	// Output:
	// Light is on
	// Light is off
}
