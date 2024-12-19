package main

import (
	"fmt"
	"net/http"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello，我是Golang服务消息")
	if err != nil {
		return
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Golang后台服务已启动")
	if err != nil {
		return
	}
}

func main() {
	var exposePort = ":9000"
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/hello", httpHandler)
	err := http.ListenAndServe(exposePort, nil)
	if err != nil {
		return
	}
}
