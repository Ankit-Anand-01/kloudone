//use of complex numbers
package main

import "fmt"

func main() {
	//declaration of complex numbers
	var x complex64 = complex(6, 2)
	var y complex128 = complex(9, 2)
	fmt.Println(x)
	fmt.Println(y)
	//display the type of x and y
	fmt.Printf("the type of x is %T"+"\n the type of y is %T ", x, y)

}
