package main

import (
	"fmt"
)

type user struct{
	name string
	email string
}

func (u user) notify() {
	fmt.Println(u.name,u.email)
}

func (u user) changev(newname string) {
	u.name = newname
	fmt.Println(u.name)
}

func (u *user) change(email string) {
	u.email = email
}

func main() {
	bill := user{
		name: "bill",
		email: "bill@email.com",
	}
	bill.notify()

	lisa := &user{"lisa", "lisa@email.com"}
	lisa.notify()	
	fmt.Println("---------")
	bill.change("bill@newemail.com")
	bill.notify()
	lisa.change("lisa@newemail.com")
	lisa.notify()	
	fmt.Println("---------")
	fmt.Println(bill)
	fmt.Println(lisa)
	fmt.Println("---------")
	bill.changev("newbill")	
	bill.notify()
	fmt.Println("---------")
	lisa.changev("newlisa")
	lisa.notify()
}


