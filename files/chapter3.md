### 模式三 单例模式
```go
package main

import (
	"sync"
	"fmt"
)

type singleton map[string]string

var (
	once     sync.Once
	instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})
	return instance
}

func main() {
	s := New()
	s["test1"] = "dafd"
	s["test2"] = "fff"
	fmt.Println(s)
}
```

### 总结

>单例模式是为了减少对象的声明，从而提高gc效率。对于一些需要对于给对象属性赋值的时候，就需要特别小心使用了，因为每次赋值都会覆盖上一次的值

>单一职责原则，对于一个结构体，只能够有一中引起其变化的因素，如果过多引起变化的因素，就需要考虑再一次抽象和拆成多个。
