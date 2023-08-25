package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func Hello() {
	// 如果是第一次执行
	now := time.Now()
	once.Do(func() {
		fmt.Printf("第一次执行 Hello 函数\n")
	})

	// 每次执行的时候都会执行这里
	fmt.Printf("%v Hello 函数的主逻辑 \n", now)
}

func main() {
	// 每 1 条 执行一次 Hello 函数
	for range time.Tick(1 * time.Second) {
		go Hello()
	}
}
