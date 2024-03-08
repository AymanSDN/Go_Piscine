package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
			os.Exit(1)
		}
	} else {
		for _, filename := range os.Args[1:] {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: open %s: %v\n", filename, err)
				continue
			}
			defer file.Close()

			_, err = io.Copy(os.Stdout, file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading from %s: %v\n", filename, err)
				os.Exit(1)
			}
		}
	}
}
