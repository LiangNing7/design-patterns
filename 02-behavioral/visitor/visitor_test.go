package visitor

func Example() {
	// 创建动物园
	zoo := NewZoo()

	// 添加场景
	zoo.Add(NewDolphinSpot())
	zoo.Add(NewDolphinSpot())

	// 访问场景
	zoo.Accept(NewStudentVisitor())
	zoo.Accept(NewCommonVisitor())
	// Output:
	// 学生游客游览海豚馆票价: 7
	// 学生游客游览海豚馆票价: 7
	// 普通游客游览海豚馆票价: 15
	// 普通游客游览海豚馆票价: 15
}
