package main

import (
	"fmt"

	"github.com/skeptycal/ansi"
)

func main() {
	// main := ansi.NewColor(7, 32, 40)

	var t = ansi.NewCLI(nil)

	t.CLS()
	// t.Hr()

	fmt.Println(t)

	// t.Println("This is bright white text.")
	// t.Println("This is bright white text.")

	// t.Fg(ansi.Green)
	// fmt.Println("test")

	// // fmt.Println(t.Writer)

	// t.Println("This is green text.")
	// t.Println("This is green text.")
	// t.Println("This is green text.")
	// t.Println("This is green text.")
}
