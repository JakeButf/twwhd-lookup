package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: twwhd-lookup <vers>")
		os.Exit(1)
	}

	input := os.Args[1]
	fmt.Printf("Test arg: %s\n", input)
}
