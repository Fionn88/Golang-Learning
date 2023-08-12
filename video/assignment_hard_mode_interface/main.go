package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// dat, err := os.ReadFile(os.Args[1])
	// check(err)
	// fmt.Println(string(dat))

	f, err := os.Open(os.Args[1])
	check(err)
	io.Copy(os.Stdout, f)

}

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(1)
	}
}
