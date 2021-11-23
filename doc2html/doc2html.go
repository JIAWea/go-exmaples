package main

import (
	"fmt"
	"log"
	"os"

	"code.sajari.com/docconv"
)

func main() {
	f, _ := os.Open(".\\购物合同纠纷.doc")
	res, _, err := docconv.ConvertDoc(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
