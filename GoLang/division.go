//Division in Go
package main

import "fmt"

func main() {
	fmt.Println("enter num:")
	var x int
	//take input from user
	fmt.Scanln(&x)
	fmt.Println("divided by:")
	var y int
	//take input from user
	fmt.Scanln(&y)
	z := x / y
	fmt.Println("Division of num1 and num2 =", z)
	a := x % y
	fmt.Println("remainder = ", a)
}
