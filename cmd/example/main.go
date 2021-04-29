package main

// ansi example demo

import (
	"fmt"

	"github.com/skeptycal/ansi"
)

var (
	c      = ansi.NewColor // alias for creating new colors
	green  = c(ansi.Green, ansi.Black, ansi.Bold)
	blue   = c(ansi.Blue, ansi.Black, ansi.Bold)
	red    = c(ansi.Red, ansi.Black, ansi.Bold)
	invert = c(172, ansi.Black, ansi.Inverse)
)

func Green(args ...interface{}) {
	fmt.Print(green)
	fmt.Print(args...)
	fmt.Println(ansi.Reset)
}

func Blue(args ...interface{}) {
	fmt.Print(blue)
	fmt.Print(args...)
	fmt.Println(ansi.Reset)
}

func Red(args ...interface{}) {
	fmt.Print(red)
	fmt.Print(args...)
	fmt.Println(ansi.Reset)
}

func main() {
	Green("Green Title")
	Blue("Blue Text")
	Red("Red Text")
	fmt.Printf("%v%s\n", invert, "Orange Inverse Text")
	fmt.Printf("%v%s\n", c(155, 0, 4), "Yellow Italic Text")

}
