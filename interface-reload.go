package main

import (
	"fmt"
)
//定义只读接口
type read interface{
	get() file
}
//实现只读接口
func (f file) get() file{
	return f
}
//定义只写接口
type write interface{
	set(v int)
}
//实现只写接口
func (s *socket) set(v int) {
	s.version = v
}
//定义读写接口
type readwrite interface{
	read
	write
}
//定义file类
type file struct {
	name string
}
//定义socket类
type socket struct{
	version int
}
//定义读写类
type wr struct{
	file
	socket
}
func main() {
//实例化file和socket并打印值
	f1 := file{"k8s"}
	s1 := &socket{1}
	fmt.Println(f1,s1)
//实例化只写接口，并测试写入
	var w1 write = s1
	w1.set(2)
	fmt.Println(f1,s1)
//实例化读写类，并赋值。读写类赋值给读写接口，调用set方法，并验证结果。	
	wr1 := wr{f1, *s1}
	var readw1 readwrite = &wr1
	readw1.set(3)
	fmt.Println(wr1)
}
