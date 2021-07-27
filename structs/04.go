package main

import (
	"fmt"
)

func main() {
	anonymousStruct := struct {
		nome       string
		quantidade int
		elementos  [4]string
	}{
		nome:       "exemplo",
		quantidade: 5,
		elementos:  [4]string{"pedra", "água", "lava", "ar"},
	}

	fmt.Println(anonymousStruct)
}
