package main

import (
	"fmt"
	"os"
)

type Strategy interface {
	Run()
}

type StrategyA struct {
}

func NewStrategyAFactory() *StrategyA {
	obj := new(StrategyA)
	return obj
}

func (this *StrategyA) Run() {
	fmt.Printf("i am go to airport by %s", "car")
}

type StrategyB struct {
}

func NewStrategyBFactory() *StrategyB {
	obj := new(StrategyB)
	return obj
}

func (this *StrategyB) Run() {
	fmt.Printf("i am go to airport by %s", " own car")
}

type Route struct {
	obj Strategy
}

func (this *Route) declaration(a Strategy) *Route {
	if a, ok := a.(Strategy); ok {
		this.obj = a
	}
	return this

}

func (this *Route) ChoseStrategy(scheme string) *Route {
	obj := new(Route)
	switch scheme {
	case "A":
		obj = obj.declaration(NewStrategyAFactory())
	case "B":
		obj = obj.declaration(NewStrategyBFactory())
	default:

	}
	return obj
}

func main() {
	choose := os.Args[0]
	r := new(Route)
	r = r.ChoseStrategy(choose)
	r.obj.Run()
}
