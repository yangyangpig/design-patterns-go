package main

/**
*模式名称:行为结构->职责链模式
*概念:使多个对象都有机会处理请求，从而避免请求的发送者和接收者之间的耦合关系。将这些对象连成一条链，并沿着这条链传递该请求，直到有一个对象处理它为止
*
*
*
*
*
 */
import "fmt"

type IHandle interface {
	SetSuccessor(IHandle)
	HandleRequest(int)
}

type Handle struct {
	successor IHandle
}

func (this *Handle) SetSuccessor(i IHandle) {
	this.successor = i
}

type CreateA struct {
	Handle
}

func (this *CreateA) HandleRequest(a int) {
	if this.successor != nil {
		this.successor.HandleRequest(a) //可以让CreateB也可以处理请求的参数a，而不需要把参数又重新传进CreateB中，解耦了发送者和接收者
		return
	}
	fmt.Printf("CreateA HandleRequest 输出内容为%d", a)
}

type CreateB struct {
	Handle
}

func (this *CreateB) HandleRequest(a int) {
	if this.successor != nil {
		this.successor.HandleRequest(a)
	}
	fmt.Printf("CreateB HandleRequest 输出内容为%d", a)
}

func main() {
	a := &CreateA{}
	b := &CreateB{}
	a.SetSuccessor(b)
	a.HandleRequest(1)
}
