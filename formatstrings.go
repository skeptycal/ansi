package ansi

// Format Strings - 3/4 bit Ansi printf commands.
//
// ESC[⟨e⟩;⟨x⟩⟨n⟩m
//
// Select 3/4 bit color
//
// n in [0..7]; basic colors
//
// (x in [3, 4]); 3 = foreground; 4 = background
//
// e can be any valid ANSI escape effect; these are common:
//  0 Normal            // may be left out to maintain current ANSI effect
//  1 Bold              // bold or increased intensity
//  2 Faint             // faint, decreased intensity or second color
//  3 Italics
//  4 Underline
//  7 Inverse
//  8 Conceal
//  9 Strikeout
//
const (
	fmtANSI           string = "\033[%dm"
	fmtANSIFG         string = "\033[3%dm"
	fmtANSIBG         string = "\033[4%dm"
	fmtANSIWithEffect string = "\033[%d;%dm"
	fmtANSIBright     string = "\033[1;%dm"
	fmtANSIDim        string = "\033[2;%dm"
)

// Format Strings 24 bit Ansi printf commands.
//
// ESC[⟨x⟩8;2;⟨R⟩;⟨G⟩;⟨B⟩m
//
// Select RGB color
//
// R, G, B in [0..255]
//
// (x in [3, 4]); 3 = foreground; 4 = background
const (
	fmt24bit   string = "\033[%d8;2;%d;%d;%dm"
	fmt24bitFG string = "\033[38;2;%d;%d;%dm"
	fmt24bitBG string = "\033[48;2;%d;%d;%dm"
)
