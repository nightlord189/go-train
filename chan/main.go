package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)
	fmt.Println("channel closed")
	reader(ch1)
}

func reader(ch chan int) {
	fmt.Println("start reading")
	for {
		val, ok := <-ch
		if ok {
			fmt.Println(val)
		} else {
			break
		}
	}
}
