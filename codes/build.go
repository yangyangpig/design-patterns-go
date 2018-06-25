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
