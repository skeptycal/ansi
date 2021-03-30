// Copyright (c) 2020 Michael Treanor
// MIT License

// Package ansi provides fast ansi escape sequence processing based on strings.Builder.
// The standard is defined by the ECMA-48 standard "Control Functions for Coded Character Sets - Fifth Edition"
package ansi

import (
	"fmt"

	"github.com/skeptycal/util/stringutils"
)

const (
	// Clear screen
	ansiCLS = "\033[2J"
	// Character used for HR function
	HrChar string = "="
)

func ByteToNum(c byte) byte {
	if stringutils.IsDigit(c) {
		return c - 65
	}
	return '0'
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
