package main

import (
	"fmt"
)

func main() {

  x := 3

	switch {

	case x == 0:
		fmt.Println("x é 0")

	case x == 1:
		fmt.Println("x é 1")

	default:
		fmt.Println("x é nada")

	}

}
