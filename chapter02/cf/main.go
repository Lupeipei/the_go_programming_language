package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/Lupeipei/the_go_programming_book/chapter02/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempcov.Fahrenheit(t)
		c := tempcov.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempcov.FToC(f), c, tempcov.CToF(c))
	}
}
