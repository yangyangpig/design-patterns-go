### 模式四 装饰者模式
```go
package main

import "fmt"

type BaseComputerConcept interface {
	Description() string
	Price() float32
}

type TheoreticalCourses struct {
	name  string
	price float32
}

func (t TheoreticalCourses) Description() string {
	return t.name
}

func (t TheoreticalCourses) Price() float32 {
	return t.price
}

//此时如果想增加多一个系统科目，如何在不改变原来的理论科目对象代码前提下实现呢？此刻就需要用到
//装饰者模式，把理论科目对象装饰

type SystemCourses struct {
	theoreticalcouses BaseComputerConcept
	name              string
	price             float32
}

//将需要装饰的对象放到装饰者对象里面
func (s SystemCourses) SetTheoreticalcouses(course BaseComputerConcept) {
	s.theoreticalcouses = course
}

func (s SystemCourses) Description() string {

	return s.theoreticalcouses.Description() + "+" + s.name
}

func (s SystemCourses) Price() float32 {
	return s.theoreticalcouses.Price() + s.price
}

func main() {
	books1 := TheoreticalCourses{name: "深入理解计算机系统", price : 110}
	book2 := SystemCourses{theoreticalcouses:books1, name : "linux系统", price : 96}
	book3 := SystemCourses{theoreticalcouses:book2, name : "网络学习", price : 120}

	fmt.Print(book3.Description())
	fmt.Println(book3.Price())
}
```

#### 总结

>装饰者模式特点：
1、理论上可以无限包装
2、装饰者和被装饰者门之间有相同的超类型
3、想要拓展功能无需修改原有的代码，定义一个装饰者就可以

>适配器和装饰者区别
适配器模式是在类型不匹配的时候使用，目的是将一种类型伪装成另外一种类型，以便我们的代码可以正常使用
装饰者模式中装饰者和被装饰者拥有相同的超类，目的是为了增强功能或者方便使用

>装饰者模式设计原则
1、多用组合，少用继承
2、遵循开闭原则

#### 定义
>装饰模式(Decorator Pattern) ：动态地给一个对象增加一些额外的职责(Responsibility)，就增加对象功能来说，装饰模式比生成子类实现更为灵活。其别名也可以称为包装器(Wrapper)，与适配器模式的别名相同，但它们适用于不同的场合。根据翻译的不同，装饰模式也有人称之为“油漆工模式”，它是一种对象结构型模式

#### 结构
* Component: 抽象构件
* ConcreteComponent: 具体构件
* Decorator: 抽象装饰类
* ConcreteDecorator: 具体装饰类

![](https://github.com/developersPHP/design-patterns-go/blob/master/images/Decorator.jpg)

>时序图

![](https://github.com/developersPHP/design-patterns-go/blob/master/images/seq_Decorator.jpg)

#### 优点

* 装饰模式与继承关系的目的都是要扩展对象的功能，但是装饰模式可以提供比继承更多的灵活性。
* 可以通过一种动态的方式来扩展一个对象的功能，通过配置文件可以在运行时选择不同的装饰器，从而实现不同的行为。
* 通过使用不同的具体装饰类以及这些装饰类的排列组合，可以创造出很多不同行为的组合。可以使用多个具体装饰类来装饰同一对象，得到功能更为强大的对象。
* 具体构件类与具体装饰类可以独立变化，用户可以根据需要增加新的具体构件类和具体装饰类，在使用时再对其进行组合，原有代码无须改变，符合“开闭原则”


#### 缺点

* 使用装饰模式进行系统设计时将产生很多小对象，这些对象的区别在于它们之间相互连接的方式有所不同，而不是它们的类或者属性值有所不同，同时还将产生很多具体装饰类。这些装饰类和小对象的产生将增加系统的复杂度，加大学习与理解的难度。

*这种比继承更加灵活机动的特性，也同时意味着装饰模式比继承更加易于出错，排错也很困难，对于多次装饰的对象，调试时寻找错误可能需要逐级排查，较为烦琐。

#### 使用环境

* 在不影响其他对象的情况下，以动态、透明的方式给单个对象添加职责。
* 需要动态地给一个对象增加功能，这些功能也可以动态地被撤销。
* 当不能采用继承的方式对系统进行扩充或者采用继承不利于系统扩展和维护时。不能采用继承的情况主要有两类：第一类是系统中存在大量独立的扩展，为支持每一种组合将产生大量的子类，使得子类数目呈爆炸性增长；第二类是因为类定义不能继承（如final类）.


