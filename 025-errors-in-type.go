package main

import "fmt"

type ArgError struct {
	// 用作异常的类要实现 Error 方法
	arg     int
	message string
}

func (err *ArgError) Error() string {
	return fmt.Sprintf("{arg:%d, message:%s}", err.arg, err.message)
}

func fa(i int) (int, error) {
	return 0, &ArgError{arg: 0, message: "failed"}
}

func main() {
	res, err := fa(100)
	fmt.Println("res:", res)
	fmt.Println("err:", err)
}
