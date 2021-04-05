package ansi

import "fmt"

func (t *Terminal) on() {
	if t.useColor {
		fmt.Fprint(t.Writer, t.Color)
	}
}

// Print wraps args in ANSI 8-bit color codes (256 color codes)
func (t *Terminal) Print(args ...interface{}) {
	t.on()
	fmt.Fprint(t.Writer, args...)
	t.Reset()
}

// Printf wraps args in ANSI 8-bit color codes (256 color codes)
func (t *Terminal) Printf(fmt string, args ...interface{}) {
	t.on()
	// fmt.Fprintf(t.Writer, fmt, args...)
	t.Reset()
}

// Println wraps args in ANSI 8-bit color codes (256 color codes)
// and adds a newline character
func (t *Terminal) Println(args ...interface{}) {
	t.on()
	fmt.Fprintln(t.Writer, args...)
	t.Reset()
}
