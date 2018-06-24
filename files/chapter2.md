### 模式二 策略模式
```go
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

```
### 总结

>策略模式可以让用户自己通过路由器去选择具体的实现算法逻辑，算法与算法之间互不影响，高度解耦，可扩展性非常高

>策略模式与简单工厂模式区别：其都是通过接口实现不同的对象。但是工厂模式在路由时候仅仅传递一个条件参数就能够获取一个对象，然后通过对象实现算法操作
而策略模式在路由方法中，需要传递一个需要使用的对象，通过该对象调用不同算法。它们最大的区别在于在路由层中传递的是一个条件还是一个对象。
