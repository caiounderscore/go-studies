package main

import (
	"fmt"
)

func main() {
	movies := []string{"Matrix", "The Godfather", "Batman", "Mulholland Drive", "Inception"}
	movies = append(movies[:2], movies[3:]...)
	for i := 0; i < len(movies); i++ {
		fmt.Println(movies[i])
	}
}
