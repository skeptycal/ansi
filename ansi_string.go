package ansi

import "fmt"

// NewAnsiString returns a new AnsiString that wrapped
// in an Ansi color.
func NewAnsiString(color Ansi, str string) AnsiString {
	return &ansiString{
		s: fmt.Sprintf("%s%s%s", color, str, Reset),
	}
}

// AnsiString implements a Stringer interface that wraps
// a string in an ansi color format.
//
// The Stringer wraps a string in Ansi color
// It adds a prefix for the Ansi code set (fg,bg,effect)
// and a suffix for the ANSI Reset code.
type AnsiString interface {
	String() string
}

type ansiString struct {
	s string
}

func (s *ansiString) String() string {
	return s.s
}
