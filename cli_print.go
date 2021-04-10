package ansi

import "fmt"

func (t *Terminal) doCheckColor() {
	if t.useColor {
		_, _ = t.Write(t.colorBytes)
	}
}

func (t *Terminal) noOp() {}

// Print wraps args in ANSI 8-bit color codes (256 color codes)
func (t *Terminal) Print(args ...interface{}) {
	t.on()
	fmt.Fprint(t.Writer, args...)
	t.Reset()
}

// Printf wraps args in ANSI 8-bit color codes (256 color codes)
func (t *Terminal) Printf(s string, args ...interface{}) {
	t.on()
	fmt.Fprintf(t.Writer, s, args...)
	t.Reset()
}

// Println wraps args in ANSI 8-bit color codes (256 color codes)
// and adds a newline character
func (t *Terminal) Println(args ...interface{}) {
	t.on()
	fmt.Fprintln(t.Writer, args...)
	t.Reset()
}
