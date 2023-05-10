package main

import (
	"fmt"
	"time"
)

// 通过 chan 接收与返回
func worker(message chan string) {
	fmt.Println("worker func start .")
	time.Sleep(time.Second)
	fmt.Println("worker func done .")

	message <- "worker func chan return value ."
}

func main() {
	fmt.Println("main func start .")

	done := make(chan string)
	go worker((done))

	message := <-done
	fmt.Println("got  routine return value: ", message)
	fmt.Println("main func done .")

}
