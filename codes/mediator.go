package main

import "fmt"

//抽象中介者
type AbstractChatroom interface {
	register()
}

//抽象成员
type AbstractMember interface {
	receiveImage()
	receiveText()
	sendImage()
	sendText()
	sendChatRoom()
}

//具体中介者
type chatRoom struct {
	com AbstractMember
}

func (this *chatRoom) register(o interface{}) *chatRoom {

	switch o.(type) {
	case *commonMember:
		this.com = o.(*commonMember)
	case *diamondMember:
		this.com = o.(*diamondMember)
	default:

	}
	fmt.Println("you have register the chatroom")
	return this

}

func (this *chatRoom) sendImage() {
	this.com.sendImage()
	fmt.Println("you have send a image")
}

func (this *chatRoom) sendText() {
	this.com.sendText()
	fmt.Println("you have send  words")
}

type commonMember struct {
}

func (this commonMember) receiveImage() {
	fmt.Println("common member only can send text")
}

func (this commonMember) receiveText() {
	fmt.Println("common member only can send text")
}

func (this commonMember) sendImage() {
	fmt.Println("common member only can send text")
}

func (this commonMember) sendText() {
	fmt.Println("common member only can send text")
}

func (this commonMember) sendChatRoom() {
	fmt.Println("common member only can send text")
}

type diamondMember struct {
}

func (this diamondMember) receiveImage() {
	fmt.Println("diamondmember can send text")
}

func (this diamondMember) sendImage() {
	fmt.Println("diamond member can send image")
}

func (this diamondMember) receiveText() {
	fmt.Println("diamond member can send image")
}

func (this diamondMember) sendText() {
	fmt.Println("diamond member can send image")
}

func (this diamondMember) sendChatRoom() {
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
