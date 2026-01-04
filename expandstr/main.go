package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	words := strings.Fields(os.Args)

	if len(words) == 0 {
		fmt.Println()
		return
	}

	fmt.Println(strings.Join(words, "   "))

}
