//Channels in Golang
package main

import (
	"fmt"
	"time"
)
//defining func to print
func display(ch chan string) {
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
func main() {
	//initializing the channel 
	ch := make(chan string)
       //calling goroutine
	go display(ch)
       //sending values to the channels
	ch <- "Ping"
	ch <- "Pong"
	time.Sleep(100 * time.Millisecond)
}
