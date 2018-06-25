### 模式7 原型模式
```go
package main

type Base interface {
	SetName(string)
	SetPrice(float64)
}

type Foot struct {
	name  string
	price float64
}

func (this *Foot) SetName(name string) {
	this.name = name
}

func (this *Foot) SetPrice(p float64) {
	this.price = p
}

func (this *Foot) Clone() Base {
  t := this
	return t
}

func main()  {
  obj := new(Foot)
  obj.SetName("大白菜")
  obj.SetPrice(6)
  
  obj1 := obj.Clone()
  obj1.SetPrice(12)
}
```

### 总结

>通过克隆原来的实例来指定创建对象的种类，并且通过拷贝这些原型创建新的对象
>当一个系统应该独立于它的产品创建、构成和表示时，要使用原型模式
>可以避免创建一个和产品类层次平行的工厂类层次时
>当一个类的实例只能有几个不同状态组合中的一种时，建立相应数目的原型并克隆它们可能比每次用合适的状态手工实例化对象更加方便
