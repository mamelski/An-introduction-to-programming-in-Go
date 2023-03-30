package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	fileBytes := make([]byte, stat.Size())
	_, err = file.Read(fileBytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	str := string(fileBytes)
	fmt.Println(str)
}
