package ansi

import "fmt"

const (
	// set ANSI 8-bit (256 color) foreground and background color with leading effect
	ansiEncode string = "\x1b[%v;38;5;%v;48;5;%vm"
	fgEncode   string = "\x1b[%d;38;5;%d;48;5;%dm%s%s"
)

// all this stuff is cool and fun, but it is slower than just using strings ...

type Ansier interface {
	String() string
}

type ansi struct {
	foreground string `default:"15"` // ANSI White ForeGround 37
	background string // default is Zero Value (0)
	effect     string // default is Zero Value (0)
}

func (c *ansi) String() string {
	return fmt.Sprintf(ansiEncode, c.effect, c.foreground, c.background)
}

func newColorStruct(foreground, background, effect string) *ansi {
	return &ansi{foreground, background, effect}
}
