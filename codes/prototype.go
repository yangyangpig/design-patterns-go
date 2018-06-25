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

func main() {
	obj := new(Foot)
	obj.SetName("大白菜")
	obj.SetPrice(6)

	obj1 := obj.Clone()
	obj1.SetPrice(12)
}
