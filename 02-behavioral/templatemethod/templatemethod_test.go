package templatemethod

import "fmt"

func Example() {
	fmt.Println("=======制作红豆豆浆=======")
	redBeanSoyaMilk := NewRedBeanSoyaMilk()
	DoMake(redBeanSoyaMilk)
	fmt.Println("=======制作花生豆浆=======")
	peanutSoyaMilk := NewPeanutSoyaMilk()
	DoMake(peanutSoyaMilk)
	// Output:
	// =======制作红豆豆浆=======
	// 第 1 步：选择新鲜的豆子
	// 第 2 步：加入上好的红豆
	// 第 3 步：豆子和配料开始浸泡 3H
	// 第 4 步：豆子和配料放入豆浆机榨汁
	// =======制作花生豆浆=======
	// 第 1 步：选择新鲜的豆子
	// 第 2 步：加入上好的花生
	// 第 3 步：豆子和配料开始浸泡 3H
	// 第 4 步：豆子和配料放入豆浆机榨汁
}
