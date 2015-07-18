package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	io.WriteString(w, "Hello World")
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	port := 8000
	server := http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: &myHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = hello

	fmt.Println("Now listening to port ", port)
	server.ListenAndServe()

}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "My server: "+r.URL.String())
}
