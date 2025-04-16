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
