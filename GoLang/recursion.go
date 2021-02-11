//finding factorial using recursion in golang
package main

import "fmt"

//defining a func for factorial
func factorial(i int) int {

	if i <= 1 {
		return 1
	}
	//calling the recursive func
	return i * factorial(i-1)
}
func main() {
	i := 5
	fmt.Println(factorial(i))
	j := 4
	fmt.Println(factorial(j))
	fmt.Println(factorial(3))
	fmt.Println(factorial(2))
	fmt.Println(factorial(1))
}
