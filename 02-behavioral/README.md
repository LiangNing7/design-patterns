# 中介者模式

> [中介者模式（Mediator Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/02-behavioral/mediator)

![image-20250415001905735](http://images.liangning7.cn/typora/202504150019797.png)

中介者模式，主要用来减少对象之间的直接依赖关系，以促进对象之间的松耦合。在中介者模式中，所有对象不再直接彼此通信，而是通过一个中介者对象进行通信。这种模式有助于降低系统的复杂性，提高可维护性，并支持系统的扩展性。

中介者模式的核心概念包括以下几个关键角色：

* 中介者（Mediator）：定义一个接口用于与各个同事对象之间通信。
* 具体中介者（Concrete Mediator）：实现中介者接口，协调各个同事对象的行为。
* 同事类（Colleague）：每个同事类都知道中介者对象，并与中介者通信。

中介者模式的优点包括降低系统的耦合度、集中控制对象之间的交互、促进代码重用和增强可扩展性等。

下面我们通过一个简单的聊天室示例来演示中介者模式在 Go 语言中的实现：

第一步，我们需要定义中介者需要做什么，在这个简单聊天室中，中介者只需要进行转发消息即可。

```go
// Mediator 是中介者接口.
type Mediator interface {
	SendMessage(sender Colleague, message string) // 发送消息给参与者.
}
```

第二步，定义参与者之间要怎么沟通。

```go
// Colleague 是参与者接口
type Colleague interface {
	SendMessage(message string)    // 发送消息给其他参与者
	ReceiveMessage(message string) // 接收消息
	SetMediator(mediator Mediator) // 设置中介者
}
```

第三步，中介者需要管理参与者列表，并且当一个参与者发送消息时，中介者负者将消息转发给其他人。

```go
// ConcreteMediator 是具体中介者.
type ConcreteMediator struct {
	colleagues []Colleague
}

// NewConcreteMediator 创建一个新的具体中介者实例.
func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{colleagues: make([]Colleague, 0)}
}

// AddColleague 添加参与者到中介者.
func (m *ConcreteMediator) AddColleague(colleague Colleague) {
	m.colleagues = append(m.colleagues, colleague)
}

// SendMessage 实现中介者接口，发送消息给其他参与者.
func (m *ConcreteMediator) SendMessage(sender Colleague, message string) {
	for _, c := range m.colleagues {
		if c != sender { // 不发送给自己.
			c.ReceiveMessage(message)
		}
	}
}
```

第四步，实现参与者，用户要能发消息、接收消息和绑定中介：

```go
// User 是用户参与者.
type User struct {
	name     string
	mediator Mediator
}

// NewUser 创建一个新的用户参与者.
func NewUser(name string) *User {
	return &User{name: name}
}

// SendMessage 实现参与者接口，通过中介者发送消息.
func (u *User) SendMessage(message string) {
	u.mediator.SendMessage(u, message)
}

// ReceiveMessage 实现参与者接口，接收消息.
func (u *User) ReceiveMessage(message string) {
	fmt.Printf("[%s] Received message: %s\n", u.name, message)
}

// SetMidiator 实现参与者接口，设置中介者.
func (u *User) SetMidiator(mediator Mediator) {
	u.mediator = mediator
}
```

完整代码如下：

```go
package mediator

import "fmt"

// Mediator 是中介者接口.
type Mediator interface {
	SendMessage(sender Colleague, message string) // 发送消息给参与者.
}

// Colleague 是参与者接口.
type Colleague interface {
	SendMessage(message string)    // 发送消息给其他参与者.
	ReceiveMessage(message string) // 接收消息.
	SetMidiator(mediator Mediator) // 设置中介者
}

// ConcreteMediator 是具体中介者.
type ConcreteMediator struct {
	colleagues []Colleague
}

// NewConcreteMediator 创建一个新的具体中介者实例.
func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{colleagues: make([]Colleague, 0)}
}

// AddColleague 添加参与者到中介者.
func (m *ConcreteMediator) AddColleague(colleague Colleague) {
	m.colleagues = append(m.colleagues, colleague)
}

// SendMessage 实现中介者接口，发送消息给其他参与者.
func (m *ConcreteMediator) SendMessage(sender Colleague, message string) {
	for _, c := range m.colleagues {
		if c != sender { // 不发送给自己.
			c.ReceiveMessage(message)
		}
	}
}

// User 是用户参与者.
type User struct {
	name     string
	mediator Mediator
}

// NewUser 创建一个新的用户参与者.
func NewUser(name string) *User {
	return &User{name: name}
}

// SendMessage 实现参与者接口，通过中介者发送消息.
func (u *User) SendMessage(message string) {
	u.mediator.SendMessage(u, message)
}

// ReceiveMessage 实现参与者接口，接收消息.
func (u *User) ReceiveMessage(message string) {
	fmt.Printf("[%s] Received message: %s\n", u.name, message)
}

// SetMidiator 实现参与者接口，设置中介者.
func (u *User) SetMidiator(mediator Mediator) {
	u.mediator = mediator
}
```

在上述示例中，我们定义了中介者模式的中介者接口、具体中介者类、参与者接口以及具体参与者类。通过中介者模式，用户可以通过中介者对象在聊天室中发送和接收消息，而不需要直接与其他用户进行通信。

上述代码的单元测试用例如下：

```go
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
```

运行上述单元测试用例，输出如下：

```bash
=== RUN   ExampleMediator
--- PASS: ExampleMediator (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/02-behavioral/mediator    0.003s
```

# 观察者模式

> [观察者模式（Obserser Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/02-behavioral/visitor)

为了能在股市中获利，股民们**时刻关注着股市的风吹草动**，其正类似于捉摸不定的数据对象状态。但是大部分是无用功，倒不如股票的价格发生变动的时候主动的通知股民。

再或者，由于 HTTP 无状态连接协议的特性，服务端无法主动推送（Push）消息给 Web 客户端，因此我们常常会用到轮询策略，也就是持续轮番询问服务端状态有无更新。然而当访问高峰期来临时，成千上万的客户端（观察者）轮询会让服务端（被观察者）不堪重负。

观察者模式，定义了一种一对多的依赖关系，让多个观察者对象同时监听并收到被观察对象的状态变化通知。在观察者模式中，主题（Subject）维护一个观察者（Observer）列表，并在状态发生变化时通知观察者。

观察者模式的类图如下：

![image-20250415103234649](http://images.liangning7.cn/typora/202504151032790.png)

在观察者模式中，主要包含以下几个角色：

* Subject（目标）：被观察者，它是指被观察的对象。 从类图中可以看到，类中有一个用来存放观察者对象的Vector 容器（之所以使用Vector而不使用List，是因为多线程操作时，Vector在是安全的，而List则是不安全的），这个Vector容器是被观察者类的核心，另外还有三个方法：attach方法是向这个容器中添加观察者对象；detach方法是从容器中移除观察者对象；notify方法是依次调用观察者对象的对应方法。这个角色可以是接口，也可以是抽象类或者具体的类，因为很多情况下会与其他的模式混用，所以使用抽象类的情况比较多；
* ConcreteSubject（具体目标）：具体目标是目标类的子类，通常它包含经常发生改变的数据，当它的状态发生改变时，向它的各个观察者发出通知。同时它还实现了在目标类中定义的抽象业务逻辑方法（如果有的话）。如果无须扩展目标类，则具体目标类可以省略；
* Observer（观察者）：观察者将对观察目标的改变做出反应，观察者一般定义为接口，该接口声明了更新数据的方法 update()，因此又称为抽象观察者；
* ConcreteObserver（具体观察者）：在具体观察者中维护一个指向具体目标对象的引用，它存储具体观察者的有关状态，这些状态需要和具体目标的状态保持一致；它实现了在抽象观察者 Observer 中定义的 update()方法。通常在实现时，可以调用具体目标类的 attach() 方法将自己添加到目标类的集合中或通过 detach() 方法将自己从目标类的集合中删除。

首先，观察者模式的核心在于解耦主体（Subject）与观察者（Observer）。主体维护一个观察者列表，当状态发生变化时，它会通知所有注册过的观察者。观察者一旦收到通知，会根据自身实现进行相应的处理。

对于主体，要能注册观察者，注销观察者，并能向所有观察者发送通知。这样，我们主题的接口定义如下：

```go
// Subject 主题接口定义了主题对象应该实现的方法.
type Subject interface {
	Register(observer Observer)   // Register 注册一个观察者.
	Deregister(observer Observer) // Deregister 注销一个观察者.
	Notify(message string)        // Notify 发送通知给所有观察者.
}
```

而观察者，只需要根据接受到的通知进行更新即可，其接口定义如下：

```go
// Observer 观察者接口定义了观察者对象应该实现的方法.
type Observer interface {
	Update(message string) // Update 接收更新通知.
}
```

然后我们再实现具体的 Subject：

```go
// ConcreteSubject 具体主题实现了 Subject 接口，维护了观察者列表.
type ConcreteSubject struct {
	observers []Observer // 观察者列表.
}

// NewConcreteSubject 创建一个新的 ConcreteSubject 实例.
func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{}
}

// Register 实现了 Subject 接口的注册方法.
func (s *ConcreteSubject) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

// Deregister 实现了 Subject 接口的注销方法.
func (s *ConcreteSubject) Deregister(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			// s.observers = append(s.observers[:i], s.observers[i+1:]...)
			s.observers = slices.Delete(s.observers, i, i+1)
			break
		}
	}
}

// Notify 实现了 Subject 接口的通知方法，并通知所有观察者.
func (s *ConcreteSubject) Notify(message string) {
	fmt.Println("系统：韭菜们，股票暴涨，大家快买！")
	for _, observer := range s.observers {
		observer.Update(message)
	}
}
```

然后再实现具体的订阅者：

```go
// ConcreteObserver 具体观察者实现了 Observer 接口.
type ConcreteObserver struct {
	name string // 观察者名称.
}

// NewConcreteObserver 创建一个新的 ConcreteObserver 实例.
func NewConcreteObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{name: name}
}

// Update 实现了 Observer 接口的更新方法.
func (o *ConcreteObserver) Update(message string) {
	fmt.Printf("%s: 收到信息<%s>并激进购入股票！\n", o.name, message)
}
```

这样就完成了观察者模式，完整的代码如下：

```go
package observer

import (
	"fmt"
	"slices"
)

// Subject 主题接口定义了主题对象应该实现的方法.
type Subject interface {
	Register(observer Observer)   // Register 注册一个观察者.
	Deregister(observer Observer) // Deregister 注销一个观察者.
	Notify(message string)        // Notify 发送通知给所有观察者.
}

// Observer 观察者接口定义了观察者对象应该实现的方法.
type Observer interface {
	Update(message string) // Update 接收更新通知.
}

// ConcreteSubject 具体主题实现了 Subject 接口，维护了观察者列表.
type ConcreteSubject struct {
	observers []Observer // 观察者列表.
}

// NewConcreteSubject 创建一个新的 ConcreteSubject 实例.
func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{}
}

// Register 实现了 Subject 接口的注册方法.
func (s *ConcreteSubject) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

// Deregister 实现了 Subject 接口的注销方法.
func (s *ConcreteSubject) Deregister(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			// s.observers = append(s.observers[:i], s.observers[i+1:]...)
			s.observers = slices.Delete(s.observers, i, i+1)
			break
		}
	}
}

// Notify 实现了 Subject 接口的通知方法，并通知所有观察者.
func (s *ConcreteSubject) Notify(message string) {
	fmt.Println("系统：韭菜们，股票暴涨，大家快买！")
	for _, observer := range s.observers {
		observer.Update(message)
	}
}

// ConcreteObserver 具体观察者实现了 Observer 接口.
type ConcreteObserver struct {
	name string // 观察者名称.
}

// NewConcreteObserver 创建一个新的 ConcreteObserver 实例.
func NewConcreteObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{name: name}
}

// Update 实现了 Observer 接口的更新方法.
func (o *ConcreteObserver) Update(message string) {
	fmt.Printf("%s: 收到信息<%s>并激进购入股票！\n", o.name, message)
}
```

在上面的示例中，我们定义了观察者模式的主题（Subject）、观察者（Observer）、具体主题（ConcreteSubject）和具体观察者（ConcreteObserver）。在 main 函数中，我们创建了一个具体主题对象，并注册了两个具体观察者。通过 Notify 方法通知观察者接收新的消息，并演示了观察者被移除后再次通知的情况。

通过观察者模式，我们可以实现对象之间的松耦合，让主题和观察者之间的交互更加灵活和可扩展。在 Go 语言中，利用接口和结构体的特性，可以轻松实现观察者模式，提高代码的可读性和可维护性。

上述代码的单元测试用例如下：

```go
package observer

func Example() {
	// 创建一个主题对象.
	subject := NewConcreteSubject()

	// 创建两个观察者对象并注册到主题中.
	observer1 := NewConcreteObserver("韭菜一号")
	observer2 := NewConcreteObserver("韭菜二号")
	subject.Register(observer1)
	subject.Register(observer2)

	// 发送通知给所有观察者.
	subject.Notify("腾讯股票即将暴涨")

	// 注销观察者 observer2.
	subject.Deregister(observer2)

	// 再次发送通知给观察者.
	subject.Notify("字节股票即将暴涨")
	// Output:
	// 系统：韭菜们，股票暴涨，大家快买！
	// 韭菜一号: 收到信息<腾讯股票即将暴涨>并激进购入股票！
	// 韭菜二号: 收到信息<腾讯股票即将暴涨>并激进购入股票！
	// 系统：韭菜们，股票暴涨，大家快买！
	// 韭菜一号: 收到信息<字节股票即将暴涨>并激进购入股票！
}
```

运行上述单元测试用例，输出如下：

```bash
$ go test -v .
=== RUN   Example
--- PASS: Example (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/02-behavioral/observer    (cached)
```

# 命令模式

> [命令模式（Command Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/02-behavioral/command)

命令模式，将请求以命令的形式包裹在对象里面，传递给调用对象，调用对象寻找匹配该命令的对象，将命令给该对象执行。调用对象可以在执行命令前和执行命令后，进行一些通用或者增强操作，例如：排队请求、记录命令运行日志、撤销命令等。命令模式可以将请求的发送者和接收者解耦，使系统更加灵活和可扩展。

命令模式执行命令时，步骤大概可以分为以下三步：

1. 命令被包裹在请求对象里，传递给调用对象；
2. 调用对象查找匹配该命令（可以处理该命令）的对象，将该命令传递给匹配的对象；
3. 该对象执行传递给它的命令。

一般而言，在软件开发中，行为的请求者和行为的执行者是紧密耦合在一起的，调用关系简单易懂，但是这样不容易拓展，有些时候，我们需要记录、撤销、重做处理的时候，不易修改。因此我们需要将命令抽象化，封装起来，不直接调用真正的执行者方法，易于拓展。

举个现实中的例子，比如我们去餐厅吃饭：点餐（写好菜单，发起请求） -->  订单系统处理，生成订单（创建命令） --> 厨师获取到订单，开始做菜（执行命令），在这个过程中我们并没有直接与厨师交谈，不知道那个厨师做，厨师也不知道是哪个顾客需要，只需要按照订单处理就可以。

又比如，我们经常使用智能音响，我经常叫它 ”小度小度，帮我打开空调“，”小度小度，帮我打开窗帘“等等，在整个过程中，我发出命令 --> 小度接受到命令，包装成为请求 --> 让真正接收命令的对象处理（空调或者窗帘控制器），我没有手动去操作空调和窗帘，小度也可以接受各种各样的命令，只要接入它，我都通过它去操作。

在命令模式中，通常包含以下几个关键角色：

1. **命令接口（Command Interface）：**将命令抽象成一个接口，不同的命令执行不同的操作；
2. **具体命令（Concrete Command）：**实现了命令接口，包含了实际执行操作的方法；
3. **调用者（Invoker）：**请求的封装发送者，它通过命令对象来执行请求，不会直接操作接受者，而是直接关联命令对象，间接调用到接受者的相关操作；
4. **接收者（Receiver）：**真正执行命令的对象，接收并执行命令对象指定的操作；
5. **Client(客户端)：**一般我们在客户端中创建调用者对象，具体的命令类，去执行命令。

以下通过一个简单的例子来演示命令模式在Go语言中的应用。假设我们有一个灯和对应的开关作为接收者，可以实现通过命令对象来控制灯的开关操作。

```go
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
```

在以上示例中，我们定义了灯作为接收者 Light，开灯命令 TurnOnCommand 和关灯命令 TurnOffCommand 作为具体命令，以及调用者 Invoker 来执行命令对象。通过命令模式，可以实现对灯开关操作的封装和执行。

```go
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
```

运行上述单元测试用例，输出如下：

```bash
$ go test -v .
=== RUN   Example
--- PASS: Example (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/02-behavioral/command     0.002s
```

# 迭代器模式

> [迭代器模式（Iterator Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/02-behavioral/iterator)

迭代器模式用于提供一种方法来顺序访问一个聚合对象中的各个元素，而无需暴露其内部实现方式。迭代器模式将迭代操作从聚合对象中分离出来，让客户端可以独立地遍历聚合对象中的元素，同时保持代码的简洁性和可维护性。

绝大多数情况下 Go 程序是不需要用迭代器的。因为内置的 slice 和 map 两种容器都可以通过 range 进行遍历，并且这两种容器在性能方面做了足够的优化。只要没有特殊的需求，通常是直接用这两种容器解决问题。即使不得不写了一个自定义容器，我们几乎总是可以实现一个函数，把所有元素（的引用）拷贝到一个 slice 之后返回，这样调用者又可以直接用 range 进行遍历了。

当然某些特殊场合迭代器还是有用武之地。比如迭代器的 `Next()` 是个耗时操作，不能一口气拷贝所有元素；再比如某些条件下需要中断遍历。

在迭代器模式中，通常包含以下几个关键角色：

1. **迭代器（Iterator）接口：**定义访问和遍历元素的方法，包括 `Next()`、`HasNext()` 等；
2. **具体迭代器（Concrete Iterator）：**实现迭代器接口，并提供对具体聚合对象（如数组、列表）的迭代逻辑；
3. **聚合对象（Aggregate）接口：**定义创建迭代器的方法；
4. **具体聚合对象（Concrete Aggregate）：**实现聚合对象接口，提供具体数据结构（如数组、列表）的管理和遍历逻辑。

下面通过一个示例来演示迭代器模式在 Go 语言中的应用。假设我们有一个简单的数组实现的聚合对象，需要实现迭代器模式来让客户端可以遍历数组中的元素。

```go
package iterator

// Iterator 定义迭代器.
type Iterator interface {
	HasNext() bool // 是否有下一个元素.
	Next() string  // 获取下一个元素.
}

// ConcreteIterator 具体迭代器实现.
type ConcreteIterator struct {
	index int      // 迭代器当前位置.
	data  []string // 数据集合.
}

// NewConcreteIterator 创建新的具体迭代器实例.
func NewConcreteIterator(data []string) *ConcreteIterator {
	return &ConcreteIterator{index: 0, data: data}
}

// HasNext 实现迭代器接口，判断是否有下一个元素.
func (it *ConcreteIterator) HasNext() bool {
	return it.index < len(it.data)
}

// Next 实现迭代器接口，获取下一个元素.
func (it *ConcreteIterator) Next() string {
	if !it.HasNext() {
		return ""
	}
	value := it.data[it.index]
	it.index++
	return value
}

// Aggregate 聚合对象接口.
type Aggregate interface {
	CreateIterator() Iterator // 创建迭代器.
}

// ConcreteAggregate 具体聚合对象实现.
type ConcreteAggregate struct {
	data []string // 数据集合.
}

// NewConcreteAggregate 创建新的具体聚合对象实例.
func NewConcreteAggregate(data []string) *ConcreteAggregate {
	return &ConcreteAggregate{data: data}
}

// CreateIterator 实现聚合对象接口，创建迭代器.
func (a *ConcreteAggregate) CreateIterator() Iterator {
	return NewConcreteIterator(a.data)
}
```

在以上示例中，我们定义了迭代器接口 `Iterator` 和具体迭代器 `ConcreteIterator`，以及聚合对象接口 `Aggregate` 和具体聚合对象 `ConcreteAggregate`。通过迭代器模式，客户端可以通过聚合对象获取迭代器并进行遍历操作，而不必直接暴露聚合对象的内部结构。

上述代码的单元测试用例如下：

```go
package iterator

import "fmt"

func Example() {
	// 创建具体聚合对象.
	aggregate := NewConcreteAggregate([]string{"apple", "banana", "cherry", "date"})
	// 获取迭代器.
	iterator := aggregate.CreateIterator()

	// 遍历元素并输出.
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
	// Output:
	// apple
	// banana
	// cherry
	// date
}
```

运行上述单元测试，输出如下：

```bash
$ go test -v .
=== RUN   Example
--- PASS: Example (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/02-behavioral/iterator    0.003s
```

# 模板方法模式

> [模板方法模式（Template Method Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/02-behavioral/templatemethod)

模板方法模式是对多种事物的结构、形式、行为的模式化总结，而模板方法模式（Template Method）则是对一系列类行为（方法）的模式化。我们将总结出来的行为规律固化在基类中，对具体的行为实现则进行抽象化并交给子类去完成，如此便实现了子类对基类模板的套用。

简单来说，模板模式就是将一个类中能够公共使用的方法放置在抽象类中实现，将不能公共使用的方法作为抽象方法，强制子类去实现，这样就做到了将一个类作为一个模板，让开发者去填充需要填充的地方。

例如：编写制作豆浆的程序，说明如下：

1. 制作豆浆的流程：选材—>添加配料—>浸泡—>放到豆浆机打碎
2. 通过添加不同的配料，可以制作出不同口味的豆浆
3. 选材、浸泡和放到豆浆机打碎这几个步骤对于制作每种口味的豆浆都是一样的

以下是模板方法模式的一个实现：

```go
package templatemethod

import "fmt"

// 抽象方法.
type Milker interface {
	// 选择材料.
	SelectBean()

	// 浸泡.
	Soak()

	// 榨汁.
	Beat()

	// 添加配料，子类实现.
	AddCondiment()
}

// 基类
type SoyaMilk struct{}

func (s *SoyaMilk) SelectBean() {
	fmt.Println("第 1 步：选择新鲜的豆子")
}

func (s *SoyaMilk) AddCondiment() {}

func (s *SoyaMilk) Soak() {
	fmt.Println("第 3 步：豆子和配料开始浸泡 3H")
}

func (s *SoyaMilk) Beat() {
	fmt.Println("第 4 步：豆子和配料放入豆浆机榨汁")
}

type RedBeanSoyaMilk struct {
	SoyaMilk
}

func NewRedBeanSoyaMilk() *RedBeanSoyaMilk {
	return &RedBeanSoyaMilk{}
}

func (r *RedBeanSoyaMilk) AddCondiment() {
	fmt.Println("第 2 步：加入上好的红豆")
}

type PeanutSoyaMilk struct {
	SoyaMilk
}

func NewPeanutSoyaMilk() *PeanutSoyaMilk {
	return &PeanutSoyaMilk{}
}

func (p *PeanutSoyaMilk) AddCondiment() {
	fmt.Println("第 2 步：加入上好的花生")
}

// 模板方法
func DoMake(milk Milker) {
	milk.SelectBean()
	milk.AddCondiment()
	milk.Soak()
	milk.Beat()
}
```

上述代码的单元测试用例如下：

```go
package templatemethod

import "fmt"

func Example() {
	fmt.Println("=======制作红豆豆浆=======")
	redBeanSoyaMilk := NewRedBeanSoyaMilk()
	DoMake(redBeanSoyaMilk)
	fmt.Println("=======制作花生豆浆=======")
	peanutSoyaMilk := NewPeanutSoyaMilk()
	DoMake(peanutSoyaMilk)
	// Output:
	// =======制作红豆豆浆=======
	// 第 1 步：选择新鲜的豆子
	// 第 2 步：加入上好的红豆
	// 第 3 步：豆子和配料开始浸泡 3H
	// 第 4 步：豆子和配料放入豆浆机榨汁
	// =======制作花生豆浆=======
	// 第 1 步：选择新鲜的豆子
	// 第 2 步：加入上好的花生
	// 第 3 步：豆子和配料开始浸泡 3H
	// 第 4 步：豆子和配料放入豆浆机榨汁
}
```

运行上述单元测试用例，输出如下：

```bash
$ go test -v .
=== RUN   Example
--- PASS: Example (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/02-behavioral/templatemethod      0.002s
```

# 策略模式

> [策略模式（Strategy Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/02-behavioral/strategy)

策略模式（Strategy Pattern）定义了一系列算法，将每个算法封装起来，并使它们可以互相替换。客户端可以在运行时动态地选择需要的算法。

在什么时候，我们需要用到策略模式呢？

在项目开发中，我们经常要根据不同的场景，采取不同的措施，也就是不同的策略。比如，假设我们需要对 a、b 这两个整数进行计算，根据条件的不同，需要执行不同的计算方式（加法、减法等）。我们可以把所有的操作都封装在同一个函数中，然后通过 `if ... else ...` 的形式来调用不同的计算方式，这种方式称之为硬编码。

在实际应用中，随着功能和体验的不断增长，我们需要经常添加/修改策略，这样就需要不断修改已有代码，不仅会让这个函数越来越难维护，还可能因为修改带来一些bug。所以为了解耦，需要使用策略模式，定义一些独立的类来封装不同的算法，每一个类封装一个具体的算法（即策略）。

下面是一个实现策略模式的代码：

```go
package strategy

// 策略模式.

// 定义一个策略类.
type IStrategy interface {
	Do(int, int) int
}

// 策略实现：加法.
type Add struct{}

func (*Add) Do(a, b int) int {
	return a + b
}

// 策略实现：减法.
type Reduce struct{}

func (*Reduce) Do(a, b int) int {
	return a - b
}

// 具体的策略执行者.
type Operator struct {
	strategy IStrategy
}

// 设置策略
func (o *Operator) SetStrategy(strategy IStrategy) {
	o.strategy = strategy
}

// 调用策略中的方法.
func (o *Operator) Calculate(a, b int) int {
	return o.strategy.Do(a, b)
}

func NewOperator(strategy IStrategy) *Operator {
	return &Operator{strategy: strategy}
}
```

在上述代码中，我们定义了策略接口 `IStrategy`，还定义了 Add 和 Reduce 两种策略。最后定义了一个策略执行者，可以设置不同的策略，并执行。

上述代码的单元测试用例如下：

```go
package strategy

import "fmt"

func Example() {
	operator := NewOperator(&Add{})
	result := operator.Calculate(1, 2)
	fmt.Println("Do add:", result)

	operator.SetStrategy(&Reduce{})
	result = operator.Calculate(2, 1)
	fmt.Println("Do reduce:", result)
	// Output:
	// Do add: 3
	// Do reduce: 1
}
```

运行以上单元测试用例，输出如下：

```bash
$ go test -v .
=== RUN   Example
--- PASS: Example (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/02-behavioral/strategy    0.002s
```

可以看到，我们可以随意更换策略，而不影响 Operator 的所有实现。

# 状态模式

> [状态模式（State Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/02-behavioral/state)

事物状态的变化驱动机制是非常普遍的存在。

状态模式，主要解决的是复杂的状态之间的转换问题。其核心目的就是设计一个状态机，用状态的改变/流转驱动行为变化，同时将状态的实现放在外部，以方便扩展状态。

以交通信号灯为例，它一般包括红、黄、绿3种颜色状态，不同状态之间的切换包含这样的逻辑：红灯只能切换为黄灯，黄灯可以切换为绿灯或红灯，绿灯只能切换为黄灯。

**状态模式的类结构如下图所示：**

![statePattern](http://images.liangning7.cn/typora/202504291316168.webp)

1. State（状态接口）：定义通用的状态规范标准，其中处理请求方法 handle() 将系统环境 Context 作为参数传入。对应本课程中的状态接口 State。
2. ConcreteStateA、ConcreteStateB、ConcreteStateC（状态实现A、状态实现B、状态实现C）：具体的状态实现类，根据系统环境用于表达系统环境 Context 的各个状态，它们都要符合状态接口的规范。对应本章例程中的红灯状态 Red、绿灯状态 Green 以及黄灯状态 Yellow。
3. Context（系统环境）：系统的环境，持有状态接口的引用，以及更新状态方法 setState()，对外暴露请求发起方法 request()，对应本课程中的交通灯类 TrafficLight。

下面我们通过一个简单的信号灯状态转换示例来演示状态模式在 Go 语言中的实现：

```go
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
```

上述代码的单元测试用例如下：

```go
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
```

运行上述单元测试用例，输出如下：

```bash
$ go test -v .
=== RUN   Example
--- PASS: Example (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/02-behavioral/state       0.002s
```

# 备忘录模式

> [备忘录模式（Memento Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/02-behavioral/memento)

浏览器前进与后退、撤销文档修改、数据库备份与恢复、游戏存盘载入、操作系统快照恢复、手机恢复出厂设置等操作稀松平常。

再深入到面向对象层面，我们知道当程序运行时一个对象的状态有可能随时发生变化，而当修改其状态时我们可以对其进行记录，如此便能够将对象恢复到任意记录的状态。备忘录模式正是采用这种理念，让历史重演。我们常常经历文档丢失的痛苦。

![mementoPattern](http://images.liangning7.cn/typora/202504291319359.webp)

备忘录模式（Memento Pattern）允许将对象在不暴露其内部状态的情况下捕获并存储这些状态，并在后续将对象恢复到先前的状态。备忘录模式通常由三个核心角色组成：发起人（Originator）、备忘录（Memento）、管理者（Caretaker）。在实际应用中，备忘录模式常用于实现撤销操作、历史记录功能等。

在备忘录模式中，主要包含以下几个角色：

1. **发起人（Originator）：**负责创建备忘录对象，可以根据需要将当前状态保存到备忘录中，或者从备忘录中恢复状态。
2. **备忘录（Memento）：**用于存储发起人的内部状态，可以包含多个状态属性，并提供读取和设置状态的方法。
3. **管理者（Caretaker）：**负责管理备忘录对象，可以保存和获取备忘录对象，但不会对备忘录对象的状态进行操作。

下面通过一个简单的示例来演示备忘录模式在 Go 语言中的实现：

```go
package memento

import "fmt"

// History 定义历史记录结构体.
type History struct {
	body string // 历史记录内容.
}

// GetBody 获取历史记录内容.
func (h *History) GetBody() string {
	return h.body
}

// SetBody 设置历史记录内容.
func (h *History) SetBody(body string) {
	h.body = body
}

// NewHistory 创建新的历史记录对象实例.
func NewHistory(body string) *History {
	return &History{body: body}
}

// Doc 定义文档结构体.
type Doc struct {
	title string // 文档标题.
	body  string // 文档内容.
}

// Title 获取文档标题.
func (d *Doc) Title() string {
	return d.title
}

// SetTitle 设置文档标题.
func (d *Doc) SetTitle(title string) {
	d.title = title
}

// Body 获取文档内容.
func (d *Doc) Body() string {
	return d.body
}

// SetBody 设置文档内容.
func (d *Doc) SetBody(body string) {
	d.body = body
}

// NewDoc 创建新的文档对象实例.
func NewDoc(title string) *Doc {
	return &Doc{title: title}
}

// CreateHistory 创建文档的历史记录.
func (d *Doc) CreateHistory() *History {
	return &History{body: d.body}
}

// RestoreHistory 恢复文档的历史记录.
func (d *Doc) RestoreHistory(history *History) {
	d.body = history.GetBody()
}

// Editor 定义编辑器结构体.
type Editor struct {
	doc       *Doc       // 当前编辑的文档.
	histories []*History // 编辑历史记录数组.
	position  int        // 当前历史记录位置.
}

// NewEditor 创建新的编辑器实例.
func NewEditor(doc *Doc) *Editor {
	fmt.Println("打开文档" + doc.Title())
	e := &Editor{doc: doc}
	e.histories = make([]*History, 0)
	return e
}

// Append 在文档末尾添加文本内容.
func (e *Editor) Append(text string) {
	e.backup()

	e.doc.SetBody(e.doc.Body() + text + "\n")

	fmt.Println("===> 插入操作，文档内容如下：")
	e.Show()
}

// Save() 保存文档操作.
func (e *Editor) Save() {
	fmt.Println("===> 存盘操作")
}

// Delete 删除文档内容.
func (e *Editor) Delete() {
	e.backup()

	fmt.Println("===> 删除操作，文档内容如下：")
	e.doc.SetBody("")
}

// Show 显示文档内容
func (e *Editor) Show() {
	fmt.Println(e.doc.Body())
}

// backup 备份当前文档状态
func (e *Editor) backup() {
	e.histories = append(e.histories, e.doc.CreateHistory())
	e.position++
}

// Undo 撤销操作，恢复历史状态
func (e *Editor) Undo() {
	// 到头了不可以再撤回
	if e.position == 0 {
		return
	}
	e.position--
	history := e.histories[e.position]
	e.doc.RestoreHistory(history)

	fmt.Println("===> 撤销操作，文档内容如下：")
	e.Show()
}
```

在上面的示例中，我们定义了备忘录模式的发起人（Originator）和备忘录（Memento）接口，以及备忘录的具体实现。

上述代码的测试用例如下：

```go
package memento

func Example() {
	editor := NewEditor(NewDoc("《孔令飞沉思录》"))
	editor.Append("标题：自我晋升篇")
	editor.Append("标题：自我觉醒篇")
	editor.Append("标题：自我反思篇")

	editor.Delete()
	editor.Show()

	editor.Undo()
	editor.Undo()
	// Output:
	// 打开文档《孔令飞沉思录》
	// ===> 插入操作，文档内容如下：
	// 标题：自我晋升篇
	//
	// ===> 插入操作，文档内容如下：
	// 标题：自我晋升篇
	// 标题：自我觉醒篇
	//
	// ===> 插入操作，文档内容如下：
	// 标题：自我晋升篇
	// 标题：自我觉醒篇
	// 标题：自我反思篇
	//
	// ===> 删除操作，文档内容如下：
	//
	// ===> 撤销操作，文档内容如下：
	// 标题：自我晋升篇
	// 标题：自我觉醒篇
	// 标题：自我反思篇
	//
	// ===> 撤销操作，文档内容如下：
	// 标题：自我晋升篇
	// 标题：自我觉醒篇
}
```

在  Example 函数中，我们创建了一个发起人对象 Doc，并展示了创建文档、恢复文档的过程。最终输出了撤销文档后的内容。

通过备忘录模式，我们可以实现对象状态的保存和恢复，使得对象状态的管理更加灵活和可控。在 Go 语言中，通过接口和结构体的组合，可以轻松实现备忘录模式，提高代码的可读性和可维护性。

运行上述测试用例，输入如下：

```bash
$ go test -v .
=== RUN   Example
--- PASS: Example (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/02-behavioral/memento     0.003s
```

# 解释器模式

> [解释器模式（Interpreter Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/02-behavioral/interpreter)

解释器模式，用于定义一种语言的文法，并提供一个解释器来解释该语言中的表达式。解释器模式将文法表示成一个语法树，并通过解释器逐个解释语法树节点，从而实现根据特定规则解释和执行程序。

解释器模式（Interpreter Pattern）的本质是就是自己定一套规则、语言、表达方式,也就是所谓的 DSL（Domain Specific Language），然后按照定义解析执行 DSL，常见的自定义协议，私有协议，就是一个中解释器模式的概念，使用者按照协议的规则做事。

解释器模式的意义在于，它分离多种复杂功能的实现，每个功能只需关注自身的解释。对于调用者不用关心内部的解释器的工作，只需要用简单的方式组合命令。

常见的 Redis 协议就是一个很好解释器模式实现，通过 redis-cli 可以发送各种指令给到 redis-server，服务端解释执行后返回结果。

我们常说的计算器，是一个很典型的解释器模式，它用来解释执行加减乘除的运行规则，现实生活中，哑语翻译、音乐乐谱、摩尔斯电码，道路交通指示系统，等都是解释器模式的好例子。

我们来设计一个简单交流标识系统，来表达会议中两个人交流时候的谈话方式，以代替冗长的语言表述：

```text
A "->"  B 表示  A 说，B 听，此时 B 不能发言。

A "<-"  B 表示  B 说，A 听，此时 B 不能发言。

A "<->" B 表示  A 和 B 可以自由发言。
```

在解释器模式中，通常包含以下几个关键角色：

* **抽象表达式（Abstract Expression）：**定义了解释器的接口，包括一个Interpret()方法，用于解释表达式。
* **终端表达式（Terminal Expression）：**表示语言中的最小单位，实现了抽象表达式的接口。
* **非终端表达式（Non-terminal Expression）：**由多个终端表达式组合而成，也符合抽象表达式的接口。

**解释器模式示例**

以下通过一个简单的例子来演示解释器模式在Go语言中的应用。假设我们有一个简单的算术表达式解释器，可以实现对表达式中加法和减法的解释和计算。

```go
package interpreter

// Expression 是解释器接口，定义了解释器的方法.
type Expression interface {
	Interpret() int
}

// Number 表示一个数字表达式.
type Number struct {
	value int
}

// Interpret 实现了 Expression 接口的 Interpret 方法，返回数字值.
func (n *Number) Interpret() int {
	return n.value
}

// Add 表达加法表达式.
type Add struct {
	left  Expression
	right Expression
}

// Interpret 实现了 Expression 接口的 Interpret 方法，对左右表达式进行相加操作.
func (a *Add) Interpret() int {
	return a.left.Interpret() + a.right.Interpret()
}

// Subtract 表示减法表达式
type Subtract struct {
	left  Expression
	right Expression
}

// Interpret 实现了Expression接口的Interpret方法，对左右表达式进行相减操作
func (s *Subtract) Interpret() int {
	return s.left.Interpret() - s.right.Interpret()
}
```

在以上示例中，我们定义了数字表达式 Number 和加法表达式 Add、减法表达式 Subtract。通过构建不同的表达式组合，客户端可以实现对复杂的算术表达式的解释和计算。

上述代码的单元测试用例如下：

```go
package interpreter

import (
	"fmt"
)

func Example() {
	// 构建表达式：4 + 2 - 3
	expression := Subtract{
		left: &Add{
			left:  &Number{value: 4},
			right: &Number{value: 2},
		},
		right: &Number{value: 3},
	}

	// 解释表达式并计算结果
	result := expression.Interpret()
	fmt.Println("Result:", result)
	// Output:
	// Result: 3
}
```

运行上述单元测试用例，输出如下：

```bash
$ go test -v .
=== RUN   Example
--- PASS: Example (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/02-behavioral/interpreter 0.002s
```

# 责任链模式

> [责任链模式（Chain of Responsibility Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/02-behavioral/chainofresponsibility)

在日常生活中，我们经常会申请各种各样的权限。每个审批过程中又会出现多个节点，每个节点都会有对应的同学进行审批。每个权限节点的同学通常只负责审批自己权限内的资源。

例如，**报销审批的流程**示例图：

![image-20250429132704498](http://images.liangning7.cn/typora/202504291327785.png)

我们来看一个带有一些逻辑的责任链：报销审批流程。

公司为了更高效、安全规范地把控审核工作，通常会将整个审批工作过程按负责人或者工作职责进行拆分，并组织好各个环节中的逻辑关系及走向，最终形成标准化的审批流程。审批流程需要依次通过**财务专员、财务经理、财务总监**的审批。如果申请金额在审批人的审批职权范围内则审批通过并终止流程，反之则会升级至更高层级的上级去继续审批，直至最终的财务总监，如果仍旧超出财务总监的审批金额则驳回申请，流程终止。但是实际的业务场景中这套审批流程可能会**由于业务的变动，而不断变化**。话说要方便地对**业务链条进行拆分、重组，以及对单独节点的增、删、改**。结构松散的业务处理节点让系统具备更加灵活的**可伸缩性、可扩展性**。

责任链模式，用于将请求从一个处理程序传递到另一个处理程序，直到找到能够处理该请求的处理程序为止。在 Go 语言中，责任链模式通常通过函数闭包和链式调用实现，提高代码的灵活性和可扩展性

在下面的示例中，我们将创建一个简单的审批流程链，包括经理、总监和 CFO 三个层级的审批。每个层级有不同的审批金额上限，根据请求金额决定是否批准请求或将请求传递给下一级审批人。代码如下：

```go
package chainofresponsibility

import "fmt"

// IApprover 接口定义了审批人的方法.
type IApprover interface {
	// SetNext 设置下一个审批人.
	SetNext(approver IApprover) IApprover
	// Approver 批准金额.
	Approve(amount float64)
}

// Approver 结构体表示一个审批人.
type Approver struct {
	limit float64
	next  IApprover
}

// NewApprover 创建一个新的审批人.
func NewApprover(limit float64) *Approver {
	return &Approver{limit: limit}
}

// SetNext 设置下一个审批人
func (a *Approver) SetNext(approver IApprover) IApprover {
	a.next = approver
	return a
}

// Approve 空实现了批准金额的操作，具体逻辑由子类实现.
func (a *Approver) Approve(amount float64) {}

// Manager 结构体表示经理职级的审批人.
type Manager struct {
	*Approver
}

// NewManager 创建一个新的经理审批人.
func NewManager(limit float64) *Manager {
	return &Manager{Approver: NewApprover(limit)}
}

// Approve 实现经理审批金额的逻辑.
func (m *Manager) Approve(amount float64) {
	if amount <= m.limit {
		fmt.Printf("Manager approved the request for $%.2f\n", amount)
	} else if m.next != nil {
		m.next.Approve(amount)
	} else {
		fmt.Println("Request cant't be approved at Manager level")
	}
}

// Director 结构体表示总监职级的审批人.
type Director struct {
	*Approver
}

// NewDirector 创建一个新的审批人.
func NewDirector(limit float64) *Director {
	return &Director{Approver: NewApprover(limit)}
}

// Approve 实现总监审批金额的逻辑.
func (d *Director) Approve(amount float64) {
	if amount <= d.limit {
		fmt.Printf("Director approved the request for $%.2f\n", amount)
	} else if d.next != nil {
		d.next.Approve(amount)
	} else {
		fmt.Println("Request can't be approved at Director level")
	}
}

// CFO 结构体表示CFO职级的审批人
type CFO struct {
	*Approver
}

// NewCFO 创建一个新的CFO审批人
func NewCFO(limit float64) *CFO {
	return &CFO{Approver: NewApprover(limit)}
}

// Approve 实现CFO审批金额的逻辑
func (c *CFO) Approve(amount float64) {
	if amount <= c.limit {
		fmt.Printf("CFO approved the request for $%.2f\n", amount)
	} else {
		fmt.Println("Request can't be approved at CFO level")
	}
}
```

在上述示例代码中，我们定义了经理、总监和CFO三个审批层级作为审批处理程序，每个层级根据不同的审批金额上限决定是否批准请求或传递给下一级审批人。通过设置审批处理程序链并处理审批请求，展示了责任链模式在处理审批流。

上述代码单元测试用例如下：

```go
package chainofresponsibility

import (
	"fmt"
)

func Example() {
	// 创建审批处理程序
	manager := NewManager(1000)
	director := NewDirector(5000)
	cfo := NewCFO(10000)

	// 设置审批链
	manager.SetNext(director)
	director.SetNext(cfo)

	// 测试审批流程
	amounts := []float64{800, 3500, 6000, 12000}
	for _, amount := range amounts {
		fmt.Printf("Processing approval request for $%.2f\n", amount)
		manager.Approve(amount)
		fmt.Println()
	}
	// Output:
	// Processing approval request for $800.00
	// Manager approved the request for $800.00
	//
	// Processing approval request for $3500.00
	// Director approved the request for $3500.00
	//
	// Processing approval request for $6000.00
	// CFO approved the request for $6000.00
	//
	// Processing approval request for $12000.00
	// Request can't be approved at CFO level
}
```

运行上述单元测试用例，输出如下：

```bash
$ go test -v .
=== RUN   Example
--- PASS: Example (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/02-behavioral/chainofresponsibility       0.002s
```

# 访问者模式

> [访问者模式（Visitor Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/02-behavioral/visitor)

动物园中有多个场馆，比如豹子馆，海豚馆，大象馆等等，不同场馆的票价是不一样的。另外，动物园针对不同类型的游客收费方式也不一样，比如学生半价。

1. 这个场景下，包括以下要素：动物园，动物园中的各个场馆，不同类型的游客，不同类型的游客票价不同。
2. 动物园就相当于一个对象结构，该结构包含具体的元素（各个场馆），每个场馆（元素）都有接待游客（Visitor）的方法（Accept）。

这些被处理的数据元素相对稳定（动物园中的场馆一般比较稳定）而访问方式多种多样（比如学生散客，学生团体，普通游客，团体游客等不同的访问方式）的数据结构，如果用“访问者模式”来处理比较方便。

访问者模式允许你定义一些操作，可以应用于一个对象结构的元素，而不会改变类。这使得可以在不改变元素类的情况下，新建操作并对其应用于元素类。访问者模式能够使得执行的操作独立于对象结构，同时也增加了在结构上增加新操作的灵活性。

![image-20250429133005543](http://images.liangning7.cn/typora/202504291330882.png)

**访问者模式的结构**

通过上面场景的分析，访问者（Visitor）模式实现的关键是如何将作用于元素的操作分离出来封装成独立的类，其基本结构如下：

* 抽象的访问者（Visitor）：访问具体元素的接口，为每个具体元素类对应一个访问操作 VisitXX() ，其参数为某个具体的元素。
* 具体的访问者（ConcreteVisitor）：实现抽象访问者角色中声明的各个访问操作，确定访问者访问一个元素时该做什么。
* 抽象元素（Element）：声明一个包含接受操作 Accept() 的接口，其参数为访问者对象（游客）。
* 具体元素（ConcreteElement）：实现抽象元素角色提供的 accept() 操作，其方法体通常都是 visitor.visitXX(this) ，另外具体元素中可能还包含本身业务逻辑的相关操作。
* 对象结构（Object Structure）：一个包含元素角色的[容器](https://cloud.tencent.com/product/tke?from=10680)，提供让访问者对象遍历容器中的所有元素的方法，通常由 List、Set、Map 等聚合类实现。本例中的动物园就可抽象成一个对象结构。

在Go语言中，访问者模式可以通过接口和方法的方式来实现。首先，我们定义一个接口用于表示访问者的行为，访问者接口可以定义多个访问方法对应不同类型的元素。然后，我们为每种元素类型定义一个具体的元素结构体，并为这些元素结构体实现接口定义的访问方法。

下面以一个简单的示例来说明访问者模式在Go语言中的应用：

```go
package visitor

import "fmt"

// 抽象访问者接口.
type Visitor interface {
	// 参观猎豹馆.
	VisitLeopardSpot(leopard *LeopardSpot)
	// 参观海豚馆.
	VisitDolphinSpot(dolphin *DolphinSpot)
}

// 场馆景点.
type Scenery interface {
	// 接待访问者.
	Accept(visitor Visitor)
	// 票价.
	Price() int
}

// 定义一个动物园.
type Zoo struct {
	// 动物园包含多个景点.
	Sceneries []Scenery
}

// 创建一个动物园.
func NewZoo() *Zoo {
	return &Zoo{}
}

// 给动物园添加景点.
func (z *Zoo) Add(scenery Scenery) {
	z.Sceneries = append(z.Sceneries, scenery)
}

// 动物园接待游客.
func (z *Zoo) Accept(v Visitor) {
	for _, scenery := range z.Sceneries {
		scenery.Accept(v)
	}
}

// 豹子馆.
type LeopardSpot struct{}

func (l *LeopardSpot) Accept(visitor Visitor) {
	visitor.VisitLeopardSpot(l)
}

// 票价 15 元.
func (l *LeopardSpot) Price() int {
	return 15
}

// 海豚馆
type DolphinSpot struct{}

func NewDolphinSpot() *DolphinSpot {
	return &DolphinSpot{}
}

func (d *DolphinSpot) Accept(visitor Visitor) {
	visitor.VisitDolphinSpot(d)
}

func (d *DolphinSpot) Price() int {
	return 15
}

// 学生的访问游客.
type StudentVisitor struct{}

func NewStudentVisitor() *StudentVisitor {
	return &StudentVisitor{}
}

func (s *StudentVisitor) VisitLeopardSpot(leopard *LeopardSpot) {
	fmt.Printf("学生游客浏览豹子馆票价：%v\n", leopard.Price()/2)
}

func (s *StudentVisitor) VisitDolphinSpot(dolphin *DolphinSpot) {
	fmt.Printf("学生游客游览海豚馆票价: %v\n", dolphin.Price()/2)
}

// 普通游客的访问游客
type CommonVisitor struct{}

func NewCommonVisitor() *CommonVisitor {
	return &CommonVisitor{}
}

func (c *CommonVisitor) VisitLeopardSpot(leopard *LeopardSpot) {
	fmt.Printf("普通游客游览豹子馆票价: %v\n", leopard.Price())
}

func (c *CommonVisitor) VisitDolphinSpot(dolphin *DolphinSpot) {
	fmt.Printf("普通游客游览海豚馆票价: %v\n", dolphin.Price())
}
```

上述代码的单元测试用例如下：

```go
package visitor

func Example() {
	// 创建动物园
	zoo := NewZoo()

	// 添加场景
	zoo.Add(NewDolphinSpot())
	zoo.Add(NewDolphinSpot())

	// 访问场景
	zoo.Accept(NewStudentVisitor())
	zoo.Accept(NewCommonVisitor())
	// Output:
	// 学生游客游览海豚馆票价: 7
	// 学生游客游览海豚馆票价: 7
	// 普通游客游览海豚馆票价: 15
	// 普通游客游览海豚馆票价: 15
}
```

运行上述单元测试用例，输出如下：

```bash
$ go test -v .
=== RUN   Example
--- PASS: Example (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/02-behavioral/visitor     0.002s
```

# 注册表模式

> [注册表模式（Registry Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/02-behavioral/registry)

注册表模式，用于管理应用程序中的全局对象、服务或模块，允许将这些对象注册到一个集中的注册表中，并在需要时进行检索和使用。注册表模式提供了一种灵活的方式来管理和访问全局对象，同时可以实现对象的延迟加载和解耦应用程序组件。

注册表模式基本原理如下：

1. **注册和检索：**注册表模式允许将对象注册到一个全局的注册表中，并通过相应的键来检索和使用这些对象；
2. **延迟加载：**注册表模式支持延迟加载对象，只在需要时才实例化和注册到注册表中；
3. **解耦组件：**通过注册表模式，不同组件之间可以松耦合地进行对象的注册和访问，提高应用程序的灵活性和可维护性。

注册表模式是一种常用的设计模式，特别适用于需要管理和访问全局对象或服务的场景。

在 Go 语言中，可以使用注册表模式来实现全局对象的管理和访问，提高代码的可维护性和扩展性。下面是一个示例，演示了如何在 Go 语言中实现注册表模式：

```go
package registry

// Registry 定义注册表结构.
type Registry struct {
	registry map[string]any
}

// Register 方法用于向注册表中注册对象.
func (r *Registry) Register(key string, value any) {
	r.registry[key] = value
}

// Get 方法用于从注册表中检索对象.
func (r *Registry) Get(key string) any {
	return r.registry[key]
}
```

在上述示例中，我们定义了 Registry 结构体表示注册表，包含 registry 字段用于存储注册的对象。注册表提供了 Register 方法用于注册对象并将其存储在注册表中，以及 Get 方法用于从注册表中检索对象。

上述代码的单元测试用例如下：

```go
package registry

import (
	"fmt"
)

func Example() {
	// 创建注册表实例
	reg := Registry{
		registry: make(map[string]any),
	}

	// 注册对象到注册表中
	reg.Register("logger", "LoggerInstance")
	reg.Register("database", "DatabaseInstance")

	// 从注册表中获取对象并使用
	logger := reg.Get("logger").(string)
	database := reg.Get("database").(string)

	fmt.Println("Logger:", logger)
	fmt.Println("Database:", database)
	// Output:
	// Logger: LoggerInstance
	// Database: DatabaseInstance
}
```

在 ExampleRegistry 函数中，我们创建了注册表实例 reg，并向注册表中注册了日志对象和数据库对象。然后使用 Get 方法从注册表中获取这些对象，并进行使用。

运行上述单元测试用例，输出如下：

```bash
$ go test -v .
=== RUN   Example
--- PASS: Example (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/02-behavioral/registry    0.002s
```

# 上下文模式

> [上下文模式（Context）](https://github.com/LiangNing7/design-patterns/tree/main/02-behavioral/context)

上下文模式（Context）用于在应用程序中传递和管理请求间的各种数据和信息。通过上下文对象，可以在不同组件之间共享数据、状态和配置信息，以实现解耦和提高灵活性。

上下文模式基本原理如下：

* **数据共享：**上下文模式用于在应用程序中传递和共享数据，便于不同组件之间访问和处理；
* **解耦组件：**通过上下文对象，不同组件之间可以松耦合地进行通信和数据传递；
* **请求范围：**上下文对象通常在请求的生命周期内存在，用于传递请求参数、认证信息等。

上下文模式是一种灵活且常用的设计模式，有助于组织和管理应用程序中的数据流和状态信息。

在 Go 语言中，上下文模式经常用于传递请求范围的上下文信息，例如请求参数、用户认证信息等。下面是一个示例，演示了如何在 Go 语言中使用上下文对象传递请求信息：

```go
package context

import (
	"context"
	"fmt"
	"time"
)

type RequestInfo struct {
	Username  string
	IPAddress string
}

func processRequest(ctx context.Context, req RequestInfo) {
	// 从上下文中获取请求信息.
	username := ctx.Value("request").(RequestInfo).Username
	// 模拟处理请求的逻辑.
	fmt.Printf("Processing request from user %s at IP %s\n", username, req.IPAddress)
	// 等待一段时间模拟处理请求.
	time.Sleep(2 * time.Second)
}
```

在上述示例中，我们定义了 RequestInfo 结构体表示请求信息，包括用户名和 IP 地址。在 processRequest 函数中，我们通过上下文对象 context.Context 获取请求信息，并进行处理。

上述代码的单元测试用例如下：

```go
package context

import (
	"context"
)

func Example() {
	// 创建上下文对象，并设置请求信息
	ctx := context.Background()
	req := RequestInfo{
		Username:  "Alice",
		IPAddress: "192.168.1.1",
	}

	// 在上下文中传递请求信息
	ctx = context.WithValue(ctx, "request", req)

	// 处理请求
	processRequest(ctx, req)
	// Output:
	// Processing request from user Alice at IP 192.168.1.1
}
```

在 `ExampleContext` 函数中，我们创建了一个空的上下文对象 ctx，并通过 `context.WithValue` 方法将请求信息 req 存储在上下文中。然后调用 `processRequest` 函数处理请求，其中可以通过上下文对象获取请求信息并进行处理。

运行上述单元测试用例，输出如下：

```bash
$ go test -v .
=== RUN   Example
--- PASS: Example (2.00s)
PASS
ok      github.com/LiangNing7/design-patterns/02-behavioral/context     2.002s
```

