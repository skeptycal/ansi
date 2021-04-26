package main

// ansi dev testing demo

import (
	"fmt"
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
	return fmt.Sprintf("%s%-20.20s (len %d,type %v) - %v - {abcdefghijklmnopqrstuvwxyz01234567890}%s\n", ansi.Reset, cmd, l, tt, v, ansi.Reset)
}

func main() {
	fmt.Println("")
	fmt.Println("dev function results examples")
	// effectSamples()
	// formatSamples()
	rainbowSamples()
	constantValues()
	// strPointerExample()
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

	strList := []string{ansi.Default, ansi.Reset, ansi.DefaultEffectString, ansi.DefaultForegroundString, ansi.DefaultBackgroundString, redstringer}

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

	sb.WriteString("=====================================================================================\n")
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
