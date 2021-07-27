package main

import (
	"fmt"
)

func main() {

	slice := make([]int, 5, 10)
	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5

	slice = append(slice, 6, 7, 8, 9, 10, 11)

	fmt.Println(slice, len(slice), cap(slice))

}
