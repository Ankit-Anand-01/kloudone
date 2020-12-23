//use of integers
package main

import (
	"fmt"
)

func main() {
	//using 32bit signed integer
	var A int32 = 36588
	fmt.Println(A+1, A-1, A*2, A/5)
	//using  16bit unsigned integer
	var B uint16 = 222
	fmt.Println(B+1, B-1, B*2, B/5)
}
