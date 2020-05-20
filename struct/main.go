package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Person struct
type Person struct {
	AgeA int
	AgeB *int
	
	CreateAT time.Time
	DeleteAT *time.Time
}

func main(){
	// struct 中带指针是允许为空，nil(null)，没带指针则是零值
	var p Person
	text := `{}`
	if err := json.Unmarshal([]byte(text), &p); err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("反序列化: ", p)

	a := Person{}
	fmt.Println("初始值:", a)
	b, _ := json.Marshal(a)
	fmt.Println("json:", string(b))
}