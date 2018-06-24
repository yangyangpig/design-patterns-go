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

### 总结
>代理模式主要是在某些方法不像被直接修改或者访问的时候，可以用代理的方法，把需要访问或者扩展的方法被包含封装，
写成一个代理，再进行扩展
