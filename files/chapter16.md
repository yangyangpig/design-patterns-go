### 模式十六 命令模式

```go
package main

import "fmt"

type command interface {
	execute()
}

type receiver struct {
}

func (this *receiver) action() {
	fmt.Println("the recever has received command")
}

type invoker struct {
	c command
}

func (this *invoker) action() {
	this.c.execute()
}

type createcommand struct {
	r *receiver
}

func (this *createcommand) execute() {
	this.r.action()
}

func main() {
	r := new(receiver)
	c := &createcommand{r}

	invoker := invoker{c}
	invoker.action()
}
```
#### 定义
>命令模式(Command Pattern)：将一个请求封装为一个对象，从而使我们可用不同的请求对客户进行参数化；对请求排队或者记录请求日志，以及支持可撤销的操作。命令模式是一种对象行为型模式，其别名为动作(Action)模式或事务(Transaction)模式。

#### 结构
>命令模式包含如下角色：

* Command: 抽象命令类
* ConcreteCommand: 具体命令类
* Invoker: 调用者
* Receiver: 接收者
* Client:客户类

![](https://github.com/developersPHP/design-patterns-go/blob/master/images/Command.jpg)

>时序图

![](https://github.com/developersPHP/design-patterns-go/blob/master/images/seq_Command.jpg)

#### 优点
* 降低系统的耦合度。
* 新的命令可以很容易地加入到系统中
* 可以比较容易地设计一个命令队列和宏命令（组合命令）。
* 可以方便地实现对请求的Undo和Redo。

#### 缺点
* 使用命令模式可能会导致某些系统有过多的具体命令类。因为针对每一个命令都需要设计一个具体命令类，因此某些系统可能需要大量具体命令类，这将影响命令模式的使用。

