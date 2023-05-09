package main

import (
	"fmt"
	"time"
)

func foo(message string) {
	fmt.Println(message)
	time.Sleep(time.Second)
}

func main() {
	// 直接把 5 个协程放后台执行，每一个大概耗时 1s , 5 个总耗时也就是 1s 多一点点
	for i := 0; i < 5; i++ {
		go foo(fmt.Sprintf("message:%s, index:%d", "foo", i))
	}

	// 直接把 5 个协程放后台执行，每一个大概耗时 1s , 5 个总耗时也就是 1s 多一点点
	for i := 0; i < 5; i++ {
		go foo(fmt.Sprintf("message:%s, index:%d", "bar", i))
	}

	// 主线程 sleep 1 ，有可能遇到程序退出了，但是后台的协程没有运行完的情况如：
	time.Sleep(time.Second)
}
