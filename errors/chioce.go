package main

import (
	"fmt"
)

// MyErrOne 错误1
type MyErrOne struct {
	status int
	err    string
}

// MyErrTwo 错误2
type MyErrTwo struct {
	status int
	err    string
}

func (e *MyErrOne) Error() string {
	return fmt.Sprintf("自定义错误1, err: %s", e.err)
}

func (e *MyErrTwo) Error() string {
	return fmt.Sprintf("自定义错误2, err: %s", e.err)
}

func createErrOne() error {
	return &MyErrOne{400, "超时啦"}
}

func createErrTwo() error {
	return &MyErrTwo{500, "内部错误啦"}
}

func main() {
	err1 := createErrOne()
	err2 := createErrTwo()

	fmt.Println("err1: ", err1)
	fmt.Println("err2: ", err2)

	handlerErr(err1)
	handlerErr(err2)

}

func handlerErr(err error) {

	switch err.(type) {
	case *MyErrOne:
		fmt.Println("自定义错误1") 
	case *MyErrTwo:
		fmt.Println("自定义错误2") 
	case nil:
		fmt.Println("不是错误")
	default:
		fmt.Println("其他错误")
	}

}