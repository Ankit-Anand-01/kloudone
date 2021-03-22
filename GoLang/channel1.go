package main

import (
	"fmt"
	"time"
)

func display(ch chan string) {
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
func main() {
	ch := make(chan string)

	go display(ch)

	ch <- "Ping"
	ch <- "Pong"
	time.Sleep(100 * time.Millisecond)
}
