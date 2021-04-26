// Copyright (c) 2020 Michael Treanor
// MIT License

package ansi

import (
	"bytes"
	"strings"
)

const (
	prefix   string = "\x1b["
	fg       string = "38;5;"
	bg       string = "48;5;"
	suffix   string = "m"
	ansiSep  string = ";"
	special1 string = prefix + suffix
	special2 string = prefix + suffix
	special3 string = prefix + suffix
	reset           = Reset
)

var (
	bprefix  []byte = []byte(prefix)
	bfg      []byte = []byte(fg)
	bbg      []byte = []byte(bg)
	bsuffix  []byte = []byte(suffix)
	bansiSep []byte = []byte(ansiSep)
	arr             = [32]byte{}
)

func ansiOneAppend(b ...byte) (retval []byte) {
	if len(b) == 0 {
		return bDefaultColor
	}
	retval = append(retval, bprefix...)
	retval = append(retval, b...)
	retval = append(retval, bsuffix...)
	return
}

// func ansiArray(b *[32]byte) []byte {

// 	len(b)

// 	return []byte{}
// 	// return []byte(b...)
// }

func ansiBytesBB3d(b ...byte) []byte {
	if len(b) == 0 {
		return bDefaultColor
	}
	defer buf.Reset()
	switch len(b) {
	case 1:
		// "\x1b[  38;5;  %d   m"
		buf.Write(bprefix)
		buf.Write(bfg)
		buf.WriteByte(b[0])
		buf.Write(bsuffix)
		return buf.Bytes()

	case 3:
		// "\x1b[  %d  ;  38;5;  %d  ;  48;5;  %d  m"
		buf.Write(bprefix)
		buf.WriteByte(b[2])
		buf.Write(bansiSep)
		buf.Write(bfg)
		buf.WriteByte(b[0])
		buf.Write(bansiSep)
		buf.Write(bbg)
		buf.WriteByte(b[1])
		buf.Write(bsuffix)
		return buf.Bytes()

	case 2:
		// "\x1b[   38;5;  %d  ;  48;5;  %d  m"
		buf.Write(bprefix)
		buf.Write(bfg)
		buf.WriteByte(b[0])
		buf.Write(bansiSep)
		buf.Write(bbg)
		buf.WriteByte(b[1])
		buf.Write(bsuffix)
		return buf.Bytes()

	default:
		// "\x1b[  %d  ;  %d  ;  %d  ;  38;5;  %d  ;  48;5;  %d  m"
		buf.Write(bprefix)
		buf.Write(bytes.Join(bytes.Split(b[2:], []byte{}), bansiSep))
		buf.Write(bansiSep)
		buf.Write(bfg)
		buf.WriteByte(b[0])
		buf.Write(bansiSep)
		buf.Write(bbg)
		buf.WriteByte(b[1])
		buf.Write(bsuffix)
		return buf.Bytes()
	}
}

func ansiBytesBB2d(b ...byte) []byte {
	if len(b) == 0 {
		return bDefaultColor
	}
	defer buf.Reset()
	switch len(b) {
	case 0:
		return bDefaultColor

	case 1:
		// "\x1b[  38;5;  %d   m"
		buf.Write(bprefix)
		buf.Write(bfg)
		buf.WriteByte(b[0])
		buf.Write(bsuffix)
		return buf.Bytes()

	case 2:
		// "\x1b[   38;5;  %d  ;  48;5;  %d  m"
		buf.Write(bprefix)
		buf.Write(bfg)
		buf.WriteByte(b[0])
		buf.Write(bansiSep)
		buf.Write(bbg)
		buf.WriteByte(b[1])
		buf.Write(bsuffix)
		return buf.Bytes()

	default:
		// "\x1b[  %d  ;  %d  ;  %d  ;  38;5;  %d  ;  48;5;  %d  m"
		buf.Write(bprefix)
		buf.Write(bytes.Join(bytes.Split(b[2:], []byte{}), bansiSep))
		buf.Write(bansiSep)
		buf.Write(bfg)
		buf.WriteByte(b[0])
		buf.Write(bansiSep)
		buf.Write(bbg)
		buf.WriteByte(b[1])
		buf.Write(bsuffix)
		return buf.Bytes()
	}
}

func ansiBytesSB3d(b ...byte) string {
	if len(b) == 0 {
		return DefaultColor
	}
	defer sb.Reset()
	switch len(b) {
	case 0:
		return DefaultColor

	case 1:
		// "\x1b[  38;5;  %d   m"
		sb.Write(bprefix)
		sb.Write(bfg)
		sb.WriteByte(b[0])
		sb.Write(bsuffix)
		return sb.String()

	case 2:
		// "\x1b[   38;5;  %d  ;  48;5;  %d  m"
		sb.Write(bprefix)
		sb.Write(bfg)
		sb.WriteByte(b[0])
		sb.Write(bansiSep)
		sb.Write(bbg)
		sb.WriteByte(b[1])
		sb.Write(bsuffix)
		return sb.String()

	case 3:
		// "\x1b[  %d  ;  38;5;  %d  ;  48;5;  %d  m"
		sb.Write(bprefix)
		sb.WriteByte(b[2])
		sb.Write(bansiSep)
		sb.Write(bfg)
		sb.WriteByte(b[0])
		sb.Write(bansiSep)
		sb.Write(bbg)
		sb.WriteByte(b[1])
		sb.Write(bsuffix)
		return sb.String()

	default:
		// "\x1b[  %d  ;  %d  ;  %d  ;  38;5;  %d  ;  48;5;  %d  m"
		sb.Write(bprefix)
		sb.Write(bytes.Join(bytes.Split(b[2:], []byte{}), bansiSep))
		sb.Write(bansiSep)
		sb.Write(bfg)
		sb.WriteByte(b[0])
		sb.Write(bansiSep)
		sb.Write(bbg)
		sb.WriteByte(b[1])
		sb.Write(bsuffix)
		return sb.String()
	}
}

func ansiBytesSB2d(b ...byte) string {
	if len(b) == 0 {
		return DefaultColor
	}
	defer sb.Reset()
	switch len(b) {
	case 0:
		return DefaultColor

	case 1:
		// "\x1b[  38;5;  %d   m"
		sb.Write(bprefix)
		sb.Write(bfg)
		sb.WriteByte(b[0])
		sb.Write(bsuffix)
		return sb.String()

	case 2:
		// "\x1b[   38;5;  %d  ;  48;5;  %d  m"
		sb.Write(bprefix)
		sb.Write(bfg)
		sb.WriteByte(b[0])
		sb.Write(bansiSep)
		sb.Write(bbg)
		sb.WriteByte(b[1])
		sb.Write(bsuffix)
		return sb.String()

	default:
		// "\x1b[  %d  ;  %d  ;  %d  ;  38;5;  %d  ;  48;5;  %d  m"
		sb.Write(bprefix)
		sb.Write(bytes.Join(bytes.Split(b[2:], []byte{}), bansiSep))
		sb.Write(bansiSep)
		sb.Write(bfg)
		sb.WriteByte(b[0])
		sb.Write(bansiSep)
		sb.Write(bbg)
		sb.WriteByte(b[1])
		sb.Write(bsuffix)
		return sb.String()
	}
}

var buf = bytes.Buffer{}
var sb = strings.Builder{}

// type bufPool struct {
// 	buf bytes.Buffer
// }
