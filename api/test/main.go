package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	ch <- 1
	select {
	case m := <-ch:
		fmt.Println("channel value is", m)
		// fmt.Println("channel value is", <-ch)
	default:
		fmt.Println("channel blocking")
	}
}
