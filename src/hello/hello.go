package main

import (
	"fmt"	
	"util"
)

func main() {
	fmt.Println(util.Reverse("!oG ,olleH"))
	r := []int{41, 24, 76, 11, 45, 64, 21, 69, 19, 36}
	s := "acbfd"
	fmt.Println(s)
	fmt.Println(util.Sort(s))
	fmt.Println(util.BubbleSort(s))
	fmt.Println(util.QuickSort(r, 0, len(r)-1))
}