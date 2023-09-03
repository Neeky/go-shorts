package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("enter hello func ")

	fmt.Fprintf(w, "hell world")

	log.Printf("exit  hello ")
}

func main() {
	// http.HandleFunc 用于测试处理程序
	http.HandleFunc("/hello", hello)

	// 启动监听
	http.ListenAndServe(":8375", nil)
}
