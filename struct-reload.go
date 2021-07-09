package main

import (
	"fmt"
)

type Fly struct{
}

func (f *Fly) fly() {
	fmt.Println("can fly")
}

type Walk struct{
}

func (w *Walk) Walk() {
	fmt.Println("can walk")
}

type Bird struct{
	Fly
	Walk
}

type Human struct{
	Walk
}

func main () {
	h1 := new(Human)
	h1.Walk.Walk()
	b1 := &Bird{}
	b1.fly()
}
