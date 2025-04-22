package interpreter

import (
	"fmt"
)

func Example() {
	// 构建表达式：4 + 2 - 3
	expression := Subtract{
		left: &Add{
			left:  &Number{value: 4},
			right: &Number{value: 2},
		},
		right: &Number{value: 3},
	}

	// 解释表达式并计算结果
	result := expression.Interpret()
	fmt.Println("Result:", result)
	// Output:
	// Result: 3
}
