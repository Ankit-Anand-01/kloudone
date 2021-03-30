package main

import "fmt"

func rd(a []int) []int {
	result := []int{}
	n := len(a)
	for i := 0; i < n; i++ {
		j := 0
		for ; j < i; j++ {
			if a[i] == a[j] {
				break
			}
		}
		if j == i {
			result = append(result, a[i])

		}
	}
	return result

}
func main() {
	s := []int{5, 6, 3, 2, 1, 4, 5, 6, 3, 2, 1}
	fmt.Println(rd(s))
}
