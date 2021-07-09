package main

import(
  "fmt"
)

type str struct{
}

type intf interface{
	get()
}

func (s str) get() {
	fmt.Println("I am get")
}

func main () {
	var s1 str
	s1.get()
//	var intf1 intf = s1
//	intf1.get()
}
