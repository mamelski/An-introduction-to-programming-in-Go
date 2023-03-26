package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("test", "te"))

	fmt.Println(strings.Count("test", "t"))

	fmt.Println(strings.HasPrefix("test", "tes"))

	fmt.Println(strings.HasSuffix("test", "est"))

	fmt.Println(strings.Index("test", "t"))

	fmt.Println(strings.Join(strings.Split("test", ""), " "))

	fmt.Println(strings.Repeat("test ", 6))

	fmt.Println(strings.Replace("aaaaa", "a", "b", 3))

	fmt.Println(strings.Split("a-b-c-d-e", "-"))

	fmt.Println(strings.ToLower("TEST"))
	fmt.Println(strings.ToUpper("test"))

	arr := []byte("test")
	str := string(arr)

	fmt.Println(str)

}
