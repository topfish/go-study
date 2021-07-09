package main

import (
	"fmt"
)

func main() {
	f1()
}

func f1() {
	s1 := make([]string,3,5)
	s1[0] = "q"
	s1[1] = "r"
	fmt.Println("f1:", s1)
	f2(s1)
	fmt.Println("f1:", s1)
	
}
 func f2(s []string) {
	a := append(s,"w","e","r")
//	a := s
//	a[1] = "w"
	a[1] = "k"
	fmt.Println("f2:", a)
	fmt.Println("------")
	for index, value := range a {
		fmt.Printf("%v,\t%v,\t%p\n",index, value, &value)
	}
}
