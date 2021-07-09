package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	fmt.Println(ch1)
	go gg(ch1)
	b := <-ch1
	fmt.Println(b)
}

func gg(ch chan int) {
        time.Sleep(3 * time.Second)
	ch <- 1
}
