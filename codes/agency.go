package main

import "fmt"

/**
*模式名称:结构模式->代理模式
*概念:在代理模式中，我们创建具有现有对象的对象，以便向外界提供功能接口
作用:想在访问一个类时做一些控制
*
*
*
*
*
*
********************************************************************************************/

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
