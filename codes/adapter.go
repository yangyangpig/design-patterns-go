package main

import "fmt"

//鸭接口
type Duck interface {
	fly()
	quack()
}

//火鸡接口
type Turkey interface {
	gobble()
	fly()
}

//具体的鸭
type MallardDuck struct {
}

func (this *MallardDuck) fly() {
	fmt.Println("mallardDuck can fly")
}

func (this *MallardDuck) quack() {
	fmt.Println("mallardDuck can quack")
}

//具体火鸡
type WildTurkey struct {
}

func (this *WildTurkey) gobble() {
	fmt.Println("WildTurkey can gobble")
}

func (this *WildTurkey) fly() {
	fmt.Println("WildTurkey can fly")
}

//因为包装了鸭子适配器的火鸡，其外表函数是鸭子，但是实际调用了火鸡方法
type TurkeyAdapter struct {
	turkey Turkey
}

func NewTurkeyAdapter(turkey Turkey) *TurkeyAdapter {
	return &TurkeyAdapter{turkey: turkey}
}

func (this *TurkeyAdapter) fly() {
	this.turkey.fly()
}

func (this *TurkeyAdapter) quack() {
	this.turkey.gobble()
}

func main() {
	t := new(WildTurkey)
	d := NewTurkeyAdapter(t)

	d.fly()
	d.quack()
}
