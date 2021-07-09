package main

import (
	"fmt"
)

func main() {
	fmt.Printf("type:v---->%v\n", f1())
	fmt.Printf("type:T---->%T\n", f1())
	fmt.Printf("type:v---->%v\n", f2())
	fmt.Printf("type:T---->%T\n", f2())
}

func f1()int {
	t := 2
	return t
}
 
func f2() string {
	a := "hello"
	return a
}
