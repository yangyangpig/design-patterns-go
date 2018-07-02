### 模式五 代理模式
```go
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
	obj := shanghaijiaoda{name:"上海交大"}
	obj1 := selectschoolname{shcool:obj}
	fmt.Println(obj1.showname())
}
```

### 定义
>代理模式(Proxy Pattern) ：给某一个对象提供一个代 理，并由代理对象控制对原对象的引用。
代理模式的英 文叫做Proxy或Surrogate，它是一种对象结构型模式。

#### 结构
* Subject: 抽象主题角色
* Proxy: 代理主题角色
* RealSubject: 真实主题角色

![](https://github.com/developersPHP/design-patterns-go/blob/master/images/Proxy.jpg)

>时序图

![](https://github.com/developersPHP/design-patterns-go/blob/master/images/seq_Proxy.jpg)





