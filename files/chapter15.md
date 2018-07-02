### 模式十五 享元模式

```go
package main

import "fmt"

type geometricPattern interface {

	//几何图案的描述
	draw()
}

type Roundeness struct {
	Color  string
	Radius int64
}

func (this Roundeness) draw() {
	fmt.Printf("the roundensee radius is %d,color is %s", this.Radius, this.Color)
}

//工厂类

type ObjectFactory struct {
	o map[string]interface{}
}

func (this *ObjectFactory) createObject(c string) interface{} {
	g, ok := this.o[c]
	if !ok {
		g = Roundeness{Color: c}
		this.o[c] = g
		fmt.Printf("%s is assigned to map\n", c)
	}
	return g
}

func main() {
	c := [...]string{"Red", "Yellow", "Green"}
	for i := 0; i < len(c); i++ {
		o := &ObjectFactory{o: make(map[string]interface{})}
		//注意，这里如果返回的是interface需要指定一个具体的结构，不然golang不知道会找到那个对象的方法,而且在这里的方法只能是对象方法而不是指针方法
		g := o.createObject(c[i]).(Roundeness)
		g.Radius = 22
		g.draw()
	}

}

```

#### 定义

>享元模式(Flyweight Pattern)：运用共享技术有效地支持大量细粒度对象的复用。系统只使用少量的对象，而这些对象都很相似，状态变化很小，可以实现对象的多次复用。由于享元模式要求能够共享的对象必须是细粒度对象，因此它又称为轻量级模式，它是一种对象结构型模式。

#### 结构
>享元模式包含如下角色：
* Flyweight: 抽象享元类
* ConcreteFlyweight: 具体享元类
* UnsharedConcreteFlyweight: 非共享具体享元类
* FlyweightFactory: 享元工厂类

![](https://github.com/developersPHP/design-patterns-go/blob/master/images/Flyweight.jpg)

>时序图

![](https://github.com/developersPHP/design-patterns-go/blob/master/images/seq_Flyweight.jpg)


#### 优点
* 享元模式的优点在于它可以极大减少内存中对象的数量，使得相同对象或相似对象在内存中只保存一份。
* 享元模式的外部状态相对独立，而且不会影响其内部状态，从而使得享元对象可以在不同的环境中被共享。

#### 缺点
* 享元模式使得系统更加复杂，需要分离出内部状态和外部状态，这使得程序的逻辑复杂化。
* 为了使对象可以共享，享元模式需要将享元对象的状态外部化，而读取外部状态使得运行时间变长

#### 模式扩展
* 单纯享元模式：在单纯享元模式中，所有的享元对象都是可以共享的，即所有抽象享元类的子类都可共享，不存在非共享具体享元类。
* 复合享元模式：将一些单纯享元使用组合模式加以组合，可以形成复合享元对象，这样的复合享元对象本身不能共享，但是它们可以分解成单纯享元对象，而后者则可以共享

> 享元模式与其他模式的联用
* 在享元模式的享元工厂类中通常提供一个静态的工厂方法用于返回享元对象，使用简单工厂模式来生成享元对象。
* 在一个系统中，通常只有唯一一个享元工厂，因此享元工厂类可以使用单例模式进行设计。
* 享元模式可以结合组合模式形成复合享元模式，统一对享元对象设置外部状态。


