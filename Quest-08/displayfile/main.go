package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	size := len(args)
	if size == 0 {
		fmt.Printf("File name missing\n")
	}
	if size > 1 {
		fmt.Printf("Too many arguments\n")
	}
	if size == 1 {
		content, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Printf("failed to read file\n")
		}
		fmt.Printf("%v", string(content))
	}
}
