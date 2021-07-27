package main

import (
	"fmt"
)

func main() {
	array := [5]int{0, 1, 2, 3, 4}
	fmt.Printf("%T - %v\n", array, array)
	slice := []int{0, 1, 2, 3, 4}
	fmt.Printf("%T - %v\n", slice, slice)
	slice = append(slice, 5)
	fmt.Println(slice)
}
