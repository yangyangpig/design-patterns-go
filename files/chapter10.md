### 模式十 建造者模式

```go
package main

import "fmt"

//具体的产品类
type Produce struct {
	name  string
	price float64
}

func (this *Produce) SetName(name string) {
	this.name = name
}

func (this *Produce) SetPrice(p float64) {
	this.price = p
}

func (this *Produce) GetName() string {
	return this.name
}

func (this *Produce) GetPrice() float64 {
	return this.price
}

//抽象建造者接口
type Builder interface {
	SetName(name string) Builder
	SetPrice(p float64) Builder
	Build() *Produce
}

//具体建造者类
type ProductBuilder struct {
	p *Produce
}

func (this *ProductBuilder) SetName(name string) Builder {
	if this.p == nil {
		this.p = new(Produce)
	}
	this.p.SetName(name)
	return this
}

func (this *ProductBuilder) SetPrice(p float64) Builder {
	if this.p == nil {
		this.p = new(Produce)
	}

	this.p.SetPrice(p)
	return this

}

func (this *ProductBuilder) Build() *Produce {
	return this.p
}

//规范对象组件创建流程的类
type Director struct {
	build Builder
}

func (this *Director) Create(name string, price float64) *Produce {
	return this.build.SetName(name).SetPrice(price).Build()
}

func main() {
	var builder Builder = &ProductBuilder{}
	var director *Director = &Director{build: builder}
	var p *Produce = director.Create("rich", 12)
	fmt.Printf("foot name is %s,price is %d", p.GetName(), p.GetPrice())
}
```

### 总结

>建造者模式中有以下角色
Product 这是我们要创建的复杂对象(一般都是很复杂的对象才需要使用建造者模式)。
Builder 抽象的一个东西， 主要是用来规范我们的建造者的。
ConcreteBuilder 具体的Builder实现， 这是今天的重点，主要用来根据不用的业务来完成要创建对象的组建的创建。
Director 这个的作用是规范复杂对象的创建流程

>建造者模式可以将一个复杂对象的构造与它的表示分离，使同样的构建过程可以创建不同的表示
