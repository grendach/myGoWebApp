package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func helloWeb(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //parse argument, you have to call it by yourself
	fmt.Println(r.Form) //print form information from server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Dmytro!!\n") //send data to client side
}

func main() {
	http.HandleFunc("/", helloWeb)           //set route to your function
	err := http.ListenAndServe(":9090", nil) //set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
