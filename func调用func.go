package main

import (
	"fmt"
)

func f1() {
	fmt.Println("I am F1")
}

func f2(f func()) {
	fmt.Println("say your name :", f)
	f()
	
}
func main() {
	f2(f1)
//	var f3 func()
//	fmt.Printf("%T\n", f3)
}
