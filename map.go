package main 

import (
	"fmt"
)

func main() {
	map4 := map[int]string{1:"hello", 2:"kitty"}

	fmt.Println(map4)
	v1 := map4[1]
	v2 := map4[2]
	map4[3] = "yuyang"
	fmt.Printf("%s,%s\n",v1,v2)
 	fmt.Println("v1: ", v1, "v1-addr: ", &v1)
	for index,value := range map4 {
		fmt.Println(index,value)
	}
	delete(map4, 2)
	fmt.Println("------------")
	for index,value := range map4 {
		fmt.Println(index,value)
	}
}
