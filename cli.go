package ansi

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
// defaultCLIforeground byte = 15
// defaultCLIbackground byte = 0
// defaultCLIeffect          = 0
)

var (
	defaultWriter      io.Writer = newAnsiStdout()
	defaultErrorWriter io.Writer = NewAnsiStderr()
)

// newAnsiStdout returns stdout which converts escape sequences
// to Windows API calls on Windows environment.
func newAnsiStdout() io.Writer {
	return os.Stdout
}

// NewAnsiStderr returns stdout which converts escape sequences
// to Windows API calls on Windows environment.
func NewAnsiStderr() io.Writer {
	return os.Stderr
}

// CLI implements an ANSI compatible terminal interface.
type CLI interface {
	String() string
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
	SetColor(color Ansi)
	Reset()
	CLIControls
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
	Writer                           io.Writer `default:"defaultWriter"`
	useColor                         bool      `default:"true"`
	devMode                          bool
	Color                            string
	DefaultForeground, fg            byte `default:"15"`
	DefaultBackground, DefaultEffect byte // default is Zero Value (0)
}

// Inverse sets the inverse ANSI effect if the terminal supports it.
func (t *Terminal) Inverse() {
	if t.useColor {
		t.Print(simpleEncode(Inverse))
	}
}

// Reset sets the ANSI foreground, background, and effect to default.
func (t *Terminal) Reset() {
	if t.useColor {
		fmt.Fprint(t.Writer, Reset)
	}
}

// SetColor sets the ANSI foreground, background, and effect codes
// for upcoming output.
func (t *Terminal) SetColor(color Ansi) {
	if t.useColor {
		t.Color = color.String()
	}
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
		fmt.Fprintf(sb, "%vANSI color output (fg = %v, bg = %v, ef = %v) %v\n", t.Color, t.DefaultForeground, t.DefaultBackground, t.DefaultEffect, Reset)
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
		fmt.Fprintf(sb, "CLI variable (Color): %q\n", t.Color)
		fmt.Fprintf(sb, "CLI variable (UseColor): %t\n", t.useColor)
		fmt.Fprintf(sb, "CLI variable (devMode): %t\n", t.devMode)
		fmt.Fprintf(sb, "CLI writer pointer: %v\n\n", &t.Writer)
	}
	return sb.String()
}
