package main

import "fmt"

type school interface {
	showname() string
}

type shanghaijiaoda struct {
	name string
}

func (this shanghaijiaoda) showname() string {
	name := this.name
	return name
}

type selectschoolname struct {
	shcool school
}

func (this selectschoolname) showname() string {
	name := this.shcool.showname()
	return name
}

func main() {
	obj := shanghaijiaoda{name: "上海交大"}
	obj1 := selectschoolname{shcool: obj}
	fmt.Println(obj1.showname())
}
