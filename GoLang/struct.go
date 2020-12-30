//structure in Go
package main

import "fmt"

//defining a struct type
//employee structure
type employee struct {
	name       string
	age        int
	city       string
	state      string
	experience string
	contactNo  int
}

func main() {
	// Naming fields while
	// initializing a struct
	emp1 := employee{
		name:       "Dhruv",
		age:        25,
		city:       "Jaipur",
		state:      "Rajasthan",
		experience: "2-years",
		contactNo:  881234536799,
	}
	//without using fields name
	emp2 := employee{"Animesh", 32, "Pune", "Maharastra", "8-years", 9829210018}
	emp3 := employee{"Harry", 28, "Roorkee", "Uttarakhand", "3-years", 7856432100}
	fmt.Println("employee 1:", emp1)
	fmt.Println("employee 2:", emp2)
	fmt.Println("employee 3:", emp3)

}
