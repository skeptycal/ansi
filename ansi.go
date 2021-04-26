// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

// Package ansi provides fast ansi escape sequence processing based on strings.Builder.
// The standard is defined by the ECMA-48 standard "Control Functions for Coded Character Sets - Fifth Edition"
//
// The default zero values for Go variables are used whenever possible.
// Reference: https://golangbyexample.com/go-default-zero-value-all-types/
//
// ANSI escape codes for Command Line applications are well defined.
// Reference: https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html
//
// Many of the ANSI cursor codes are from https://github.com/k0kubun/go-ansi
package ansi

import (
	"fmt"
	"strings"
)

const (
	// set ANSI 8-bit (256 color) foreground and background color with leading effect
	ansiEncode = "\x1b[%d;38;5;%d;48;5;%dm"
	fgEncode   = "\x1b[%d;38;5;%d;48;5;%dm%s%s"
)

const (
	// ResetEffects string = "\x1b[39;49;0m"  // Reset foreground, background with no effects
	// DefaultText  string = "\x1b[39;49;22m" // Default text color and intensity
	Default                 string = "\x1b[39;49m" // Default foreground and background color
	Reset                   string = "\x1b[0m"     // Turn off all effects
	SetInverse              string = "\x1b[7m"
	DefaultEffectString     string = "\x1b[1m"
	DefaultForegroundString string = "\x1b[37m"
	DefaultBackgroundString string = "\x1b[47m"
	FmtANSI                 string = "\x1b[%dm"
)

var (
	bDefault                 []byte = []byte(Default)
	bReset                   []byte = []byte(Reset)
	bSetInverse              []byte = []byte(SetInverse)
	bDefaultEffectString     []byte = []byte(DefaultEffectString)
	bDefaultForegroundString []byte = []byte(DefaultForegroundString)
	bDefaultBackgroundString []byte = []byte(DefaultBackgroundString)
	bFmtANSI                 []byte = []byte(FmtANSI)
	ResetBytes               []byte = []byte(Reset)

	CurrentColor     Ansi   = NewColor(214, 0, 1)
	DefaultColor     string = NewColor(White, Black, Normal).String()
	bDefaultColor    []byte = []byte(DefaultColor)
	DefaultReversed  string = NewColor(White, Black, Inverse).String()
	bDefaultReversed []byte = []byte(DefaultColor)
)

// TODO - for dev tests ...
var (
	BDefault                 = bDefault
	BReset                   = bReset
	BSetInverse              = bSetInverse
	BDefaultEffectString     = bDefaultEffectString
	BDefaultForegroundString = bDefaultForegroundString
	BDefaultBackgroundString = bDefaultBackgroundString
	BFmtANSI                 = bFmtANSI
	BesetBytes               = ResetBytes
)

// var Output = newAnsiStdout()

// ANSI escape codes for text effects
//
// These are the most commonly used. ANSI codes above 9 are very
// rare and many are not fully implemented.
//  Normal      = 0
//  Bold        = 1
//  Faint       = 2
//  Italics     = 3
//  Underline   = 4
//  Inverse     = 7
//  Conceal     = 8
//  Strikeout   = 9
//
const (
	Normal byte = iota
	Bold        // bold or increased intensity
	Faint       // faint, decreased intensity or second color
	Italics
	Underline
	// Blink  		// not implemented
	// FastBlink 	// not implemented
	Inverse byte = iota + 2
	Conceal
	Strikeout
	Superscript byte = 73
	Subscript   byte = 74
)

// ANSI escape codes for basic ANSI colors
const (
	Black byte = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

func Rainbow(start, end byte, s string) string {
	return rainbow1(start, end, s)
}

func rainbow1(start, end byte, s string) string {
	sb := strings.Builder{}
	defer sb.Reset()

	colorRange := end - start
	strLength := len(s)

	var slope float64 = float64(colorRange) / float64(strLength)

	color := func(i int) byte {
		return byte(slope*float64(i) + float64(start))
	}

	for i := 0; i < strLength; i++ {
		c := NewColor(color(i), 0, 0)
		sb.WriteString(fmt.Sprintf("%s%s", c, string(s[i])))
	}
	sb.WriteString(Reset)
	return sb.String()
}

// constant color set
func rainbow2(start, end byte, s string) string {
	sb := strings.Builder{}
	defer sb.Reset()

	colorRange := end - start
	strLength := len(s)

	var slope float64 = float64(colorRange) / float64(strLength)

	color := func(i int) byte {
		return byte(slope*float64(i) + float64(start))
	}

	for i := 0; i < strLength; i++ {
		c := NewColor(color(i), 0, 0)
		sb.WriteString(fmt.Sprintf("%s%s", c, string(s[i])))
	}
	sb.WriteString(Reset)
	return sb.String()
}
