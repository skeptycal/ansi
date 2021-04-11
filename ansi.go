// Copyright (c) 2020 Michael Treanor
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

const (
	// ResetEffects string = "\033[39;49;0m"  // Reset foreground, background with no effects
	// DefaultText  string = "\033[39;49;22m" // Default text color and intensity
	Default               string = "\033[39;49m" // Default foreground and background color
	Reset                 string = "\033[0m"     // Turn off all effects
	SetInverse            string = "\x1b[7m"
	FmtANSI               string = "\033[%dm"
	CharCommand           string = `⌘`
	CharPlusMinus         string = `±` // Option Shift =
	SymDegree             string = `°` // Option Shift 8
	SymSquared            string = `²` // (no default macOS keymapping)
	SymCubed              string = `³` // (no default macOS keymapping)
	SymUpsideDownQuestion string = `¿` // Option Shift /
	SymCopyright          string = `©` // Option g
	SymRegistered         string = `®` // Option r
	SymTrademark          string = `™` // Option 2
	SymSummation          string = `∑` // Option w
	SymFrontTick          string = `´` // Option e
	SymDager              string = `†` // Option t

	SymEyes        string = `¨` // Option u
	SymHighCarrot  string = `ˆ` // Option i
	SymPhi         string = `ø` // Option o
	SymDiameter    string = `⌀` // (no default macOS keymapping)
	SymSlashO      string = `Ø`
	SymSmallSlashO string = `ø`
	SymEmptySet    string = `∅`
)

var Currency = map[string]string{
	`Yen`:        `¥`, // Option y
	`Currency`:   `¤`,
	`Pound`:      `£`, // Option 3
	`SymEuro`:    `€`, // Option Shift 2
	`SymBitcoin`: `₿`, // (no default macOS keymapping)
	`SymCents`:   `¢`, // Option 4
	`SymPeso`:    `₱`, // (no default macOS keymapping)
}

// var Output = newAnsiStdout()

// func simpleEncode(b byte) string {
// 	return fmt.Sprintf(FmtANSI, b)
// }

// func encode(fg, bg, ef byte) string {
// 	return fmt.Sprintf(ansiEncode, ef, fg, bg)
// }

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
	// Blink
	// FastBlink
	Inverse byte = iota + 2
	Conceal
	Strikeout
	Superscript byte = 73
	Subscript   byte = 74
)

// ANSI escape codes for basic foreground colors
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

// ANSI escape codes for basic background colors
/*
 BlackBackground byte = 40
 RedBackground byte = 41
 GreenBackground byte = 42
 YellowBackground byte = 43
 BlueBackground byte = 44
 MagentaBackground byte = 45
 CyanBackground byte = 46
 WhiteBackground byte = 47
*/
const (
	BlackBackground byte = iota + 40
	RedBackground
	GreenBackground
	YellowBackground
	BlueBackground
	MagentaBackground
	CyanBackground
	WhiteBackground
)

var (
	// Normal is the default color used when no other is specified
	DefaultColor Ansi = NewColor(White, BlackBackground, Normal)
	// Reversed is the default reverse color used when no other is specified
	DefaultReversed Ansi = NewColor(White, BlackBackground, Inverse)
)
