package main

import (
	"fmt"
)

func main() {
	map1 := make(map[int]string)
	map1[1] = "hello"
	fmt.Printf("%T--%v\n", map1[1], map1[1])
}
