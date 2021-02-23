//gorotine in Golang
package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for w := 0; w < 6; w++ {
		time.Sleep(5 * time.Second)
		fmt.Println(str)
	}
}

func main() {

	// using go,we can call the go routine
	go display("Hi")

	display("Ankit")
	//go display("Anand")
}
