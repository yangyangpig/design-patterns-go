package main

/*******************************************************************************************************
*模式名称：结构模式->装饰者模式
*概念：在不改变对象内部结构前提下，动态扩展它的功能，它提供了灵活的方法来拓展对象的功能
*角色：被装饰者，已存在需要扩展的对象
*      抽象接口
*
*
*
********************************************************************************************************* */
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

/*func main() {
	books1 := TheoreticalCourses{name: "深入理解计算机系统", price: 110}
	book2 := SystemCourses{theoreticalcouses: books1, name: "linux系统", price: 96}
	book3 := SystemCourses{theoreticalcouses: book2, name: "网络学习", price: 120}

	fmt.Print(book3.Description())
	fmt.Println(book3.Price())
}*/

type Noodles interface {
	Price() float64
	Describe()
}

type Remen struct {
	name  string
	price float64
}

func (this *Remen) Describe() {
	fmt.Println(this.name)
}

func (this *Remen) Price() float64 {
	return this.price
}

//突然间，需要在这碗面上增加鸡蛋，不可能重新再写一个remen对象（相当于重新点一碗加上鸡蛋的面），所以
//需要在原来的这碗面上，添加一个鸡蛋就可以了，这里remen是被装饰者，egg是装饰者

type Egg struct {
	nooles Noodles
	price  float64
	name   string
}

func (this *Egg) Price() float64 {
	p := this.nooles.Price() + this.price
	fmt.Println(p)
	return p
}

func (this *Egg) Describe() {
	fmt.Println(this.name)
}

func main() {
	remen := &Remen{name: "面条", price: 12}
	remenAdnEgg := &Egg{remen, 6, "remen and egg"}
	remenAdnEgg.Price()
	remenAdnEgg.Describe()
}
