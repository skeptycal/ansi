package main

// ansistrings demo

import (
	"fmt"

	"github.com/skeptycal/ansi"
)

func main() {

	s := "This is a string."

	fmt.Println(s)

	color := ansi.NewColor("172", "0", "1")

	a := ansi.NewFromAnsi(color, s)

	fmt.Println(a)
}
