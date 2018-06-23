package main

import (
	"fmt"
	"os"
)

type Car interface {
	Run()
}

type VW struct {
	wheel  string
	engine string
}

func NewVW() VW {
	obj := new(VW)
	return *obj
}

func (this VW) Run() {
	fmt.Printf("%s在奔跑", "vm")
}

type BMW struct {
	wheel  string
	engine string
}

func NewBMW() BMW {
	obj := new(BMW)
	return *obj
}

func (this BMW) Run() {
	fmt.Printf("%s在奔跑", "BMW")
}

type TOYOTA struct {
	wheel  string
	engine string
}

func NewTOYOTA() TOYOTA {
	obj := new(TOYOTA)
	return *obj
}

func (this TOYOTA) Run() {
	fmt.Printf("%s在奔跑", "TOYOTA")
}

type Rout struct {
	obj Car
}

func main() {
	r := new(Rout)
	brand := os.Args[0]
	switch brand {
	case "TOYOTA":
		r.obj = NewTOYOTA()
	case "BMW":
		r.obj = NewBMW()
	case "VW":
		r.obj = NewVW()
	default:

	}
	r.obj.Run()
}
