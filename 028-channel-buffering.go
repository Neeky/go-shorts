package main

import "fmt"

func main() {
	// buffer 的大小为 2
	messages := make(chan string, 2)

	messages <- "hello"
	messages <- "world"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
