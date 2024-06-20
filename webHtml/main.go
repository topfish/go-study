package main

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("hello")
	user := User{
		Name:   "yuyang",
		Gender: "man",
		Age:    33,
	}

	tmp, err := template.ParseFiles("myTemplates/hello.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmp.Execute(w, user)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("http server start failed ", err)
	}
}
