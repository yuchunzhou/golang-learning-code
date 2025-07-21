package main

import (
	"fmt"
	"net/http"
)

type MyHttpHandler struct{}

func (handler MyHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("path: %v\n", r.RequestURI)
	fmt.Printf("args: %v\n", r.URL.Query())
	fmt.Printf("headers: %v\n", r.Header)
	_, _ = fmt.Fprint(w, "hello world!\n")
}

func main() {
	err := http.ListenAndServe(":8080", MyHttpHandler{})
	if err != nil {
		fmt.Printf("start http service failed: %s\n", err.Error())
		return
	}
}
