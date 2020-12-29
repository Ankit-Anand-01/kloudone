//go program of if..else..if
package main

import "fmt"

func main() {
	x := 500
	//checking the condition
	if x == 100 {
		// if condition is true then
		// display the following */
		fmt.Printf("value of x is 100\n")
	} else if x == 200 {
		fmt.Printf("value of x is 200\n")
	} else if x == 300 {
		fmt.Printf("value of x is 300\n")
	} else if x == 400 {
		fmt.Printf("value of x is 400\n")
	} else if x == 500 {
		fmt.Printf("value of x is 500\n")
	} else {
		//if none of the condions is true
		// display the following
		fmt.Printf("none of the values is matching\n")
	}
}
