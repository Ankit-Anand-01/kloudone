package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{59, 3, 7, 5, 6, 44, 12, 53, 8}
	//check the slice is sorted or not
	s := sort.IntsAreSorted(a)
	//print false because slice is not sortd
	fmt.Println(s)
	//sorting in ascending order by using inbuilt method
	sort.Ints(a)
	su := sort.IntsAreSorted(a)
	fmt.Println(a)
	//print true because slice is sorted now
	fmt.Println(su)

}
