package main

import (
	"fmt"
)

type flyer interface {
	fly()
}

type walker interface {
	walk()
}

type bird struct {
}

func (b bird) walk() {
	fmt.Println("bird walk")
}

func (b bird) fly() {
	fmt.Println("fly")
}

type pig struct {
}

func (p pig) walk() {
	fmt.Println("pig walk")
}

func main() {
	b1 := new(bird)
	fmt.Println(b1)
}
