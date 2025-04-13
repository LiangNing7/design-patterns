package new

// 示例模式：New 模式.

// 定义 Product 结构体, 用来表示一个商品.
type Product struct {
	Name  string
	Price float64
}

// Name 方法返回产品的名字.
func (p *Product) GetName() string {
	return p.Name
}

// Price 方法返回产品的价格.
func (p *Product) GetPrice() float64 {
	return p.Price
}

// 使用 NewProduct 创建并返回实例化后的实例，*Product 类型.
func NewProduct(name string, price float64) *Product {
	return &Product{
		Name:  name,
		Price: price,
	}
}
