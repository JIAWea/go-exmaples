package main

import (
	"flag"
	"fmt"
)

func main() {

	// 第一个为flag名称
	// 第二个为默认值
	// 第三个为提示语
	fileName := flag.String("filename", "logfile", "File name for the log file")
	logLevel := flag.Int("loglevel", 0, "An integer value for Level (0-4)")
	isEnable := flag.Bool("enable", false, "A boolean value for enabling log options")

	var num int
	// 判定变量
	flag.IntVar(&num, "num", 25, "An integer value")

	flag.Usage = usage
	// Parse parses flag definitions from the argument list.
	flag.Parse()

	// Get the values from pointers
	fmt.Println("filename:", *fileName)
	fmt.Println("loglevel:", *logLevel)
	fmt.Println("enable:", *isEnable)

	// Get the value from a variable
	fmt.Println("num:", num)

	// 返回没有定义的flag
	args := flag.Args()
	// fmt.Printf("args type: %T\n", args)  // []string
	if len(args) > 0 {
		fmt.Println("没有申明的flag名称:")
		// Print the arguments
		for _, v := range args {
			fmt.Println(v)
		}
	}

}

func usage() {
	fmt.Printf("start fail...\n")
	flag.PrintDefaults()
}
