// Copyright (c) 2020 Michael Treanor
// MIT License

package ansi

import "fmt"

const (
	// set ANSI 8-bit (256 color) foreground and background color with leading effect
	ansiEncode = "\033[%d;38;5;%d;48;5;%dm" // "\033[1;38;5;2;48;5;40m"
)

type Ansi interface {
	String() string
}

func NewColor(foregroundByte, backgroundByte, effectByte byte) Ansi {
	return &ansiCode{
		foreground: foregroundByte,
		background: backgroundByte,
		effect:     effectByte,
	}
}

type ansiCode struct {
	foreground byte `default:"15"` // ANSI White ForeGround 37
	background byte // default is Zero Value (0)
	effect     byte // default is Zero Value (0)
}

func (c *ansiCode) String() string {
	return fmt.Sprintf(ansiEncode, c.effect, c.foreground, c.background)
}
