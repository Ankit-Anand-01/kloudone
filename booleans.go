//use of boolean datatype
package main

import "fmt"

func main() {
	//string variables
	str1 := "My Name Is Ankit Anand"
	str2 := "my name is Ankit Anand"
	str3 := "My Name Is Ankit Anand"
	result1 := str1 == str2
	result2 := str1 == str3
	//display the result
	fmt.Println(result1)
	fmt.Println(result2)
	//display the type of result1 and result2
	fmt.Printf("The type of result1 is %T"+"\n The type of result2 is %T", result1, result2)
}
