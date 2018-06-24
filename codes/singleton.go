package main

import (
	"fmt"
	"sync"
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
