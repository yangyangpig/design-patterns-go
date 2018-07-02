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
