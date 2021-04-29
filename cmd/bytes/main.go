package main

import "github.com/skeptycal/ansi"

func main() {

	fg := "160"
	bg := "191"
	ef := "1"

	ansi.NewColorMake(fg, bg, ef)

	fg = "2"
	bg = "14"
	ef = "4"

	ansi.NewColorMake(fg, bg, ef)

}
