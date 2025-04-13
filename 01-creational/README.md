# 简单工厂模式

> [简单工厂模式(Simple Factory Pattern)](https://github.com/LiangNing7/design-patterns/tree/main/01-creational/simplefactory)

简单工厂模式：定义一个工厂函数 / 方法，它可以**根据参数的不同，返回不同类的实例**，被创建的实例通常都具有共同的父类。因为在简单工厂模式中用于创建实例的方法是静态(static)方法，因此简单工厂模式又被称为静态工厂方法(Static Factory Method)模式，它属于类创建型模式。

简单工厂模式的要点在于：当你需要什么，只需要传入一个正确的参数，就可以获取你所需要的对象，而无需知道其创建细节。简单工厂模式结构比较简单，其核心是工厂方法的设计。

下面通过一个简单的示例来演示如何在 Go 语言中实现简单工厂模式。假设我们有一个形状接口 Shape 和两个具体形状类 Circle 和 Rectangle，我们可以通过 NewShape 函数来创建不同类型的形状对象。

```go
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
```

运行以上代码，输出如下：

```bash
$ go test -v .
=== RUN   TestCircle
--- PASS: TestCircle (0.00s)
=== RUN   TestRectangle
--- PASS: TestRectangle (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/01-creational/simplefactory       (cached)
```

简单工厂模式，可以使我们根据需要创建预期类型的实例。但简单工厂模式也存在以下问题：

1. 通过 NewShape 方法，我们可以知道，在简单工厂方法中，为了根据类型创建预期类型的实例，方法中，需要包含 if ... else ...**分支判断**。当类型较小时，问题不大。但是当类型很多时，就会有很多 if ... else ...判断，导致维护和测试难度变大；
2. 另外，如果我们要增加一个新的类型，就需要修改静态工厂方法的业务逻辑，违反了“**开闭原则**【OCP】”。

# 工厂方法模式

> [工厂方法模式（Factory Method Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/01-creational/factorymethod)

![Factory Method Pattern](http://images.liangning7.cn/typora/202504101123279.webp)

工厂方法模式：定义了一个用于创建对象的接口，**让子类决定实例化哪个类**。这种模式通过定义一个工厂接口来创建对象，将对象的实例化延迟到子类的实现。 

以下是一个使用 Go 语言实现的工厂方法模式示例：

```go
package factorymethod

import (
	"fmt"
)

// Logger 是日志记录器接口.
type Logger interface {
	Log(message string)
}

// FileLogger 是文件记录器实现.
type FileLogger struct{}

func (f *FileLogger) Log(message string) {
	fmt.Println("Log to file: " + message)
}

// ConsoleLogger 是控制台日志记录器实现.
type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println("Log to console: " + message)
}

// LoggerFactory 是工厂方法接口，定义了创建日志记录器的方法.
type LoggerFactory interface {
	CreateLogger() Logger
}

// FileLoggerFactory 是文件日志记录器的工厂实现.
type FileLoggerFactory struct{}

func (f *FileLoggerFactory) CreateLogger() Logger {
	return &FileLogger{}
}

// ConsoleLoggerFactory 是控制台日志记录器工厂实现.
type ConsoleLoggerFactory struct{}

func (c *ConsoleLoggerFactory) CreateLogger() Logger {
	return &ConsoleLogger{}
}
```

在这个示例中，定义了日志记录器接口 Logger，以及文件日志记录器实现 FileLogger 和控制台日志记录器实现 ConsoleLogger。定义了工厂方法接口LoggerFactory，以及文件日志记录器工厂实现 FileLoggerFactory 和控制台日志记录器工厂实现 ConsoleLoggerFactory ，分别实现了创建日志记录器的方法。

上述代码测试用例如下：

```go
package factorymethod

import (
	"fmt"
)

// Logger 是日志记录器接口.
type Logger interface {
	Log(message string)
}

// FileLogger 是文件记录器实现.
type FileLogger struct{}

func (f *FileLogger) Log(message string) {
	fmt.Println("Log to file: " + message)
}

// ConsoleLogger 是控制台日志记录器实现.
type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println("Log to console: " + message)
}

// LoggerFactory 是工厂方法接口，定义了创建日志记录器的方法.
type LoggerFactory interface {
	CreateLogger() Logger
}

// FileLoggerFactory 是文件日志记录器的工厂实现.
type FileLoggerFactory struct{}

func (f *FileLoggerFactory) CreateLogger() Logger {
	return &FileLogger{}
}

// ConsoleLoggerFactory 是控制台日志记录器工厂实现.
type ConsoleLoggerFactory struct{}

func (c *ConsoleLoggerFactory) CreateLogger() Logger {
	return &ConsoleLogger{}
}
```

在 ExampleFactoryMethod 函数中 ，我们使用不同的工厂创建不同的日志记录器，并记录日志，展示了工厂方法模式的实现方式。

```bash
$ go test -v .
=== RUN   ExampleFactoryMethod
--- PASS: ExampleFactoryMethod (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/01-creational/factorymethod       (cached)
```

工厂方法模式可以帮助我们避免直接使用 New 关键字实例化具体类，使系统具有更好的可扩展性和灵活性。通过定义一个工厂接口及其具体实现，我们可以根据具体需求选择不同的工厂实现，从而创建不同的产品对象。

# 抽象工厂模式

> [抽象工厂模式（Abstract Factory Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/01-creational/abstractfactory)

![Abstract Factory Pattern](http://images.liangning7.cn/typora/202504101206455.webp)

抽象工厂模式：提供了一个接口用于创建一系列相关对象。这些对象构成一个产品族，而不需要指定具体的类。抽象工厂模式适用于需要创建多个相互关联的对象，使得系统独立于具体产品类，同时提供了易于扩展和替换的能力。

比如说我有一个木门工厂，这个工厂可以生产木门相关的对象，例如：门、门把手等。在实际开发中，我们可以使用这个木门工厂，生产出需要的相关联对象。具体的示例如下：

```go
package abstractfactory

import "fmt"

// DoorFactory 是抽象工厂接口，定义了创建门和门把手的方法.
type DoorFactory interface {
	CreateDoor() Door
	CreateDoorHandle() DoorHandle
}

// Door 是门接口.
type Door interface {
	Open()
	Close()
}

// DoorHandle 是门把手接口.
type DoorHandle interface {
	Press()
}

// WoodenDoorFactory 是一个具体的木门工厂，实现了 DoorFactory 接口.
type WoodenDoorFactory struct{}

func (f *WoodenDoorFactory) CreateDoor() Door {
	return &WoodenDoor{}
}

func (f *WoodenDoorFactory) CreateDoorHandle() DoorHandle {
	return &WoodenDoorHandle{}
}

// WoodenDoor 是木门实现.
type WoodenDoor struct{}

func (d *WoodenDoor) Open() {
	fmt.Println("Wooden door is opened")
}

func (d *WoodenDoor) Close() {
	fmt.Println("Wooden door is closed")
}

// WoodenDoorHandle 是木门把手实现.
type WoodenDoorHandle struct{}

func (h *WoodenDoorHandle) Press() {
	fmt.Println("Press wooden door handle")
}
```

在这个示例中，定义了抽象工厂DoorFactory接口，以及相应的门和门把手接口。然后实现了具体的木门工厂WoodenDoorFactory，并分别实现了创建门和门把手的方法。木门实现WoodenDoor和木门把手实现WoodenDoorHandle分别实现了门和门把手的操作。

上述代码的测试用例如下：

```go
package abstractfactory

func Example() {
	// 创建一个木门工厂.
	woodenFactory := &WoodenDoorFactory{}

	// 使用木门工厂创建门和门把手.
	door := woodenFactory.CreateDoor()
	doorHandle := woodenFactory.CreateDoorHandle()

	// 使用创建的门和门把手.
	door.Open()
	doorHandle.Press()

	// Output:
	// Wooden door is opened
	// Press wooden door handle
}
```

在 ExampleAbstractFactory 函数中，使用木门工厂创建门和门把手，并演示了门和门把手的使用。这个示例展示了抽象工厂模式的实现方式，允许我们创建一组相关的对象，并确保它们协同工作。

运行上述测试用例，输出如下：

```bash
$ go test -v .
=== RUN   Example
--- PASS: Example (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/01-creational/abstractfactory     (cached)
```

**工厂方法模式和抽象工厂模式的区别**

在实际开发过程中，很多开发者经常会把工厂方法模式和抽象工厂模式搞混淆。这里，我来介绍下二者的区别。区别如下：

**目的：**

1. 工厂方法模式旨在定义一个创建对象的接口，但**将具体实例化推迟到子类**。每个子类都可以实现工厂接口以提供自己特定的方法来创建对象；
2. 抽象工厂模式旨在提供一个接口，**用于创建一系列相关或依赖对象的“产品族”，而不需要指定具体的类**。它允许客户端使用抽象接口来创建产品，而不必关心具体实现细节。

**关注点：**

1. 工厂方法模式**着重于单一对象的创建与实例化延迟**，每个子类都可以根据需要更改对象的创建逻辑；
2. 抽象工厂模式**着重于一系列相关对象的创建**，确保这些对象是一起使用的。

# 建造者模式

> [建造者模式（Builder Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/01-creational/builder)

在日常开发中，我们经常会遇到一种情况，就是要构建的对象的**参数特别多**。有些参数是必选的，有些参数是可选。但是有的时候，我们不知道哪些参数是必选的，哪些参数是可选的，各种参数的作用和含义也不是很清楚，很容易传错，而且参数很多，构建起来非常的麻烦。

在这种情况下就可以使用建造者模式。建造者返回给客户一个完整的的产品对象，而客户端无须关心该对象所包含的额外属性和组建方式，这就是建造者模式的设计动机。建造者模式将一个复杂对象的构建与表示分离，使得同样的构建过程可以创建不同的表示。

下面演示是如何通过 `Builder` 模式来构建汽车的，示例代码：

1. 首先抽象出需要构建的产品的能力：

   ```go
   // ICar 汽车，我们要造车了.
   // Icar 车具有以下能力.
   type ICar interface {
   	Speed() int
   	Brand() string
   	Brief()
   }
   ```

2. 定义构造汽车的步骤和方法链，即 `ICarBuilder` 接口：

   ```go
   // ICarBuilder 造一辆车需要具有的部件.
   type ICarBuilder interface {
   	Wheel(wheel int) ICarBuilder
   	Engine(engine string) ICarBuilder
   	Speed(max int) ICarBuilder
   	Brand(brand string) ICarBuilder
   	Build() ICar
   }
   ```

   * B`uild()` 作为最后的方法，返回最终产品。

3. 创建具体汽车对象：

   ```go
   // CarProto 车的原型.
   type CarProto struct {
   	Wheel     int
   	Engine    string
   	MaxSpeed  int
   	BrandName string
   }
   
   // Speed 最大车速.
   func (c *CarProto) Speed() int {
   	return c.MaxSpeed
   }
   
   // Brand 车品牌.
   func (c *CarProto) Brand() string {
   	return c.BrandName
   }
   
   // Brief 简介.
   func (c *CarProto) Brief() {
   	fmt.Println("this is a cool car")
   	fmt.Println("car wheel size: ", c.Wheel)
   	fmt.Println("car MaxSpeed: ", c.MaxSpeed)
   	fmt.Println("car Engine: ", c.Engine)
   }
   ```

   * 使用结构体存储具体属性
   * 定义 `Spedd`、`Brand` 和 `Brief` 方法，用来满足 `ICar` 接口。

4. 创建可逐步配置的建造器

   ```go
   // CarStudio 打算通过成立造车实验室进行造车.
   type CarStudio struct {
   	prototype CarProto
   }
   
   // NewCarStudio 造车工作室.
   func NewCarStudio() ICarBuilder {
   	return &CarStudio{}
   }
   
   // Wheel of car.
   func (c *CarStudio) Wheel(wheel int) ICarBuilder {
   	c.prototype.Wheel = wheel
   	return c
   }
   
   // Engine of car.
   func (c *CarStudio) Engine(engine string) ICarBuilder {
   	c.prototype.Engine = engine
   	return c
   }
   
   func (c *CarStudio) Speed(maxSpeed int) ICarBuilder {
   	c.prototype.MaxSpeed = maxSpeed
   	return c
   }
   
   // Brand of car.
   func (c *CarStudio) Brand(brand string) ICarBuilder {
   	c.prototype.BrandName = brand
   	return c
   }
   ```

   * 使用 `prototype` 暂存配置参数

5. 实现 `Build` 方法：

   ```go
   // Build return a car.
   func (c *CarStudio) Build() ICar {
   	car := &CarProto{
   		Wheel:     c.prototype.Wheel,
   		Engine:    c.prototype.Engine,
   		MaxSpeed:  c.prototype.MaxSpeed,
   		BrandName: c.prototype.BrandName,
   	}
   	return car
   }
   ```

   * 每次 `Build()` 都创建新实例，确保建造器可复用；

**为什么要定义一个 `ICar` 接口？**

* 定义 ICar 接口的主要目的在于对车的行为进行抽象，这样调用者就无需关注具体的车对象是如何实现的，只需依赖接口中的方法。
* 未来如果需要构造不同种类的车（例如跑车、卡车等），只要它们都实现了 ICar 接口，就可以无缝替换。
* 接口可以只关注“能做什么”，而具体实现由 CarProto（或其它可能出现的具体实现）负责，从而符合单一职责原则

**为什么要 `ICarBuilder`接口中的前几个函数都返回`ICarBuilder`？**

* 每个设置属性的方法返回接口本身，让调用者能够链式地进行属性赋值，这样不仅方便而且代码流畅。

**为什么要定义 `CarStudio` 结构体？**

* CarStudio 作为具体的 Builder，负责分步设置车的各个部件。为了存储构造过程中的中间状态，CarStudio 内部持有一个 CarProto 实例作为“构建原型”。

**设计流程**：

1. 定义 ICar 接口，抽象出车的基本行为。
2. 定义 ICarBuilder 接口，明确车的各个部件如何被设置，并支持链式构造。
3. 利用 CarProto 作为车的原型，实现 ICar 接口，存储所有具体的属性值。
4. 使用 CarStudio 作为具体的建造者，通过内部的 prototype 字段累积各项设置，最终构造出一个完整的 ICar 实例。

完整代码：

```go
    package builder

    import (
        "fmt"
    )

    // ICar 汽车，我们要造车了.
    // Icar 车具有以下能力.
    type ICar interface {
        Speed() int
        Brand() string
        Brief()
    }

    // ICarBuilder 造一辆车需要具有的部件.
    type ICarBuilder interface {
        Wheel(wheel int) ICarBuilder
        Engine(engine string) ICarBuilder
        Speed(max int) ICarBuilder
        Brand(brand string) ICarBuilder
        Build() ICar
    }

    // CarProto 车的原型.
    type CarProto struct {
        Wheel     int
        Engine    string
        MaxSpeed  int
        BrandName string
    }

    // Speed 最大车速.
    func (c *CarProto) Speed() int {
        return c.MaxSpeed
    }

    // Brand 车品牌.
    func (c *CarProto) Brand() string {
        return c.BrandName
    }

    // Brief 简介.
    func (c *CarProto) Brief() {
        fmt.Println("this is a cool car")
        fmt.Println("car wheel size: ", c.Wheel)
        fmt.Println("car MaxSpeed: ", c.MaxSpeed)
        fmt.Println("car Engine: ", c.Engine)
    }

    // CarStudio 打算通过成立造车实验室进行造车.
    type CarStudio struct {
        prototype CarProto
    }

    // NewCarStudio 造车工作室.
    func NewCarStudio() ICarBuilder {
        return &CarStudio{}
    }

    // Wheel of car.
    func (c *CarStudio) Wheel(wheel int) ICarBuilder {
        c.prototype.Wheel = wheel
        return c
    }

    // Engine of car.
    func (c *CarStudio) Engine(engine string) ICarBuilder {
        c.prototype.Engine = engine
        return c
    }

    func (c *CarStudio) Speed(maxSpeed int) ICarBuilder {
        c.prototype.MaxSpeed = maxSpeed
        return c
    }

    // Brand of car.
    func (c *CarStudio) Brand(brand string) ICarBuilder {
        c.prototype.BrandName = brand
        return c
    }

    // Build return a car.
    func (c *CarStudio) Build() ICar {
        car := &CarProto{
            Wheel:     c.prototype.Wheel,
            Engine:    c.prototype.Engine,
            MaxSpeed:  c.prototype.MaxSpeed,
            BrandName: c.prototype.BrandName,
        }
        return car
    }
```

上述代码的测试用例如下：

```go
package builder

import (
	"fmt"
	"testing"
)

func TestBuilderCar(t *testing.T) {
	builder := NewCarStudio()
	builder.Brand("sky").Speed(120).Engine("audi")
	car := builder.Build()

	if car.Speed() != 120 {
		t.Fatalf("Builder1 fail expect 120 ,but get %d", car.Speed())
	}
	if car.Brand() != "sky" {
		t.Fatalf("Builder1 fail expect sky ,but get %s", car.Brand())
	}
	fmt.Println(car.Speed())
	fmt.Println(car.Brand())
}

func TestBuilderCarMore(t *testing.T) {
	builder := NewCarStudio()
	builder.Brand("land").Speed(110).Engine("bmw")
	builder.Engine("man made").Brand("panda").Wheel(15)
	car := builder.Build()

	fmt.Println(car.Speed())
	fmt.Println(car.Brand())
	car.Brief()
}

func Example() {
	builder := NewCarStudio()
	builder.Brand("land").Speed(110).Engine("bmw")
	builder.Engine("man made").Brand("panda").Wheel(15)
	car := builder.Build()

	fmt.Println(car.Speed())
	fmt.Println(car.Brand())
	car.Brief()
	// Output:
	// 110
	// panda
	// this is a cool car
	// car wheel size:  15
	// car MaxSpeed:  110
	// car Engine:  man made
}
```

运行上述测试用例，输出如下：

```bash
$ go test -v .
=== RUN   TestBuilderCar
120
sky
--- PASS: TestBuilderCar (0.00s)
=== RUN   TestBuilderCarMore
110
panda
this is a cool car
car wheel size:  15
car MaxSpeed:  110
car Engine:  man made
--- PASS: TestBuilderCarMore (0.00s)
=== RUN   ExampleBuilder
--- PASS: ExampleBuilder (0.00s)
PASS
ok  	github.com/superproj/design-pattern/creational/builder	(cached)
```

# 原型模式

> [原型模式（Prototype Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/01-creational/prototype)

如果你希望生成一个对象，其与另一个对象完全相同，该如何实现呢？如果遍历对象的所有成员，将其依次复制到新对象中，会稍显麻烦，而且有些对象可能会有私有成员变量遗漏。

原型模式将这个克隆的过程委派给了被克隆的实际对象，被克隆的对象就叫做“原型”。原型模式使对象能复制自身，并且暴露到接口中，使客户端面向接口编程时，不知道接口实际对象的情况下生成新的对象。原型模式配合原型管理器使用，使得客户端在不知道具体类的情况下，通过接口管理器得到新的实例，并且包含部分预设定配置。

在 Go 项目开发中，原型模式经常被用到，尤其是在日志包中。

以下是一个使用原型模式的示例代码，演示了如何在 Go 语言中实现原型模式。假设我们有一个简单的圆形形状，我们可以通过原型模式来克隆圆形形状对象：

```go
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
```

上述代码测试用例如下：

```go
package prototype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircleClone(t *testing.T) {
	assert := assert.New(t)

	circle := &Circle{Type: "Circle"}
	cloneCircle := circle.Clone()

	actual := cloneCircle.GetType()
	assert.Equal(circle.GetType(), actual)
}
```

运行上述 单元测试用例，输出如下：

```bash
$ go test -v .
=== RUN   TestCircleClone
--- PASS: TestCircleClone (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/01-creational/prototype   (cached)
```

# 单例模式

> [单例模式（Singleton Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/01-creational/singleton)

单例模式，是最简单的一个模式。在 Go 中，单例模式指的是全局只有一个实例，并且它负责创建自己的对象。单例模式不仅有利于减少内存开支，还有减少系统性能开销、防止多个实例产生冲突等优点。

因为单例模式保证了实例的全局唯一性，而且只被初始化一次，所以比较适合全局共享一个实例，且只需要被初始化一次的场景，例如数据库实例、全局配置、全局任务池等。

单例模式又分为饿汉方式和懒汉方式。**饿汉方式**指全局的单例实例在**包被加载时创建**，而**懒汉方式**指全局的单例实例**在第一次被使用时创建**。你可以看到，这种命名方式非常形象地体现了它们不同的特点。

在实际项目开发中，饿汉方式和懒汉方式，都被大量使用。其中，饿汉方式使用的最多。饿汉方式在初始化的时候创建，因为初始化一般是单协程的，所以不存在并发初始化实例的场景，可以认为是并发安全的。但在实际开发中，为了防止多个协程同时初始化，最好还是在初始化资源的时候加锁处理。饿汉方式的一个小缺点，就是不管实例需不需要，都会被初始化，在某些场景下，会带来一些额外的资源消耗。

懒汉方式跟饿汉方式的优缺点刚好颠倒过来。懒汉方式，在真正使用的时候创建实例，所以可以避免一定的资源浪费。但是懒汉方式可能会面临多个协程都去使用资源，带来一些并发问题，所以实例初始化时，需要加锁。

在实际开发中，其实二者区别不大，你可以根据需要自行选择一个即可。如果不知道如何选择，那就直接选择最直接的饿汉方式。

## 饿汉模式

下面是一个饿汉方式的单例模式代码：

```go
package singleton

// 示例模式：单例模式 -> 饿汉模式.
var eager *Eager

// 定义单例模式类型.
type Eager struct {
	count int
}

// Eager 类型 Inc 方法.
func (e *Eager) Inc() {
	e.count++
}

// 初始化单例模式，这里不建议用 init() 函数，
// 因为使用 init 函数初始化时，
// 开发者无感知，不利于维护，容易出差错.
func InitEager(count int) {
	eager = &Eager{count: count}
}

// GetEager 获取全局的单例实例，这里只读，是并发安全的.
func GetEager() *Eager {
	return eager
}
```

你需要注意，因为实例是在包被导入时初始化的，所以如果初始化耗时，会导致程序加载时间比较长。

其测试代码如下：

```go
package singleton

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// 程序运行时初始化.
	InitEager(3)
	os.Exit(m.Run())
}

// 调用 GetEager 函数获取单例实例.
func TestGetEager(t *testing.T) {
	assert := assert.New(t)

	eager := &Eager{count: 3}
	ins := GetEager()

	assert.Equal(ins, eager)
}
```

## 懒汉模式

懒汉方式是开源项目中使用最多的，但它的缺点是非并发安全，在实际使用时需要加锁。以下是懒汉方式不加锁的一个实现：

```go
package singleton

import (
	"fmt"
	"sync"
)

// 示例模式：单例模式 -> 懒汉模式.

var (
	// 初始化一个全局的单例变量.
	lazy *Lazy
	once sync.Once
)

// 定义单例模式类型.
type Lazy struct{}

// 定义 Lazy 类型 SayHi 方法.
func (lz *Lazy) SayHi() {
	fmt.Println("Hi!")
}

// 初始化并获取全局单例实例.
// 这里需要加锁，防止多个协程同时获取实例时，造成的并发不安全
// 懒汉方法在获取并初始化实例时，可能需要传入参数，例如 GetLazy(3)，这样不是很优雅，
// 所以在实际开发中，我更喜欢用饿汉方法.
func GetLazy() *Lazy {
	once.Do(func() {
		lazy = &Lazy{}
	})
	return lazy
}
```

使用 once.Do 可以确保 lazy 实例全局只被创建一次，once.Do 函数还可以确保当同时有多个创建动作时，只有一个创建动作在被执行。使用这种方法来获取单例实例，既提高了代码效率，又保证了并发安全。

其测试代码如下：

```go
package singleton

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const lazyCount = 500

// 调用 GetLazy 函数获取单例实例.
func TestGetLazy(t *testing.T) {
	assert := assert.New(t)

	// 使用时初始化.
	lazy2 := GetLazy()
	lazy3 := GetLazy()

	// 测试 GetLazy 返回的是同一个对象.
	assert.Same(lazy2, lazy3)
}

// 获取 500 次，Lazy 是否总是同一个 Lazy.
func TestParalleGetLazy(t *testing.T) {
	assert := assert.New(t)

	wg := sync.WaitGroup{}

	wg.Add(lazyCount)
	lazies := [lazyCount]*Lazy{}
	for i := 0; i < lazyCount; i++ {
		go func(index int) {
			lazies[index] = GetLazy()
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 1; i < lazyCount; i++ {
		assert.Same(lazies[i], lazies[i-1])
	}
}
```

在使用懒汉方法时，有一个地方非常烦人，就是如果 GetLazy() 函数需要一个参数来初始化单例实例，那么每次在获取全局的单例实例时，都需要传入初始化参数，这在很多时候，并不需要，并且整个代码也不优雅（臃肿），所以在实际开发中，我更喜欢用饿汉方法。饿汉方法在绝大部分情况下，是完全可以满足开发需求的。

测试结果如下：

```bash
$ go test -v .
=== RUN   TestGetEager
--- PASS: TestGetEager (0.00s)
=== RUN   TestGetLazy
--- PASS: TestGetLazy (0.00s)
=== RUN   TestParalleGetLazy
--- PASS: TestParalleGetLazy (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/01-creational/singleton   (cached)
```

# New 模式

> [New 模式（New Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/01-creational/new)

New 模式，用于创建和初始化结构体实例。

New 模式通常用于封装结构体的创建过程，提供一个包级函数或方法，用于统一实例化对象的方式，避免直接对外暴露结构体实例的构造细节。通过 New 模式，可以更好地封装对象的创建细节，提高代码的可读性和灵活性。

New 模式是 Go 项目开发中最简单，也是用的最多的一种设计模式。简单到，很多开发者已经不认为它是一种设计模式了。

New 模式也是我日常开发中，用的最多，最喜欢用的一种模式，因为它足够简单。在 Go 项目开发中， 如果你看到了以下函数名：`New()`、`NewXXX()`等，多半说明，这是一个使用 New 模式开发的代码。

在实际开发中，我们通常会使用 `New()`、`NewXXX()`等函数，返回一个指针/非指针/接口类型的实例化后的对象。然后，使用返回对象提供的可导出字段、可导出方法进行业务逻辑处理。下面是一个 New 模式的示例代码：

```go
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
```

上述代码的测试用例如下：

```go
package new

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductGetName(t *testing.T) {
	assert := assert.New(t)

	product := NewProduct("Laptop", 20.01)
	assert.Equal("Laptop", product.GetName())
}

func TestProductGetPrice(t *testing.T) {
	assert := assert.New(t)

	product := NewProduct("Laptop", 20.01)
	assert.Equal(20.01, product.GetPrice())
}
```

运行上述测试用例，输出如下：

```bash
$ go test -v .
=== RUN   TestProductGetName
--- PASS: TestProductGetName (0.00s)
=== RUN   TestProductGetPrice
--- PASS: TestProductGetPrice (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/01-creational/new (cached)
```

# 函数选项模式

> [函数选项模式（Functional Options Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/01-creational/funcionaloption)

函数选项模式，是一种常见的设计模式，用于处理函数或对象具有多个可选参数的情况。在Go语言中，选项模式通常通过函数选项(Function Options)的方式实现，允许用户根据需求为函数传递不同的选项参数，以定制函数的行为。这种模式在Go语言中非常灵活和常用，为函数调用提供了更多的灵活性和可定制性。

> Option 模式又称傻瓜模式，常常用于参数传递，一般常常与 New 模式配合使用用于初始化的一个对象的外部参数。

例如，grpc/grpc-go 的 [NewServer](https://github.com/grpc/grpc-go/blob/v1.37.0/server.go#L514) 函数，uber-go/zap 包的 [New](https://github.com/uber-go/zap/blob/v1.16.0/logger.go#L65) 函数都用到了选项模式。使用选项模式，我们可以创建一个带有默认值的 struct 变量，并选择性地修改其中一些参数的值。

选项模式的基本原理是将函数的参数设置为选项类型，通过函数的可变参数机制或接收函数选项参数的方式，根据用户传递的选项为函数进行配置。这样可以避免函数的参数列表过长和混乱，使函数调用更加清晰和易用。

以下代码使用选项模式实现了一个 HTTP 客户端：

```go
package funcionaloption

import (
	"net/http"
	"time"
)

// HTTPClient 结构体表示 HTTP 客户端的配置选项.
type HTTPClient struct {
	Timeout time.Duration // Timeout 表示 HTTP 客户端等待响应的最大时间.
}

// DefaultHTTPClient 函数返回 HTTPClient 的默认配置.
func DefaultHTTPClient() HTTPClient {
	return HTTPClient{
		Timeout: 10 * time.Second,
	}
}

// Option 类型用于应用选项到 HTTPClient.
type Option func(*HTTPClient)

// WithTimeout 函数设置 HTTPClient 的超时时间.
func WithTimeout(timeout time.Duration) Option {
	return func(hc *HTTPClient) {
		hc.Timeout = timeout // 设置 HTTP 客户端的超时时间.
	}
}

// NewHTTPClient 函数使用给定选项创建一个新的 HTTP 客户端.
func NewHTTPClient(opts ...Option) *http.Client {
	httpClient := DefaultHTTPClient()
	for _, opt := range opts {
		opt(&httpClient)
	}
	return &http.Client{
		// 根据提供的选项设置 HTTP 客户端的超时时间.
		Timeout: httpClient.Timeout,
	}
}
```

在上面的示例代码中，我们定义了一个 HTTPClient 结构体和一系列函数，包括 DefaultHTTPClient 函数用于创建默认配置的 HTTPClient，WithTimeout 函数用于设置超时时间选项，以及 NewHTTPClient 函数用于根据传入的选项创建自定义的 HTTP 客户端。

上述代码的测试用例代码如下：

```go
package funcionaloption

import (
	"fmt"
	"time"
)

func Example() {
	// 使用自定义超时时间创建一个新的 HTTP 客户端.
	client := NewHTTPClient(WithTimeout(5 * time.Second))

	// 使用自定义 HTTP 客户端发起 GET 请求.
	resp, err := client.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.Status)
	// Output:
	// Status Code: 200 OK
}
```

执行上述测试用例，结果如下：

```bash
go test -v .
=== RUN   Example
--- PASS: Example (0.15s)
PASS
ok      github.com/LiangNing7/design-patterns/01-creational/funcionaloption     (cached)
```

# 对象池模式

> [对象池模式（Object Pool Pattern）](https://github.com/LiangNing7/design-patterns/tree/main/01-creational/objectpool)

对象池模式（Object Pool Pattern）通过预先实例化一定数量的对象（对象池），在需要时从池中获取对象并在使用完成后将对象归还给池，减少了对象的创建和销毁次数，提高了系统的效率和性能。对象池一般包括初始化、对象获取、对象归还、资源释放等方法，确保对象的有效复用和管理。

对象池模式结构：

* 对象池（Object Pool）：管理和维护一组对象的池实例，提供对象获取、对象归还等方法；
* 对象（Object）：池中存储的具体对象，可被获取和归还；
* 客户端（Client）：从对象池中获取对象并进行操作的用户，完成操作后将对象归还给对象池。

下面是一个简单的对象池模式的Go语言示例，演示了如何在Go语言中实现对象池：

```go
package objectpool

// Object 表示对象池中具体对象.
type Object struct {
	ID int
}

// ObjectPool 表示对象池，存储和管理对象.
type ObjectPool struct {
	objects chan *Object
}

// NewObjectPool 创建对象池并初始化对象.
func NewObjectPool(size int) *ObjectPool {
	pool := &ObjectPool{
		objects: make(chan *Object, size),
	}
	for i := 0; i < size; i++ {
		object := &Object{ID: i}
		pool.objects <- object
	}
	return pool
}

// AcquireObject 从对象从中获取对象.
func (p *ObjectPool) AcquireObject() *Object {
	return <-p.objects
}

// ReleaseObject 将对象归还给对象池.
func (p *ObjectPool) ReleaseObject(object *Object) {
	p.objects <- object
}
```

在上述示例中，我们定义了 Object 表示池中的具体对象，ObjectPool 表示对象池结构。通过 NewObjectPool 函数初始化对象池并创建一定数量的对象。AcquireObject 和 ReleaseObject 方法分别用于从对象池中获取对象和将对象归还给对象池。

上述代码测试用例如下：

```go
package objectpool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjectPool(t *testing.T) {
	assert := assert.New(t)

	pool := NewObjectPool(3)

	object1 := pool.AcquireObject()
	assert.Equal(0, object1.ID)

	object2 := pool.AcquireObject()
	assert.Equal(1, object2.ID)

	pool.ReleaseObject(object1)
	assert.Equal(0, object1.ID)

	object3 := pool.AcquireObject()
	assert.Equal(2, object3.ID)
}
```

运行上述测试用例，输出如下：

```bash
$ go test -v .
=== RUN   TestObjectPool
--- PASS: TestObjectPool (0.00s)
PASS
ok      github.com/LiangNing7/design-patterns/01-creational/objectpool  (cached)
```

