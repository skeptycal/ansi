package main

import (
	"fmt"

	"github.com/skeptycal/ansi"
)

var green = ansi.NewColor(ansi.Green, ansi.Black, ansi.Bold)

func Green(args ...interface{}) {
	fmt.Print(green)
	fmt.Print(args...)
	fmt.Println(ansi.Reset)
}

func main() {
	Green("Green Title")
}
