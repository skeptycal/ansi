// Copyright (c) 2020 Michael Treanor
// MIT License

// Package ansi provides fast ansi escape sequence processing based on strings.Builder.
// The standard is defined by the ECMA-48 standard "Control Functions for Coded Character Sets - Fifth Edition"
package ansi

import (
	"fmt"
	"strings"
)

const (
	ansiCLS            string = "\033[2J" // ANSI escape code - clear screen
	hrChar             string = "="
	defaultScreenWidth int    = 80
)

var (
	// Normal is the default color used when no other is specified
	DefaultColor Ansi = NewColor(White, BlackBackground, Normal)
	// Reversed is the default reverse color used when no other is specified
	DefaultReversed Ansi = NewColor(Black, WhiteBackground, Inverse)
)

func screenWidth() int {
	// todo - add support for variable width
	return defaultScreenWidth
}

func LineBreak(c string) {
	if c == "" {
		c = hrChar
	}
	fmt.Println(strings.Repeat(c, screenWidth()))
}

func Hr() {
	LineBreak(hrChar)
}

func Br() {
	fmt.Println("")
}

// Print wraps args in an ANSI 8-bit color (256 color codes)
func Print(i Ansi, args ...interface{}) {
	fmt.Print(i.String())
	fmt.Print(args...)
	fmt.Print(Reset)
}

// Println wraps args in an ANSI 8-bit color (256 color codes)
// and adds a newline character
func Println(i Ansi, args ...interface{}) {
	Print(i, args...)
	fmt.Println("")
}

func CLS() {
	fmt.Print(ansiCLS)
}
