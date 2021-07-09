package main

import (
	"fmt"
)

type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

func Serve(l net.Listener, handler Handler) error

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func HandleFunc() {}

func ListenAndServe() {}

func main() {
	HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
