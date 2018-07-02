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
