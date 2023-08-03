package main

// Github Push Code Test
import (
	"fmt"
)

func main() {
	// ======= Declare map =======
	// var colors map[string]

	// ======= Same Way To Declare =======
	colors := make(map[int]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	colors[10] = "#ffffff"

	// ======= Delete Map By Key =======
	delete(colors, 10)

	fmt.Println(colors)

}
