package main

import (
	"fmt"
)

func main() {

	startyear := 1996
	endyear := 2020
	for {
		if startyear > endyear {
			break
		}
		fmt.Println(startyear)
		startyear++
	}

}
