package main

import (
	pd "examples/protobuf/proto"
	"fmt"

	"github.com/gogo/protobuf/proto"
)

func main() {
	msg := pd.UserInfo{
		Message: "hello",
		Length:  10,
		Cnt:     1,
	}

	encode, err := proto.Marshal(&msg)
	if err != nil {
		fmt.Println("encode err: ", err.Error())
	}

	var decode pd.UserInfo

	if err := proto.Unmarshal(encode, &decode); err != nil {
		fmt.Println("decode err: ", err.Error())
	}

	fmt.Println(decode)
}
