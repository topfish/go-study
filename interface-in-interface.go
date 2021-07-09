package main

import (
	"fmt"
)

type intfA interface {
	get1()
}

type intfB interface {
	intfA
	get2()
}

type stu struct {
}

func (s *stu) get1() {
	fmt.Println("I am get1")
}

func (s *stu) get2() {
	fmt.Println("I am get2")
}

func main() {
	s1 := new(stu)
	var intf1 intfA = s1
	intf1.get1()
	
	var intf2 intfB = s1
	intf2.get2()
}
