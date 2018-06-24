### 模式六 工厂方法模式
```go
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

func main()  {
	obj := new(MulFactory)
	o := obj.CreateOperation()
	o.SetA(12)
	o.SetB(21)
	fmt.Printf("last result is %d", o.GetResult())
}
```

### 总结

>简单工厂模式和工厂方法模式区别
简单工厂模式是通过用户传进来的参数制定生成那种对象，一般通过switch实现路由
工厂方法模式通过制定生成对象的不同的类，通过该类来制定来制定生成那个对象，其对象的生成不用放到一个switch里面路由了
而是在需要的时候，再去定义对应的需要的对象的生成类即可
