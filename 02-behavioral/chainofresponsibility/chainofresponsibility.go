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
