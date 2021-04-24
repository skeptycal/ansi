package ansi

import (
	"fmt"
	"testing"
)

func Test_ansiString_String(t *testing.T) {
	tests := []struct {
		name  string
		input string
		ef    byte
		fg    byte
		bg    byte
	}{
		// TODO: Add test cases.
		{"172,0,1", "this string", 172, 0, 1},
		{"255,0,2", "this string", 255, 0, 2},
		{"2,4,7", "this string", 2, 4, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newcolor := NewColor(tt.fg, tt.bg, tt.ef)
			newstring := NewAnsiString(newcolor, tt.input)
			got := newstring.String()

			want := fmt.Sprintf("\x1b[%d;38;5;%d;48;5;%dm%s%s", tt.ef, tt.fg, tt.bg, tt.input, Reset)

			if got != want {
				t.Errorf("ansiString.String() = %v, want %v", got, want)
			}
		})
	}
}
