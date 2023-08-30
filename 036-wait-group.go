package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("main routin starts ...")
	// 创建一个 WaitGroup 对象
	var wg sync.WaitGroup

	worker := func(worker_id int) {
		// 在工作协程中要 defer 调用这个 Done 方法，让它去减计数器
		defer wg.Done()

		fmt.Printf("worker %d doing work .\n", worker_id)
		for i := 0; i < 5; i++ {
			fmt.Printf("worker-%d, crurent i = %d \n", worker_id, i)
		}
	}

	// 创建被计数器的值为 2
	wg.Add(2)

	// 创建 2 个工作协程
	go worker(100)
	go worker(200)

	// 等待两个工作协程结束
	wg.Wait()

	fmt.Println("main routin ends ...")
}
