package ansi

import "fmt"

// NewAnsiString returns a new AnsiString that has
// a Stringer interface that wraps it in an ansi color.
func NewAnsiString(color Ansi, s string) AnsiString {
	return &ansiString{color, s}
}

// AnsiString implements a Stringer interface that wraps
// a string in an ansi color format.
//
// The Stringer prefixes the string in an Ansi color code
// set specified and the ansi Reset ansi code is place at
// the end.
type AnsiString interface {
	String() string
}

type ansiString struct {
	color Ansi
	string
}

func (s *ansiString) String() string {
	return fmt.Sprintf("%s%s%s", s.color.String(), s.string, Reset)
}
