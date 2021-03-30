//remove duplicates using nested loops in Golang
package main

import "fmt"

func removeDuplicates(a []int) []int {
	result := []int{}
	n := len(a)
	//traversing array
	for i := 0; i < n; i++ {
		j := 0
		//checking the element has come before or not
		for ; j < i; j++ {
			//if element has come before
			if a[i] == a[j] {
				break
			}
		}
		//element has not come before
		if j == i {
			result = append(result, a[i])

		}
	}
	return result

}
func main() {
	s := []int{5, 6, 3, 2, 1, 4, 5, 6, 3, 2, 1}
	fmt.Println(removeDuplicates(s))
}
