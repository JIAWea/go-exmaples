package main

import (
	"fmt"
)

// Notify interface
type Notify interface {
	sendMsg()
}

// User struct
type User struct {
	Name    string
	Age     int
	IsAdmin bool
}

// sendMsg func
func (u *User) sendMsg() {
	fmt.Printf("send email to user(%s), age(%v), is admin(%t)\n", u.Name, u.Age, u.IsAdmin)
}

// Admin struct
type Admin struct {
	Name    string
	Age     int
	IsAdmin bool
}

// sendMsg func
func (a *Admin) sendMsg() {
	fmt.Printf("send email to admin(%s), age(%v), is admin(%t)\n", a.Name, a.Age, a.IsAdmin)
}

func main() {
	// 通过接口实现多态
	//     当非接口类型(T)的一个值(t)被包裹了接口类型(I)的一个接口值(i)，并且已实现了接口(I)中的方法，
	// 则称类型T实现了I接口。
	//     换句话说，调用一个接口值的方法实际上将调用此接口值的动态值的对应方法，即i.m()被调用相当于t.m()。
	// 一个接口值可以包裹不同的动态类型的动态值来表现出各种不同的行为，这就是多态。

	var n Notify

	u := User{
		Name:    "Ray",
		Age:     18,
		IsAdmin: false,
	}

	n = &u
	n.sendMsg()

	a := Admin{
		Name:    "Admin",
		Age:     18,
		IsAdmin: true,
	}

	n = &a
	n.sendMsg()

}
