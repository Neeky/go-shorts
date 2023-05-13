package main

import "fmt"

func main() {
	// 这里设置了值所以后面能读到
	ch := make(chan string, 1)
	ch <- "hello"

	// 这个 chan 没有设置值，所以后面的 select 就读不到，对应的 case 就不会执行。
	ch_no_value := make(chan string, 1)

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case msg := <-ch_no_value:
		fmt.Println(msg)
	default:
		fmt.Println("defaults")
	}
}
