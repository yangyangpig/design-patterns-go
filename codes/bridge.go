package main

import "fmt"

//不同的图案有不同颜色描述，其实这个设计就可以抽象成桥接模式
//对图案抽象，对颜色是实现，不同图案组合不同颜色

type Pattern interface {
	showColor()
}

type Color interface {
	setColor()
	getColor() string
}

//圆形
type Circular struct {
	col Color
}

func (this *Circular) showColor() {
	this.col.setColor()
	fmt.Printf("圆形的颜色是%s\n", this.col.getColor())
}

//方形
type Square struct {
	col Color
}

func (this *Square) showColor() {
	this.col.setColor()
	fmt.Printf("方形的颜色是%s\n", this.col.getColor())
}

type Yellow struct {
	color string
}

func (this *Yellow) setColor() {
	this.color = "黄色"
}

func (this *Yellow) getColor() string {
	c := this.color
	return c
}

type White struct {
	color string
}

func (this *White) setColor() {
	this.color = "白色"
}

func (this *White) getColor() string {
	c := this.color
	return c
}

func main() {
	y := new(Yellow)
	c := &Circular{col: y}
	c.showColor()
	w := new(White)
	q := &Square{col: w}
	q.showColor()
}
