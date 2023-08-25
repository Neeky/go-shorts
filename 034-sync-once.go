package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
	 * 函数只执行一次，执行完成之后其它协程里的代码算是调用了，也不会得到真正的执行
	 */
	var once sync.Once

	// 定义只执行一次的函数
	onceBody := func() {
		fmt.Println("Only once")
	}

	done := make(chan bool)
	for i := 0; i < 5; i++ {
		j := i
		go func(int) {
			// 只会在第一次执行时有效
			once.Do(onceBody)
			fmt.Println(j)
			done <- true
		}(j)
	}
	<-done

	time.Sleep(2 * time.Second)
}
