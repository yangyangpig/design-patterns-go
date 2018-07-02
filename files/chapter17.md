### 模式十七 中介模式

```go
package main

import "fmt"

//抽象中介者
type AbstractChatroom interface {
	register()
	sendImage()
	sendText()
}

type AbstractMember interface {
	receiveImage()
	receiveText()
	sendImage()
	sendText()
	sendChatRoom()
}

type chatRoom struct {
	com  commonMember
	diam diamondMember
}

func (this *chatRoom) register(o interface{}) *chatRoom {
	switch o.(type) {
	case commonMember:
		this.com = o.(commonMember)
	case diamondMember:
		this.diam = o.(diamondMember)
	default:

	}
	fmt.Println("you have register the chatroom")
	return this

}

func (this *chatRoom) sendImage() {
	this.diam.sendImage()
	fmt.Println("you have send a image")
}

func (this *chatRoom) sendText() {
	this.com.sendText()
	fmt.Println("you have send  words")
}

type commonMember struct {
}

func (this commonMember) sendText() {
	fmt.Println("common member only can send text")
}

type diamondMember struct {
}

func (this diamondMember) sendText() {
	fmt.Println("diamondmember can send text")
}

func (this diamondMember) sendImage() {
	fmt.Println("diamond member can send image")
}

func main() {
	c := new(commonMember)
	d := new(diamondMember)
	chatroon := new(chatRoom)
	chatroon.register(c)
	chatroon.sendText()

	chatroon.register(d)
	chatroon.sendImage()
}

```

#### 定义
>中介者模式(Mediator Pattern)定义：用一个中介对象来封装一系列的对象交互，中介者使各对象不需要显式地相互引用，从而使其耦合松散，而且可以独立地改变它们之间的交互。中介者模式又称为调停者模式，它是一种对象行为型模式。

#### 结构
>中介者模式包含如下角色：
* Mediator: 抽象中介者
* ConcreteMediator: 具体中介者
* Colleague: 抽象同事类
* ConcreteColleague: 具体同事类

![](https://github.com/developersPHP/design-patterns-go/blob/master/images/Mediator.jpg)

>时序图

![](https://github.com/developersPHP/design-patterns-go/blob/master/images/seq_Mediator.jpg)

#### 优点
* 简化了对象之间的交互。
* 将各同事解耦。
* 减少子类生成。
* 可以简化各同事类的设计和实现。

#### 缺点

* 在具体中介者类中包含了同事之间的交互细节，可能会导致具体中介者类非常复杂，使得系统难以维护。

