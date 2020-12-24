package main

import "fmt"

func main() {
	x := 200
	if x < 100 {
		//executes this block if condition is true
		fmt.Printf("x is less than 100\n")
	} else {
		//executes this block if  above condition is false
		fmt.Printf("x is greater than 100\n")
	}
}
