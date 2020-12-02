package main

import (
	"fmt"
)

func main() {

	movietype := "drama"

	switch movietype {

	case "action":
		fmt.Println("Action Movies: Matrix, Baby Driver...")

	case "adventure":
		fmt.Println("Adventure Movies: Robin Hood, Thor...")

	case "drama":
		fmt.Println("Drama Movies: Knives out, Parasite...")

	}

}
