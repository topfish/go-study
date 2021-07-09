package main 

import (
	"fmt"
)

func main() {
	s7 := make([]string,3,5)
	s7[0] = "yu"
	fmt.Println("s7:", s7)
	fmt.Println("s7-len:", len(s7), "<--->", "s7-cap:", cap(s7))
	s8 := append(s7, "hello")
	fmt.Println("s7增加“hello”,创建s8")
	fmt.Println("s7:", s7, "<--->", "s8:", s8 )
	s7[1] = "hh"
	fmt.Println("修改s7[1]=hh")
	fmt.Println("s7:", s7, "<--->", "s8:", s8 )
	s7[2] = "jj"
	fmt.Println("给s7[2]赋值")
	fmt.Println("s7:", s7, "<--->", "s8:", s8 )
	s7[1] = "opper"
	fmt.Println("修改s7[1]=opper")
	fmt.Println("s7:", s7, "<--->", "s8:", s8 )
	fmt.Println("s8-len:", len(s8), "<--->", "s8-cap:", cap(s8))
	s9 := append(s7,"q","w")
	fmt.Println("s7:", s7, "<--->", "s9:", s9 )
	fmt.Println("s9-len:", len(s9), "<--->", "s9-cap:", cap(s9))
	s10 := append(s7,"a","s","d")	
	fmt.Println("s7:", s7, "<--->", "s10:", s10)
	fmt.Println("s10-len:", len(s10), "<--->", "s10-cap:", cap(s10))
	s7[1] = "000"
	fmt.Println("修改s7[1]=000")
	fmt.Println("s7:", s7, "<--->", "s10:", s10)
}
