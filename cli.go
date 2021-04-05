package ansi

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	defaultDevMode            = false
	defaultCLIforeground byte = 15
	defaultCLIbackground byte = 0
	defaultCLIeffect          = 0

	hrChar             string = "="
	defaultScreenWidth int    = 80
)

var (
	defaultWriter      io.Writer = newAnsiStdout()
	defaultErrorWriter io.Writer = NewAnsiStderr()
)

// newAnsiStdout returns stdout which converts escape sequences to Windows
// API calls on Windows environment.
func newAnsiStdout() io.Writer {
	return os.Stdout
}

// NewAnsiStderr returns stdout which converts escape sequences to Windows
// API calls on Windows environment.
func NewAnsiStderr() io.Writer {
	return os.Stderr
}

// CLI implements an ANSI compatible terminal interface.
type CLI interface {
	String() string
	Print(args ...interface{})
	Println(args ...interface{})
	CLS()
	Fg(color byte)
	Bg(color byte)
	Ef(color byte)
	Hr()
	Br()
}

func NewStdout(w io.Writer) CLI {
	checkColor := true // todo bring this code over ...
	devMode := defaultDevMode
	if w == nil {
		w = defaultWriter
	}
	if _, ok := w.(io.Writer); !ok {
		w = defaultWriter
	}

	return &Terminal{
		Writer:   w,
		useColor: checkColor,
		devMode:  devMode,
	}
}

func NewStderr(w io.Writer) CLI {
	checkColor := true // todo bring this code over ...
	devMode := defaultDevMode
	if w == nil {
		w = defaultErrorWriter
	}
	if _, ok := w.(io.Writer); !ok {
		w = defaultErrorWriter
	}

	return &Terminal{
		Writer:   w,
		useColor: checkColor,
		devMode:  devMode,
	}
}

type Terminal struct {
	Writer                                   io.Writer `default:"defaultWriter"`
	useColor                                 bool      `default:"true"`
	devMode                                  bool
	Color                                    string
	DefaultForeground, fg                    byte `default:"15"`
	DefaultBackground, DefaultEffect, bg, ef byte // default is Zero Value (0)
}

// Inverse sets the inverse ANSI effect if the terminal supports it.
func (t *Terminal) Inverse() {
	if t.useColor {
		t.Print(simpleEncode(Inverse))
	}
}

// LineBreak creates a CLI line break (dotted line, underline, etc.)
// by repeating the string c enough times to span the screen width.
//
// This is useful in delimiting lines of text in terminal output.
func (t *Terminal) LineBreak(c string) {
	if c == "" {
		c = hrChar
	}
	width := t.screenWidth()

	s := strings.Repeat(c, width/len(c)+1)[:width]
	t.Println(s)
}

// Hr creates a hard return using the default character.
func (t *Terminal) Hr() {
	t.LineBreak(hrChar)
}

// Br creates a line break.
func (t *Terminal) Br() {
	fmt.Fprintln(t.Writer, "")
}

func (t *Terminal) Reset() {
	if t.useColor {
		fmt.Fprint(t.Writer, Reset)
	}
}

func (t *Terminal) Fg(color byte) {
	if t.useColor {
		t.fg = color
		t.setColorString()
	}
}

func (t *Terminal) Bg(color byte) {
	if t.useColor {
		t.bg = color
		t.setColorString()
	}
}

func (t *Terminal) Ef(color byte) {
	if t.useColor {
		t.ef = color
		t.setColorString()
	}
}

func (t *Terminal) screenWidth() int {
	// todo - add support for variable width
	return defaultScreenWidth
}

func (t *Terminal) setColorString() {
	t.Color = t.encode()
}

func (t *Terminal) encode() string {
	return encode(t.fg, t.bg, t.ef)
}

// String describes the terminal. If devMode is true, it generates a
// list of dev info.
func (t *Terminal) String() string {
	sb := &strings.Builder{}
	defer sb.Reset()
	if t.devMode {
		sb.WriteString(t.devinfo())
	}
	if !t.useColor {
		sb.WriteString("ANSI terminal - no color output.")
	} else {
		fmt.Fprintf(sb, "%vANSI color output (fg = %v, bg = %v, ef = %v) %v\n", t.encode(), t.DefaultForeground, t.DefaultBackground, t.DefaultEffect, Reset)
	}
	return sb.String()
}

func (t *Terminal) devinfo() string {
	sb := &strings.Builder{}
	defer sb.Reset()
	if t.devMode {
		t.Hr()
		fmt.Fprintf(sb, "CLI variable (DefaultForeground): %v\n", t.DefaultForeground)
		fmt.Fprintf(sb, "CLI variable (DefaultBackground): %v\n", t.DefaultBackground)
		fmt.Fprintf(sb, "CLI variable (DefaultEffect): %v\n", t.DefaultEffect)
		fmt.Fprintf(sb, "CLI variable (fg): %v\n", t.fg)
		fmt.Fprintf(sb, "CLI variable (bg): %v\n", t.bg)
		fmt.Fprintf(sb, "CLI variable (ef): %v\n", t.ef)
		fmt.Fprintf(sb, "CLI variable (UseColor): %t\n", t.useColor)
		fmt.Fprintf(sb, "CLI variable (devMode): %t\n", t.devMode)
		fmt.Fprintf(sb, "CLI writer pointer: %v\n\n", &t.Writer)
	}
	return sb.String()
}

func (t *Terminal) UseColor(b bool) {
	t.useColor = b
}

func (t *Terminal) _DevMode(b bool) {
	t.devMode = b
}