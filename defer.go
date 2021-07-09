package main

import (
	"fmt"
)

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
}

func f1() (t int) {
	defer func() {
	t = 2
}()
	return 0
}

func f2() (t int) {
	t = 5
	defer func() {
	t=t+1
}()
	return t
}
 func f3() (r int) {
	t := 5
	defer func() {
	t = t+1
}()
	return t
}
