package simplefactory

type ShapeType string

const (
	ShapeTypeCircle    ShapeType = "circle"
	ShapeTypeRectangle ShapeType = "rectangle"
)

// Shape 接口定义.
type Shape interface {
	Draw() string
}

// Circle 结构体.
type Circle struct{}

func (c Circle) Draw() string {
	return "Drawing Circle"
}

// Rectangle 结构体.
type Rectangle struct{}

func (r Rectangle) Draw() string {
	return "Drawing Rectangle"
}

// NewShape 根据 shapeType 创建具体的 Shap 示例.
func NewShape(shapeType ShapeType) Shape {
	switch shapeType {
	case ShapeTypeCircle:
		return Circle{}
	case ShapeTypeRectangle:
		return Rectangle{}
	default:
		return nil
	}
}
