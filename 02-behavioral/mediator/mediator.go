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
