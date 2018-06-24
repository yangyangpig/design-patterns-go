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
	books1 := TheoreticalCourses{name: "深入理解计算机系统", price: 110}
	book2 := SystemCourses{theoreticalcouses: books1, name: "linux系统", price: 96}
	book3 := SystemCourses{theoreticalcouses: book2, name: "网络学习", price: 120}

	fmt.Print(book3.Description())
	fmt.Println(book3.Price())
}
