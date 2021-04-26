package ansi

import (
	"fmt"
	"testing"
)

func BenchmarkAnsiString(b *testing.B) {

	const fmtfake = "\x1b[%d;38;5;%d;48;5;%dm%s%s"
	newcolor := NewColor(2, 0, 1)
	onlystring := "this string"
	fakeansi := ansiString{fmt.Sprintf("\x1b[%d;38;5;%d;48;5;%dm%s%s", 1, 2, 0, onlystring, Reset)}.s
	fakefmtansi := ansiString{fmt.Sprintf(fmtfake, 1, 2, 0, onlystring, Reset)}.s

	ansistringtest := NewAnsiString(newcolor, onlystring)
	onlystringer := ansistringtest.String()

	b.Run("no fmt.Sprintf()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = onlystring
		}
	})

	b.Run("NewAnsiString()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewAnsiString(newcolor, onlystring)
		}
	})

	b.Run("only raw, uncolored string (%s)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%s", onlystring)
		}
	})

	b.Run("only raw, uncolored string (%v)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%v", onlystring)
		}
	})

	b.Run("using printf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%v", fakeansi)
		}
	})

	b.Run("using printf with fmt constant", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%v", fakefmtansi)
		}
	})

	b.Run("passing ansiString.String()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%v", onlystringer)
		}
	})

	b.Run("ansiString with %s", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%s", ansistringtest)
		}
	})

	b.Run("passing ansiString object directly", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%v", ansistringtest)
		}
	})
}

func Test_ansiString(t *testing.T) {
	t.Run("ansiString struct test", func(t *testing.T) {
		newcolor := NewColor(2, 0, 1)
		newstring := "this string"

		want := NewAnsiString(newcolor, newstring).(*ansiString).s
		got := ansiString{fmt.Sprintf("\x1b[%d;38;5;%d;48;5;%dm%s%s", 1, 2, 0, newstring, Reset)}.s

		if len(want) != len(got) {
			t.Errorf("len(ansiString) = %v, want %v", len(got), len(want))
		}

		if got != want {
			t.Errorf("&ansiString = %v, want %v", got, want)
		}
	})
}

func Test_ansiString_String(t *testing.T) {
	tests := []struct {
		name  string
		input string
		ef    byte
		fg    byte
		bg    byte
	}{
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

// benchmark results
/*
BenchmarkAnsiString/no_fmt.Sprintf()-8         				  1000000000	      0.3216 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnsiString/NewAnsiString()-8          	 				 2142024	       552.2 ns/op	      88 B/op	       3 allocs/op
BenchmarkAnsiString/only_raw,_uncolored_string_(%s)-8         	10113024	       124.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkAnsiString/only_raw,_uncolored_string_(%v)-8         	10009150	       121.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkAnsiString/using_printf-8                            	 9181312	       133.4 ns/op	      64 B/op	       2 allocs/op
BenchmarkAnsiString/using_printf_with_fmt_constant-8          	 9243387	       130.7 ns/op	      64 B/op	       2 allocs/op
BenchmarkAnsiString/passing_ansiString.String()-8             	 9298756	       130.8 ns/op	      64 B/op	       2 allocs/op
BenchmarkAnsiString/ansiString_with_%s-8                      	 5895320	       209.8 ns/op	      48 B/op	       1 allocs/op
BenchmarkAnsiString/passing_ansiString_object_directly-8      	 5865824	       207.8 ns/op	      48 B/op	       1 allocs/op

Conclusions:
/// "Do the work once and store the result."
/// "If the work is only used once, do it another way."

for one time use, creating an ansiString object is 5x slower
for strings that are reused, ansiString.String() is comparable

clear and efficient code that relies on premade strings should benefit
constant streams of varying strings would be much more efficient another way

For premade strings:

Using printf manually is 7% slower than ansiString.String() and *much* more difficult to read
Passing the ansiString object is nearly 2x slower - use the Stringer
Passing the original uncolored string OR the colored Stringer is nearly equivalent
Stringer() takes the largest allocation
*/
