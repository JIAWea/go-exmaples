package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.WriteString("%")
	buffer.WriteString("admin")
	buffer.WriteString("%")

	fmt.Println(buffer.String())

	content := make(map[string]interface{}, 1)
	content["test"] = uint32(1)

	fmt.Printf("data: %+v\n", content)

}
