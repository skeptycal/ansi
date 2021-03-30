package ansi

import "fmt"

const (
	// set ANSI 8-bit (256 color) foreground and background color with leading effect
	ansiEncode = "\033[%d;38;5;%d;48;5;%dm"
)

type Ansi interface {
	String() string
}

func NewColor(f, b, e byte) Ansi {
	return &ansiCode{f, b, e}
}

type ansiCode struct {
	effect     byte
	foreground byte
	background byte
}

func (c *ansiCode) String() string {
	return fmt.Sprintf(ansiEncode, c.effect, c.foreground, c.background)
}
