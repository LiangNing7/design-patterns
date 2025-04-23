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
