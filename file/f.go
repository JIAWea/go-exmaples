package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// dirSize 计算目录下所使用的空间大小
func dirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})

	return size, err
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	fmt.Println(err.Error())
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	dirPath := filepath.Join(".", "root", "in", "admin")
	// ok, err := PathExists(dirPath)
	// fmt.Println(ok, err)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		fmt.Println("err: ", err.Error())
	}
	size, err := dirSize(dirPath)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Println("size: ", size)

	fmt.Println("success")
}
