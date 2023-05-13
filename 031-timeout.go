package main

import (
	"fmt"
	"time"
)

func main() {
	foo := make(chan string, 1)
	go func() {
		// 如果这里延时 2s 就会触发 select 的 timeout 逻辑
		// time.Sleep(2 * time.Second)
		foo <- "hello world"
	}()

	select {
	case msg := <-foo:
		fmt.Println("this is foo function returned return value =", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("time out logic .")

	}
}
