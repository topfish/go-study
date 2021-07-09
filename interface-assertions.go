package main

import (
	"fmt"
)

type flyer interface {
	fly()
}

type bird struct{
	name string
}

func (b *bird) fly() {
	fmt.Println("can fly")
}

func (b *bird) walk() {
	fmt.Println("can walk-bird")
}

type walker interface{
	walk()
}

type pig struct{
	name string
}

func (p *pig) walk() {
	fmt.Println("can walk-pig")
}

func main () {
	var w1 walker 
	b1 := &bird{"zhangFeiNiao"}
	w1 = b1
	fmt.Printf("w1:%T--%v\n", w1, w1)
	fmt.Printf("b1:%T--%v\n", b1, b1)
	w2, ok := w1.(*bird)
	fmt.Printf("w2:%T--%v--%v\n", w2, w2, ok)
	w2.walk()
	var f1 flyer 
	f2, ok := f1.(*pig)
	fmt.Printf("f2:%T--%v--%v\n", f2, f2, ok)
}
