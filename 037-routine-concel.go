package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建 channel 对象
	var can_stop chan bool
	can_stop = make(chan bool)

	go func() {
		for {
			select {
			case <-can_stop:
				fmt.Printf("main routine want worker routine stop, we well stop .\n")
				return
			default:
				fmt.Printf("worker routine is running . %v \n", time.Now())
				time.Sleep(time.Second * 1)
			}
		}
	}()

	time.Sleep(time.Second * 4)
	fmt.Printf("main routine is wakup .")
	can_stop <- true
	close(can_stop)

}
