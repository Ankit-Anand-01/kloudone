//Funcion in Golang
package main

import "fmt"

func add(x, y int) int {
	z := x + y
	return z
}
func mul(x, y int) int {
	z := x * y
	return z
}
func swap(x, y string) (string, string) {
	//fmt.Println(y, x)
	return y, x
}
func main() {
	x := 5
	y := 8
	res := add(x, y)
	fmt.Println(res)
	a := 2
	b := 4
	r := add(a, b)
	fmt.Println(r)
	num1 := 8
	num2 := 4
	out := mul(num1, num2)
	fmt.Println(out)
	//firstname := ("Ankit")
	//lastname := ("Anand")
	sift, sift1 := swap("Ankit", "Anand")
	fmt.Println(sift, sift1)
}
