package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	now := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		// timeStamp := now.Unix()
		// fmt.Println(timeStamp)
		s += sep + os.Args[i]
		sep = ""
	}
	fmt.Println(s)
	end := time.Now()
	fmt.Println("运行时间", end.Sub(now))

	// for index, Args := range os.Args[1:] {
	// 	//s += index
	// 	fmt.Println(index)
	// 	fmt.Println(Args)
	// }

	now2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], ""))
	end2 := time.Now()
	fmt.Println("运行时间", end2.Sub(now2))

}
