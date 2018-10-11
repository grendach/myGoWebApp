package main

import (
	"fmt"
	"net/http"
)

type dmytro int

func (d dmytro) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "HEELOOO")
}

func main() {
	var dupa dmytro
	http.ListenAndServe(":8080", dupa)

}
