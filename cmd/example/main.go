package main

import (
	"fmt"

	"github.com/skeptycal/ansi"
)

func main() {
	// main := ansi.NewColor(7, 32, 40)

	var t = ansi.NewStdout(nil)

	t.CLS()
	t.DevMode(true)
	fmt.Println(t)

	t.UseColor(true)

	t.Hr()

	t.Println("This is bright white text.")
	t.Println("This is bright white text.")

	t.SetColor(ansi.NewColor(ansi.Green, 0, 0))
	fmt.Println(t)

	t.Println("Set UseColor(false).")
	t.UseColor(false)
	t.Println("This is green text. (but isn't showing green - UseColor == false!)")
	t.Br()

	t.Println("Set UseColor(true).")
	t.UseColor(true)
	t.Println("This is green text. (color is back on now - UseColor == true!)")

	t.Println("Any text that is printed after the color is set will print in the default color.")
	t.Println("Set t.SetColor(ansi.NewColor(ansi.Blue, 0, 0)).")
	t.SetColor(ansi.NewColor(ansi.Blue, 0, 0))
	t.Println("Any text that is printed after the color is set will print in the default color.")
	t.SetColor(ansi.NewColor(ansi.Yellow, 0, ansi.Italics))
	t.Println("Any text that is printed after the color is set will print in the default color.")
	t.SetColor(ansi.NewColor(ansi.Black, ansi.RedBackground, ansi.Underline))
	t.Println("Any text that is printed after the color is set will print in the default color.")

	// // fmt.Println(t.Writer)

	// t.Println("This is green text.")
	// t.Println("This is green text.")
	// t.Println("This is green text.")
	// t.Println("This is green text.")
}
