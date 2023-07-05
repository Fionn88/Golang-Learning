package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args {
		if idx == 0 {
			fmt.Println("os.Args[0]=", arg)
		} else {
			fmt.Println(idx, "=", arg)
		}
	}
}
