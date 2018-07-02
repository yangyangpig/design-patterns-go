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
