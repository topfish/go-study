package main

import (
	"fmt"
)

func main() {
	user()
	person()
}

func person() {
	type person struct {
		name	string
		age	int
	}
	var p1 person
	fmt.Println(p1.name,p1.age)
	fmt.Println("---")
}

func user () {
	type user struct {
		name	string
		age	int
		priv	bool
	}
	type user1 struct {
		name	string
		age	int
		priv	bool
	}
	yy := user{
		name: "yuyang",
		age: 30,
		priv: true,
	}
	fmt.Println(yy)
	type ff struct{
		fri user
		name string
		age int
	}
	f1 := ff{
		fri: user{
			name: "danny",
			age: 9,
			priv: false,
		},
		name: "kity",
		age: 2,
	}
	fmt.Println(f1)
	var u2 user
	u2 = user{"tony", 5, true}
	fmt.Println("u2:",u2)
}
