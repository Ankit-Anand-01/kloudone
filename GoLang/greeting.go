package main

import "fmt"

func greet(s string) {
	fmt.Printf("Good Morning! %s\n", s)
}
func wish(s string) {
	fmt.Printf("wish you a very happy birthday! %s", s)
}

func main() {
	fmt.Println("enter your name:")
	var a string
	fmt.Scanln(&a)
	greet(a)
	wish(a)
	//greet("Ankit")
	//greet("Jeevan")
}
