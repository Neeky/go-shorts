package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, routine_id int) {
	for {
		select {
		case v := <-ctx.Done():
			fmt.Printf("routine %v got cancel sign : %v , at: %v \n", routine_id, v, time.Now())
			return
		default:
			fmt.Printf("routine is running routine id %d , at: %v \n", routine_id, time.Now())
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	var i int
	for i = 0; i < 2; i++ {
		go worker(ctx, i)
	}

	time.Sleep(time.Second * 4)
	cancel()
	fmt.Println("all workers canceled by main routine .")

	// 留点时间给工作线程终止
	time.Sleep(2 * time.Second)
	fmt.Println("main routine exit .")
}
