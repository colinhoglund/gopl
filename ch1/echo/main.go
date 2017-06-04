package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Exercise 1
	// var out, sep string
	// if len(os.Args) > 1 {
	// 	for _, v := range os.Args[1:] {
	// 		out += sep + v
	// 		sep = " "
	// 	}
	// }
	// fmt.Println(out)

	// Exercise 1.1
	if len(os.Args) > 1 {
		fmt.Println(strings.Join(os.Args[1:], " "))
	}

	// Exercise 1.2
	// for i, v := range os.Args {
	// 	fmt.Println(i, v)
	// }
}
