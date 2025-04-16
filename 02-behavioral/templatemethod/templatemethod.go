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
