package main

import (
	"fmt"
)

func main() {
	t := somenteImpares(soma, []int{50, 1, 51, 25, 34, 22, 11, 60, 33, 24, 19}...)
	fmt.Println(t)
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func somenteImpares(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 != 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}