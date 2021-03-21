//maps in Golang
package main

import "fmt"

func main() {
	//initializing map
	m := make(map[string]int)

	//keys and values
	m["Rahul"] = 45
	m["Dheeraj"] = 41
	m["Aman"] = 35
	m["Riyaz"] = 41
	m["Aayat"] = 48
	fmt.Println(m)
	//printing value of a particular key
	fmt.Println(m["Aayat"])
	//length of map
	fmt.Println(len(m))
	//another method to initialize map with keys and values
	m1 := map[int]string{1: "Rahul", 2: "Riyaz", 3: "Aman", 4: "Dheeraj", 5: "Aayat"}
	fmt.Println(m1)
	//using built-in delete function to delete a key
	delete(m, "Dheeraj")
	fmt.Println(m)
}
