### 模式八 模板模式
```go
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
```

### 总结

>模板模式主要是用于当模板不一样时候，但是提供操作的方法统一的时候，就可以把模板类组合到操作类里。
比如以上CarBand就是模板类，而BMWS是统一操作类，把模板CarBand组合到BMWS类里面，统一通过Exec方法提供操作，在这个方法里面可以调用不同模板的方法
