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
