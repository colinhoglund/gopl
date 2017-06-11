package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/colinhoglund/gopl/ch2/tempconv"
)

func printUnits(arg string) {
	floatArg, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s is not a number\n", arg)
	} else {
		celsius := tempconv.Celsius(floatArg)
		fmt.Println(celsius)
		fmt.Println(tempconv.CToF(celsius))
		fmt.Println(tempconv.CToK(celsius))
	}
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			printUnits(arg)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			printUnits(scanner.Text())
		}
	}
}
