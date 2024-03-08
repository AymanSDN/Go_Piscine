package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	size := len(args)
	for i := 0; i < size; i++ {
		if args[i] == "01" || args[i] == "galaxy" || args[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
