//Uppercase and Lowercase in Golang
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "ankit"
	s2 := "ANAND"
	//print in uppercase
	fmt.Println(strings.ToUpper(s))
	//print in lowercase
	fmt.Println(strings.ToLower(s2))
}
