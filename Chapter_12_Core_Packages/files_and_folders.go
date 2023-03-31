package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile1() {
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

func ReadFile2() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(bs)
	fmt.Println(str)
}

func CreateFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString("I created this")
}

func ReadDirectories() {
	dirr, err := os.Open(".")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dirr.Close()

	fileInfos, err := dirr.ReadDir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

}

func main() {
	CreateFile()
	ReadFile2()
	ReadDirectories()
}
