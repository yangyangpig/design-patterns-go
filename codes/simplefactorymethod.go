package main

import "fmt"

type Operation struct {
	a float64
	b float64
}

func (o *Operation) SetA(v float64) {
	o.a = v
}

func (o *Operation) SetB(v float64) {
	o.b = v
}

type OperationI interface {
	GetResult() float64
	SetA(float64)
	SetB(float64)
}

type MulOperation struct {
	Operation
}

func (this *MulOperation) GetResult() float64 {
	return this.a * this.b
}

type IFactory interface {
	CreateOperation() Operation
}

type MulFactory struct {
}

func (self *MulFactory) CreateOperation() OperationI {
	return &(MulOperation{})
}

func main() {
	obj := new(MulFactory)
	o := obj.CreateOperation()
	o.SetA(12)
	o.SetB(21)
	fmt.Printf("last result is %d", o.GetResult())
}
