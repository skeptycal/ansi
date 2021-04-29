package main

// ansi dev testing demo

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/skeptycal/ansi"
)

type Some struct{}

var (
	fg      byte = 17
	bg      byte = 214
	ef      byte = ansi.Italics
	rainbow      = ansi.Rainbow
)

func ExampleAnsi() error {
	tests := []struct {
		name string
		b    []byte
		want string
	}{
		{"no bytes(return default value)", []byte{}, "\x1b[0;38;5;7;48;5;0m"},
		{"1 byte (Green)", []byte{'2'}, "\x1b[38;5;2m"},
		{"2 bytes (fg/bg)", []byte{'2', '4'}, "\x1b[38;5;2;48;5;4m"},
		{"3 bytes (fg/bg/effect)", []byte{'2', '1', '1'}, "\x1b[1;38;5;2;48;5;1m"},
		{">3 bytes (fg/bg + extra effects)", []byte{'2', '4', '1', '3'}, "\x1b[1;3;38;5;2;48;5;4m"},
		{"bold red on red", []byte{'1', '1', '1'}, "\x1b[1;38;5;1;48;5;1m"},
		{"italic blue on red", []byte{'4', '1', '3'}, "\x1b[3;38;5;4;48;5;1m"},
		{"green bold text", []byte{'2', '0', '1'}, "\x1b[1;38;5;2;48;5;0m"},
	}
	for _, tt := range tests {
		got := string(ansiBytesBB3d(tt.b...))
		// if !reflect.DeepEqual(got, tt.want) {
		fmt.Printf(
			"  %25.25s: input = %6q,  got = %-27q (%vstuff%v),  want = %-27q (%vstuff%v)\n",
			tt.name,
			tt.b,
			got,
			got,
			ansi.Reset,
			tt.want,
			tt.want,
			ansi.Reset,
		)
		// }
	}

	// t := tests[7]

	fmt.Println()

	input := []string{"123", "123", "12"}
	fg := input[0]
	bg := input[1]
	ef := input[2]
	fmt.Printf("fg (ord: %v): %q\n", fg, fg)
	fmt.Printf("bg (ord: %v): %q\n", bg, bg)
	fmt.Printf("ef (ord: %v): %q\n", ef, ef)

	s := ansi.NewColor(fg, bg, ef).String()
	b := []byte(s)
	fmt.Printf("len(%s%v%s): %d\n", b, s, ansi.Reset, len(b))
	for i, c := range b {
		fmt.Printf("%d: %v (%q)\n", i, c, c)
	}
	printResult("example ansiBytes", b)

	fmt.Println()
	input = []byte{1, 0, 1}
	b = ansiBytesBB3d(input...)
	s = string(b)
	fmt.Printf("len(%s%v%s): %d\n", b, s, ansi.Reset, len(b))
	printResult("example ansiBytes", b)

	fmt.Println()
	return nil
}

func ansiBytesBB3d(b ...byte) []byte {
	const (
		prefix  string = "\x1b["
		fg      string = "38;5;"
		bg      string = "48;5;"
		suffix  string = "m"
		ansiSep string = ";"
	)
	if len(b) == 0 {
		return []byte(ansi.DefaultColor)
	}
	var buf = bytes.Buffer{}
	switch len(b) {
	case 1:
		// "\x1b[  38;5;  %d   m"
		buf.WriteString(prefix)
		buf.WriteString(fg)
		buf.WriteByte(b[0])
		buf.WriteString(suffix)
		return buf.Bytes()
		// "\x1b[  38;5;  %d   m"
		buf.WriteString(prefix)
		buf.WriteString(fg)
		buf.WriteByte(b[0])
		buf.WriteString(suffix)
		return buf.Bytes()
	case 3:
		// "\x1b[  %d  ;  38;5;  %d  ;  48;5;  %d  m"
		buf.WriteString(prefix)       // "\x1b["
		buf.WriteString(string(b[2])) // 3rd byte (effect)
		buf.WriteString(ansiSep)      // ";"
		buf.WriteString(fg)           // "38;5;"
		buf.WriteString(string(b[0])) // 1st byte (foreground)
		buf.WriteString(ansiSep)      // ";"
		buf.WriteString(bg)           // "48;5;"
		buf.WriteString(string(b[1])) // 2nd byte (background)
		buf.WriteString(suffix)       // "m"
		return buf.Bytes()
	case 2:
		// "\x1b[   38;5;  %d  ;  48;5;  %d  m"
		buf.WriteString(prefix)
		buf.WriteString(fg)
		buf.WriteByte(b[0])
		buf.WriteString(ansiSep)
		buf.WriteString(bg)
		buf.WriteByte(b[1])
		buf.WriteString(suffix)
		return buf.Bytes()
	default:
		// "\x1b[  %d  ;  %d  ;  %d  ;  38;5;  %d  ;  48;5;  %d  m"
		buf.WriteString(prefix)
		buf.WriteString(strings.Join(strings.Split(string(b[2:]), ""), ansiSep))
		buf.WriteString(ansiSep)
		buf.WriteString(fg)
		buf.WriteByte(b[0])
		buf.WriteString(ansiSep)
		buf.WriteString(bg)
		buf.WriteByte(b[1])
		buf.WriteString(suffix)
		return buf.Bytes()
		// "\x1b[  %d  ;  %d  ;  %d  ;  38;5;  %d  ;  48;5;  %d  m"
		buf.WriteString(prefix)
		buf.WriteString(strings.Join(strings.Split(string(b[2:]), ""), ansiSep))
		buf.WriteString(ansiSep)
		buf.WriteString(fg)
		buf.WriteByte(b[0])
		buf.WriteString(ansiSep)
		buf.WriteString(bg)
		buf.WriteByte(b[1])
		buf.WriteString(suffix)
		return buf.Bytes()
	}
}

func showResult(cmd string, v interface{}) string {
	// s := fmt.Sprintf("%v", v)

	var l int
	var tt string
	switch t := v.(type) {
	case string:
		l = len(v.(string))
		tt = "string"
	default:
		tt = fmt.Sprintf("%T", t)
		l = len(tt)
	}
	return fmt.Sprintf(
		"%s%-20.20s (len %d,type %v) - %v - {abcdefghijklmnopqrstuvwxyz01234567890}%s\n",
		ansi.Reset,
		cmd,
		l,
		tt,
		v,
		ansi.Reset,
	)
}

func main() {
	fmt.Println("")
	fmt.Println("dev function results examples")
	// effectSamples()
	// formatSamples()
	// rainbowSamples()
	// constantValues()
	// strPointerExample()
	err := ExampleAnsi()
	if err != nil {
		log.Fatal(err)
	}

}

func (p *Some) GetFieldName(f interface{}) (name string) {
	// https://stackoverflow.com/a/63745259

	val := reflect.ValueOf(p).Elem()
	val2 := reflect.ValueOf(f).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if valueField.Addr().Interface() == val2.Addr().Interface() {
			return val.Type().Field(i).Name
		}
	}
	return
}

func constantValues() {
	fmt.Println("")
	redstringer := ansi.NewColor(1, 0, 0).String()

	strList := []string{
		ansi.Default,
		ansi.Reset,
		ansi.DefaultEffectString,
		ansi.DefaultForegroundString,
		ansi.DefaultBackgroundString,
		redstringer,
	}

	for _, str := range strList {
		printResult(fmt.Sprintf("%q", str), str)
	}

	printResult("NewColor(1, 0, 0)", redstringer)
	fmt.Println(ToASCIIMap("NewColor(1, 0, 0)", []byte(redstringer)))

	red := "\x1b[31m"
	printResult("red", red)
	fmt.Println(ToASCIIMap("red", []byte(red)))

}

func ToASCIIMap(name string, b []byte) string {
	s1 := ""
	s2 := ""
	s3 := ""
	sb := strings.Builder{}
	defer sb.Reset()

	sb.WriteString(
		"=====================================================================================\n",
	)
	sb.WriteString(fmt.Sprintf("ASCIIMap of: %s\n", name))
	sb.WriteString(showResult(name, b))

	for i, c := range b {
		count := fmt.Sprint(i)
		n := int(c)
		// s := string(c)
		// n := []byte(s)

		s1 += fmt.Sprintf("%3v ", count)
		s2 += fmt.Sprintf("%3d ", n)
		s3 += fmt.Sprintf("%3c ", c)
	}
	sb.WriteString(s1)
	sb.WriteByte('\n')
	sb.WriteString(s2)
	sb.WriteByte('\n')
	sb.WriteString(s3)
	sb.WriteByte('\n')

	return sb.String()
}

func printResult(cmd string, v interface{}) string {
	retval := showResult(cmd, v)
	fmt.Println(retval)
	return retval
}

func effectSamples() {
	fmt.Println("")

	for i := 0; i < 10; i++ {
		printResult(fmt.Sprintf("ANSI effect: %2d", i), ansi.NewColor(bg, fg, byte(i)))
	}

	printResult("rainbow(160,196)", rainbow(160, 196, "Rainbow"))
}

func formatSamples() {
	fmt.Println("")
	a := ansi.NewColor(fg, bg, ef)
	printResult("ansi.NewColor(17, 214, 3)", a)

	s := a.String()
	printResult("a.String()", s)

	b := []byte(s)
	printResult("[]byte(s)", b)
}

func rainbowSamples() {
	fmt.Println("")
	fmt.Println(rainbow(88, 128, "This is a rainbow string ..."))
	fmt.Println(rainbow(160, 196, strings.Repeat("This is a rainbow string ...", 3)))
	fmt.Println(rainbow(160, 196, strings.Repeat("... and another rainbow string ...", 3)))
}

func castStr(v *string) string {
	return fmt.Sprint(uintptr(unsafe.Pointer(v)))
}

func uncastStr(s string) string {
	p, _ := strconv.ParseInt(s, 10, 64)
	return *((*string)(unsafe.Pointer(uintptr(p))))
}

func strVarName(s string) string {
	return uncastStr(s)
}

func strPointerExample() {
	// https://stackoverflow.com/a/24849574
	// onevar := "something"
	// other := "something else"
	// sa := []string{castStr(&onevar), castStr(&other)}

	// for _, v := range sa {
	// 	fmt.Printf("{{%s}}\n", v)
	// 	fmt.Printf("%v\n", uncastStr(v))
	// }

	s := "fake"

	printResult("string(fake)", s)
	p := castStr(&s)
	printResult("cast", p)
	printResult("pointer", &p)
	u := uncastStr(p)
	printResult("uncast", u)

	//for _, v := range sa {
	//  vName := fmt.Sprintf("{{%s}}", v)
	//  msg = strings.Replace(msg, vName, uncastStr(v) -1)
	//}
}
