package main 

import (
	"fmt"
)

type wr interface {
	get() file
	set(name string, v int) file
	setp(name string, v int)
}

type file struct{
	name string
	value int
}

func (f file) get() file{
	return f
}

func (f file) set(name string, v int) file{
	f.name = name
	f.value = v
	fmt.Println("set:",f)
	return f
}

func (f *file) setp(name string, v int) {   //为什么这里用【*file】就报错
	f.name = name
	f.value = v
}

func main () {
	f1 := &file{"yuyang",8}
	var w1 wr
	w1 = f1
	resp := w1.get()
	fmt.Println(resp)
	fmt.Println("调用修改file")
	w1.setp("mz", 6)
	fmt.Println(*f1)

}
