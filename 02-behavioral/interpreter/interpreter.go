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
