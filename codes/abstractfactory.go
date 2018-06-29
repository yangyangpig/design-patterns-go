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
	fmt.Printf("this is product %s\n", "A1")
}

type ProductA2 struct {
}

func (this *ProductA2) Use() {
	fmt.Printf("this is product %s\n", "A2")
}

type ProductB1 struct {
}

func (this *ProductB1) Eat() {
	fmt.Printf("this is product %s\n", "B1")
}

type ProductB2 struct {
}

func (this *ProductB2) Eat() {
	fmt.Printf("this is product %s\n", "B2")
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
