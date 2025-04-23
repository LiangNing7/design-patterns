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
