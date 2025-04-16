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
