package main

import (
	"fmt"
	"time"
)

func sleep(seconds time.Duration) {
	<-time.After(time.Second * seconds)
	return
}

func main() {
	fmt.Println("start")
	sleep(4)
	fmt.Println("stop")
}
