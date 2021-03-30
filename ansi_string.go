package ansi

import "fmt"

func NewAnsiString(color Ansi, s string) AnsiString {
	return &ansiString{color, s}
}

// AnsiString implements fmt.Stringer for an Ansi encoded string.
type AnsiString interface {
	String() string
}

type ansiString struct {
	color Ansi
	string
}

func (s ansiString) String() string {
	return fmt.Sprintf("%s%s%s", s.color.String(), s.string, Reset)
}
