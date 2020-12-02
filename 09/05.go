package main

import (
	"fmt"
)

func main() {

	movies := map[string]string{
		"Matrix":           "action",
		"The Godfather":    "crime",
		"Inception":        "action",
		"Mulholland Drive": "thriller",
		//    "Batman": "crime",
	}

  delete(movies, "Matrix")

	fmt.Println(movies)
	if batman, ok := movies["Batman"]; !ok {
		fmt.Println("Not exist! ok:", ok)
	} else {
		fmt.Println(batman)
	}

}
