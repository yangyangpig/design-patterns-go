package main

import "fmt"

type ICar interface {
	SetName(string)
	BeforeOut()
	Out()
}

type CarBand struct {
	name string
}

func (this CarBand) SetName(n string) {
	this.name = n
}

func (this CarBand) BeforeOut() {
	this.name = "CARBAND"
	fmt.Printf("car band change name %s\n", this.name)
}

func (this CarBand) Out() {
	this.BeforeOut()
	fmt.Printf("car band name is %s", this.name)
}

type BMWs struct {
	CarBand ICar
}

func (this BMWs) Exec() {
	this.CarBand.Out()
	fmt.Printf("car band change name %s\n", "exec")
}

func main() {
	bmw := new(BMWs)
	bmw.CarBand = new(CarBand)
	bmw.Exec()

}
