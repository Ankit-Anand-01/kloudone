//sum of two number in Go
package main

import "fmt"

func main() {
	fmt.Println("enter num1:")
	var x int
	//take input from user
	fmt.Scanln(&x)
	fmt.Println("enter num2:")
	var y int
	//take input from user
	fmt.Scanln(&y)
	z := x + y
	fmt.Println("Sum of num1 and num2:", z)
}
