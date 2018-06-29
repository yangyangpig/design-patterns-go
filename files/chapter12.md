### 模式十二 抽象工厂
```go
package main

import "fmt"

//抽象产品A,其中抽象产品可以有多个

type IProductA interface {
	Use()
}

//抽象产品B
type IProductB interface {
	Eat()
}

//抽象工厂
type AbstractFactory interface {
	CreateProductA()
	CreateProductB()
}

//实例产品A
type ProductA1 struct {
}

func (this *ProductA1) Use() {
	fmt.Printf("this is product %s", "A1")
}

type ProductA2 struct {
}

func (this *ProductA2) Use() {
	fmt.Printf("this is product %s", "A2")
}

type ProductB1 struct {
}

func (this *ProductB1) Eat() {
	fmt.Printf("this is product %s", "B1")
}

type ProductB2 struct {
}

func (this *ProductB2) Eat() {
	fmt.Printf("this is product %s", "B2")
}

//建立实例工厂

type ConcreteFactory1 struct {
}

func (this *ConcreteFactory1) CreateProductA() IProductA {
	obj := new(ProductA1)
	return obj
}

func (this *ConcreteFactory1) CreateProductB() IProductB {
	obj := new(ProductB1)
	return obj
}

type ConcreteFactory2 struct {
}

func (this *ConcreteFactory2) CreateProductA() IProductA {
	obj := new(ProductA2)
	return obj
}

func (this *ConcreteFactory2) CreateProductB() IProductB {
	obj := new(ProductB2)
	return obj
}

func main() {
	factory1 := new(ConcreteFactory1)
	productA1 := factory1.CreateProductA()
	productB1 := factory1.CreateProductB()

	productA1.Use()
	productB1.Eat()

	factory2 := new(ConcreteFactory2)
	productA2 := factory2.CreateProductA()
	ProductB2 := factory2.CreateProductB()

	productA2.Use()
	ProductB2.Eat()
}
```

### 总结

#### 定义

>抽象工厂模式(Abstract Factory Pattern)：提供一个创建一系列相关或相互依赖对象的接口，而无须指定它们具体的类。抽象工厂模式又称为Kit模式，属于对象创建型模式

#### 结构
* AbstractFactory：抽象工厂
* ConcreteFactory：具体工厂
* AbstractProduct：抽象产品
* Product：具体产品

![](https://github.com/developersPHP/design-patterns-go/blob/master/images/AbatractFactory.jpg)

>时序图
![](https://github.com/developersPHP/design-patterns-go/blob/master/images/seq_AbatractFactory.jpg)

#### 模式分析

>优点
* 抽象工厂模式隔离了具体类的生成，使得客户并不需要知道什么被创建。由于这种隔离，更换一个具体工厂就变得相对容易。所有的具体工厂都实现了抽象工厂中定义的那些公共接口，因此只需改变具体工厂的实例，就可以在某种程度上改变整个软件系统的行为。
另外，应用抽象工厂模式可以实现高内聚低耦合的设计目的，因此抽象工厂模式得到了广泛的应用。

* 当一个产品族中的多个对象被设计成一起工作时，它能够保证客户端始终只使用同一个产品族中的对象。这对一些需要根据当前环境来决定其行为的软件系统来说，是一种非常实用的设计模式
* 增加新的具体工厂和产品族很方便，无须修改已有系统，符合“开闭原则”。

>缺点
* 在添加新的产品对象时，难以扩展抽象工厂来生产新种类的产品，这是因为在抽象工厂角色中规定了所有可能被创建的产品集合，要支持新种类的产品就意味着要对该接口进行扩展，而这将涉及到对抽象工厂角色及其所有子类的修改，显然会带来较大的不便
* 开闭原则的倾斜性（增加新的工厂和产品族容易，增加新的产品等级结构麻烦）。

#### 适用环境
* 一个系统不应当依赖于产品类实例如何被创建、组合和表达的细节，这对于所有类型的工厂模式都是重要的
* 系统中有多于一个的产品族，而每次只使用其中某一产品族。
* 属于同一个产品族的产品将在一起使用，这一约束必须在系统的设计中体现出来。
* 系统提供一个产品类的库，所有的产品以同样的接口出现，从而使客户端不依赖于具体实现。

#### 模式扩展
* “开闭原则”要求系统对扩展开放，对修改封闭，通过扩展达到增强其功能的目的。对于涉及到多个产品族与多个产品等级结构的系统，其功能增强包括两方面：
  增加产品族：对于增加新的产品族，工厂方法模式很好的支持了“开闭原则”，对于新增加的产品族，只需要对应增加一个新的具体工厂即可，对已有代码无须做任何修改。
  增加新的产品等级结构：对于增加新的产品等级结构，需要修改所有的工厂角色，包括抽象工厂类，在所有的工厂类中都需要增加生产新产品的方法，不能很好地支持“开闭原则”。
  
* 抽象工厂模式的这种性质称为“开闭原则”的倾斜性，抽象工厂模式以一种倾斜的方式支持增加新的产品，它为新产品族的增加提供方便，但不能为新的产品等级结构的增加提供这样的方便。


