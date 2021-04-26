// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

// Package ansi provides fast ansi escape sequence processing based on strings.Builder.
// The standard is defined by the ECMA-48 standard "Control Functions for Coded Character Sets - Fifth Edition"
//
// The default zero values for Go variables are used whenever possible.
// Reference: https://golangbyexample.com/go-default-zero-value-all-types/
//
// ANSI escape codes for Command Line applications are well defined.
// Reference: https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html
//
// Many of the ANSI cursor codes are from https://github.com/k0kubun/go-ansi
package ansi

import "testing"

func BenchmarkRainbow1(b *testing.B) {

	b.Run("rainbow1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rainbow1(160, 196, "This is a rainbow string ...")
		}
	})
	b.Run("rainbow2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rainbow2(160, 196, "This is a rainbow string ...")
		}
	})
}
