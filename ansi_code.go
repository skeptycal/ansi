// Copyright (c) 2020 Michael Treanor
// MIT License

package ansi

import (
	"fmt"
)

type Ansi interface {
	String() string
}

func NewColor(foregroundByte, backgroundByte, effectByte byte) Ansi {
	return &ansiCode{foregroundByte, backgroundByte, effectByte}
}

var current = struct {
	foreground, background, effect byte
}{2, 0, 1}

type ansiCode struct {
	foreground byte `default:"15"` // ANSI White ForeGround 37
	background byte // default is Zero Value (0)
	effect     byte // default is Zero Value (0)
}

func (c *ansiCode) String() string {
	return fmt.Sprintf(ansiEncode, c.effect, c.foreground, c.background)
}

// type bufPool struct {
// 	buf bytes.Buffer
// }
