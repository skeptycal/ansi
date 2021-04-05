// Copyright (c) 2020 Michael Treanor
// MIT License

package ansi

import "fmt"

const (
	// set ANSI 8-bit (256 color) foreground and background color with leading effect
	ansiEncode = "\033[%d;38;5;%d;48;5;%dm"
)

// Format Strings 8 bit Ansi printf commands.
//
// ESC[⟨x⟩8:5:⟨n⟩m
//
// Select 8bit color
//
// n in [0..255]; 0-231 are colors; 232-255 are grayscale
//
// (x in [3, 4]); 3 = foreground; 4 = background
const (
	fmt8bit   string = "\033[%d8;5;%dm"
	fmt8bitFG string = "\033[38;5;%dm"
	fmt8bitBG string = "\033[48;5;%dm"
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
