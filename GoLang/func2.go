//function for multiple operations in golang
package main

import "fmt"

//defining the func
func calc(a, b int) (add, sub, mul int) {
	add = a + b
	sub = a - b
	mul = a * b
	return
}
func main() {
	//calling the func
	out1, out2, out3 := calc(5, 10)
	fmt.Println("add:", out1)
	fmt.Println("sub:", out2)
	fmt.Println("mul:", out3)
	res1, res2, res3 := calc(20, 10)
	fmt.Println("add:", res1)
	fmt.Println("sub:", res2)
	fmt.Println("mul:", res3)

}
