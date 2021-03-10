//defer in Golang
package main

import "fmt"

func main() {
	//last printing
	defer fmt.Println("Anand")
	//second printing
	defer fmt.Println("Ankit")
	//first printing
	fmt.Println("GOOD MORNING!")
}
