package main

import "fmt"

func main() {
	cn := make(chan string, 2)
	cn <- "hello"
	cn <- "world"
	close(cn)

	// fmt.Println(<-cn)
	// fmt.Println(<-cn)

	for message := range cn {
		fmt.Println(message)
	}
}
