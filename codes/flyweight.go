package main

import "fmt"

/**************************************************************************************************
*模式：结构模式->享元模式flyweight
*概念：主要用于减少创建对象的数量，以减少内存占用和提高性能。这种类型的设计模式属于结构型模式，它提供了减少对象数量从而改善应用所需的对象结构的方式
*角色:
*
*
*
*
*
***************************************************************************************************/
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
