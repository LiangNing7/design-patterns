package prototype

// Shape 接口定义了克隆方法.
type Shape interface {
	Clone() Shape
	GetType() string
}

// Circle 结构体表示圆形.
type Circle struct {
	Type string
}

func (c *Circle) Clone() Shape {
	return &Circle{
		Type: c.Type,
	}
}

func (c *Circle) GetType() string {
	return c.Type
}
