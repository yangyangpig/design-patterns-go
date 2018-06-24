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

### 总结

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
